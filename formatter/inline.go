package formatter

import (
	"encoding/hex"
	"fmt"
	"io"
	"strings"

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
			// default width is 5 chars
			colWidth := 5
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
func (tf *textInlineWriter) generateTableByteIntps(label string, rowIdx int, bintp *interpreter.ByteInterpretations, table *Table) int {

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

func byteAsHex(b byte) string {
	return hex.EncodeToString([]byte{b})[0:]
}

func (tf *textInlineWriter) generateTable(interpretations []*interpreter.ApduInterpretation) *Table {
	cmdWidth := tf.Widths[0]
	// intpWidth := tf.Widths[1]
	result := &Table{
		Widths:      tf.Widths,
		ColSepWidth: 3,
	}

	intpIdx := 0
	for _, intp := range interpretations {

		// add rows for cmd first
		cmd := intp.Command.Command
		curCmdBytes := tf.formatBytes(intp.Command.Command.AsBytes())
		cmdLines := len(curCmdBytes)/int(cmdWidth) + 1
		for i := 0; i < cmdLines; i++ {
			endIdx := min(len(curCmdBytes), (i+1)*int(cmdWidth))
			curLine := curCmdBytes[i*int(cmdWidth) : endIdx]
			if i != 0 {
				// indent following lines
				curLine = "  " + curLine
			}
			result.SetValue(intpIdx+i, 0, curLine)
		}
		curRow := intpIdx

		// no interpretation found
		if intp.ApduInterpreter == nil {
			// set unknown command and continue
			result.SetValue(curRow, 1, "Unknown Command")
			intpIdx = curRow + 1
			continue
		}
		// add matching cmd desc
		curRow = tf.generateTableInterpreter(curRow, intp.ApduInterpreter, result)
		// add interpretations
		curRow = tf.generateTableByteIntps(fmt.Sprintf("CLA: %s", byteAsHex(cmd.Cla)), curRow, intp.Command.ClaIntp, result)
		curRow = tf.generateTableByteIntps(fmt.Sprintf("INS: %s", byteAsHex(cmd.Ins)), curRow, intp.Command.InsIntp, result)
		curRow = tf.generateTableByteIntps(fmt.Sprintf("P1: %s", byteAsHex(cmd.P1)), curRow, intp.Command.P1Intp, result)
		curRow = tf.generateTableByteIntps(fmt.Sprintf("P2: %s", byteAsHex(cmd.P2)), curRow, intp.Command.P2Intp, result)
		p3Label := "P3:"
		if cmd.P3 != nil {
			p3Label = fmt.Sprintf("P3: %s", byteAsHex(*cmd.P3))
		}
		_ = tf.generateTableByteIntps(p3Label, curRow, intp.Command.P3Intp, result)
		intpIdx = curRow
	}
	return result
}

func (tf *textInlineWriter) Write(interpretations []*interpreter.ApduInterpretation) {

	table := tf.generateTable(interpretations)
	//print table
	table.Print(tf.writer)
}
