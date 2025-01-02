package formatter

import (
	"io"
	"log"

	"github.com/razzzp/apdu-interpreter/interpreter"
)

type textInlineWriter struct {
	CmdColWidth  int
	IntpColWidth int
	writer       io.StringWriter
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
		CmdColWidth:  cmdColWidth,
		IntpColWidth: intpColWidth,
		writer:       writer,
	}
}

func (tf *textInlineWriter) formatBytes(bytes []byte) string {
	// TODO
	return ""
}

func (tf *textInlineWriter) writeInterpretation(interpretation *interpreter.ApduInterpretation) {
	cmdAsHex := tf.formatBytes(interpretation.Command.Command.AsBytes())
	// TODO
	log.Print(cmdAsHex)
}

func (tf *textInlineWriter) Write(interpretations []*interpreter.ApduInterpretation, writer io.StringWriter) {
	for _, intp := range interpretations {
		tf.writeInterpretation(intp)
	}
}
