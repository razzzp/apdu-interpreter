package formatter

import (
	"github.com/razzzp/apdu-interpreter/interpreter"
)

type InterpretationWriter interface {
	Write(interpretations []*interpreter.ApduInterpretation)
}
