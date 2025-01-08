package interpreter

import (
	"github.com/razzzp/apdu-interpreter/apdu"
	"github.com/razzzp/apdu-interpreter/schema"
)

type responseInterpreter struct {
	ResponseDef *schema.ResponseDefinition
	Sw1         ByteInterpreter
	Sw2         ByteInterpreter
	// Data
}

func (ri *responseInterpreter) Matches(response *apdu.ApduResponse) bool {
	matches := true
	if ri.Sw1 != nil {
		matches = matches && ri.Sw1.Matches(response.SW1)
	}
	if ri.Sw2 != nil {
		matches = matches && ri.Sw2.Matches(response.SW2)
	}
	return matches
}

func (ri *responseInterpreter) Interpret(response *apdu.ApduResponse) (*ResponseInterpretation, error) {
	result := &ResponseInterpretation{}
	result.Intps.Add(ri.ResponseDef.Description)
	return result, nil
}
