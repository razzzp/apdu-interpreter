package formatter

import (
	"io"

	"github.com/razzzp/apdu-interpreter/interpreter"
)

type InterpretationWriter interface {
	Write(interpretations []*interpreter.ApduInterpretation, writer io.StringWriter)
}
