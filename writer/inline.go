package formatter

import (
	"encoding/hex"
	"io"

	"github.com/razzzp/apdu-interpreter/interpreter"
)

type textInlineWriter struct {
}

func (tf *textInlineWriter) Write(interpretations []*interpreter.ApduInterpretation, writer io.StringWriter) {
	for _, intp := range interpretations {
		cmdAsHex := hex.EncodeToString(intp.Command.Command.AsBytes())
		writer.WriteString(cmdAsHex)
	}
}
