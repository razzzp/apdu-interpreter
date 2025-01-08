package formatter

import (
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/razzzp/apdu-interpreter/apdu"
	"github.com/razzzp/apdu-interpreter/interpreter"
)

type textInlineWriter struct {
	Widths []uint
	writer io.StringWriter
}

type Table struct {
	Widths      []uint
	ColSepWidth uint
	Rows        []*Row
}

func (t *Table) DefaultColWidth() int {
	return 10
}

func (t *Table) GetColWidth(colIdx int) int {
	if len(t.Widths) > colIdx {
		return int(t.Widths[colIdx])
	} else {
		return t.DefaultColWidth()
	}
}

func (t *Table) SetValue(rowIdx int, colIdx int, val string) {
	rowDiff := rowIdx - (len(t.Rows) - 1)
	if rowDiff > 0 {
		for i := 0; i < rowDiff; i++ {
			t.Rows = append(t.Rows, &Row{})
		}
	}
	//get row
	row := t.Rows[rowIdx]
	colDiff := colIdx - (len(row.Columns) - 1)
	if colDiff > 0 {
		for i := 0; i < colDiff; i++ {
			row.Columns = append(row.Columns, &Column{})
		}
	}

	col := row.Columns[colIdx]
	col.Value = val
}
func (t *Table) Print(writer io.StringWriter) {
	for _, row := range t.Rows {
		for i, col := range row.Columns {
			// default width
			colWidth := t.DefaultColWidth()
			if i < len(t.Widths) {
				colWidth = int(t.Widths[i])
			}

			valToPrint := ""
			if len(col.Value) > 0 {
				endIdx := len(col.Value)
				// cut string if too long
				endIdx = min(colWidth, endIdx)
				valToPrint = col.Value[0:endIdx]
			}
			// pad if too short
			diffWidth := colWidth - len(valToPrint)
			if diffWidth > 0 {
				valToPrint += strings.Repeat(" ", diffWidth)
			}
			writer.WriteString(valToPrint)
			// col sep
			writer.WriteString(strings.Repeat(" ", int(t.ColSepWidth)))
		}
		writer.WriteString("\n")
	}
}

type Row struct {
	Columns []*Column
}
type Column struct {
	Value string
}

func NewTextInlineWriter(cmdColWidth int, intpColWidth int, writer io.StringWriter) *textInlineWriter {
	// ensure multiple of 3
	for cmdColWidth%3 != 0 {
		cmdColWidth++
	}
	for intpColWidth%3 != 0 {
		intpColWidth++
	}
	return &textInlineWriter{
		Widths: []uint{uint(cmdColWidth), uint(intpColWidth)},
		writer: writer,
	}
}

func (tf *textInlineWriter) formatBytes(bytes []byte) string {

	return hex.EncodeToString(bytes)
}
func (tf *textInlineWriter) generateTableByteIntps(label string, rowIdx int, bintp *interpreter.GenericInterpretations, table *Table) int {
	if bintp == nil || bintp.Count() == 0 {
		return rowIdx
	}
	table.SetValue(rowIdx, 1, label)
	rowIdx++
	for _, intp := range bintp.Intps {
		table.SetValue(rowIdx, 1, fmt.Sprintf("  %v", intp))
		rowIdx++
	}
	return rowIdx
}

func (tf *textInlineWriter) generateTableInterpreter(rowIdx int, intpr *interpreter.ApduInterpreter, table *Table) int {
	table.SetValue(rowIdx, 1, fmt.Sprintf("%s : %s %s", intpr.SchemaDef.Group, intpr.SchemaDef.Name, intpr.SchemaDef.Version))
	rowIdx++
	table.SetValue(rowIdx, 1, intpr.CommandResponseDef.Name)
	rowIdx++
	table.SetValue(rowIdx, 1, intpr.CommandResponseDef.Description)
	rowIdx++
	return rowIdx
}

func (tf *textInlineWriter) generateTableCommandBreakdown(rowIdx int, cmd *apdu.ApduCommand, table *Table) int {
	table.SetValue(rowIdx, 1, fmt.Sprintf("CLA: %s", byteAsHex(cmd.Cla)))
	rowIdx++
	table.SetValue(rowIdx, 1, fmt.Sprintf("INS: %s", byteAsHex(cmd.Ins)))
	rowIdx++
	table.SetValue(rowIdx, 1, fmt.Sprintf("P1: %s", byteAsHex(cmd.P1)))
	rowIdx++
	table.SetValue(rowIdx, 1, fmt.Sprintf("P2: %s", byteAsHex(cmd.P2)))
	rowIdx++
	if cmd.P3 != nil {
		table.SetValue(rowIdx, 1, fmt.Sprintf("P3: %s", byteAsHex(*cmd.P3)))
		rowIdx++
	}
	return rowIdx
}

func (tf *textInlineWriter) generateTableResponse(rowIdx int, resp *apdu.ApduResponse, table *Table) int {
	table.SetValue(rowIdx, 1, fmt.Sprintf("SW1: %s", byteAsHex(resp.SW1)))
	rowIdx++
	table.SetValue(rowIdx, 1, fmt.Sprintf("SW2: %s", byteAsHex(resp.SW2)))
	rowIdx++
	return rowIdx
}

func (tf *textInlineWriter) generateTableResponseIntp(rowIdx int, resp *interpreter.ResponseInterpretation, table *Table) int {
	for _, intp := range resp.Intps.Intps {
		table.SetValue(rowIdx, 1, fmt.Sprintf("%v", intp))
		rowIdx++
	}
	return rowIdx
}

func (tf *textInlineWriter) generateTableBytesHex(rowIdx, colIdx int, bytes []byte, table *Table) int {
	curCmdBytes := tf.formatBytes(bytes)
	maxWidth := table.GetColWidth(colIdx)
	cmdLines := len(curCmdBytes)/int(maxWidth) + 1
	for i := 0; i < cmdLines; i++ {
		endIdx := min(len(curCmdBytes), (i+1)*int(maxWidth))
		curLine := curCmdBytes[i*int(maxWidth) : endIdx]
		if i != 0 {
			// indent following lines
			curLine = "  " + curLine
		}
		table.SetValue(rowIdx+i, 0, curLine)
	}
	return rowIdx + cmdLines
}

func byteAsHex(b byte) string {
	return hex.EncodeToString([]byte{b})[0:]
}

func (tf *textInlineWriter) generateTable(interpretations []*interpreter.ApduInterpretation) *Table {
	// intpWidth := tf.Widths[1]
	table := &Table{
		Widths:      tf.Widths,
		ColSepWidth: 3,
	}

	intpIdx := 0
	for _, apduIntp := range interpretations {

		// add column for cmd first
		cmd := apduIntp.CommandResponse.Command
		_ = tf.generateTableBytesHex(intpIdx, 0, cmd.AsBytes(), table)

		curIntpRow := intpIdx
		// add command breakdown
		curIntpRow = tf.generateTableCommandBreakdown(curIntpRow, apduIntp.CommandResponse.Command, table)

		// no interpretation found
		if len(apduIntp.Interpretations) == 0 {
			// set unknown command and continue
			table.SetValue(curIntpRow, 1, "Unknown Command")
			curIntpRow++

		} else {
			for _, intp := range apduIntp.Interpretations {
				// add matching cmd desc
				curIntpRow = tf.generateTableInterpreter(curIntpRow, intp.Interpreter, table)
				// add interpretations
				curIntpRow = tf.generateTableByteIntps("CLA:", curIntpRow, intp.CommandIntp.ClaIntp, table)
				curIntpRow = tf.generateTableByteIntps("INS:", curIntpRow, intp.CommandIntp.InsIntp, table)
				curIntpRow = tf.generateTableByteIntps("P1:", curIntpRow, intp.CommandIntp.P1Intp, table)
				curIntpRow = tf.generateTableByteIntps("P2:", curIntpRow, intp.CommandIntp.P2Intp, table)
				if cmd.P3 != nil {
					_ = tf.generateTableByteIntps("P3:", curIntpRow, intp.CommandIntp.P3Intp, table)
				}

				// add response
				resp := apduIntp.CommandResponse.Response
				_ = tf.generateTableBytesHex(curIntpRow, 0, resp.AsBytes(), table)

				// add response interpretation
				curIntpRow = tf.generateTableResponse(curIntpRow, resp, table)
				curIntpRow = tf.generateTableResponseIntp(curIntpRow, intp.ResponseIntp, table)
			}
		}

		intpIdx = curIntpRow
	}
	return table
}

func (tf *textInlineWriter) Write(interpretations []*interpreter.ApduInterpretation) {

	table := tf.generateTable(interpretations)
	//print table
	table.Print(tf.writer)
}
