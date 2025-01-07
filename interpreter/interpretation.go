package interpreter

import "github.com/razzzp/apdu-interpreter/apdu"

type ByteInterpretations struct {
	Intps []any
}

func (bi *ByteInterpretations) Count() int {
	return len(bi.Intps)
}

type DataInterpretations struct {
}

func (bi *ByteInterpretations) Add(intp any) {
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
	ClaIntp  *ByteInterpretations
	InsIntp  *ByteInterpretations
	P1Intp   *ByteInterpretations
	P2Intp   *ByteInterpretations
	P3Intp   *ByteInterpretations
	DataIntp *DataInterpretations
}

func NewCommandInterpretation() CommandInterpretation {
	return CommandInterpretation{
		ClaIntp: &ByteInterpretations{},
		InsIntp: &ByteInterpretations{},
		P1Intp:  &ByteInterpretations{},
		P2Intp:  &ByteInterpretations{},
		P3Intp:  &ByteInterpretations{},
	}
}

type ResponseInterpretation struct {
	// TODO
}
