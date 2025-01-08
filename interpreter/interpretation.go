package interpreter

import "github.com/razzzp/apdu-interpreter/apdu"

type GenericInterpretations struct {
	Intps []any
}

func (bi *GenericInterpretations) Count() int {
	return len(bi.Intps)
}

type DataInterpretations struct {
}

func (bi *GenericInterpretations) Add(intp any) {
	bi.Intps = append(bi.Intps, intp)
}

type ApduInterpretation struct {
	CommandResponse *apdu.ApduCommandResponse
	Interpretations []*CommandResponseInterpretation
}

type CommandResponseInterpretation struct {
	Interpreter  *ApduInterpreter
	CommandIntp  *CommandInterpretation
	ResponseIntp *ResponseInterpretation
}

type CommandInterpretation struct {
	ClaIntp  *GenericInterpretations
	InsIntp  *GenericInterpretations
	P1Intp   *GenericInterpretations
	P2Intp   *GenericInterpretations
	P3Intp   *GenericInterpretations
	DataIntp *DataInterpretations
}

func NewCommandInterpretation() CommandInterpretation {
	return CommandInterpretation{
		ClaIntp: &GenericInterpretations{},
		InsIntp: &GenericInterpretations{},
		P1Intp:  &GenericInterpretations{},
		P2Intp:  &GenericInterpretations{},
		P3Intp:  &GenericInterpretations{},
	}
}

type ResponseInterpretation struct {
	Intps GenericInterpretations
}
