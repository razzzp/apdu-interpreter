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
	ApduInterpreter *ApduInterpreter
	Command         *CommandInterpretation
	Response        *ResponseInterpretation
}

type CommandInterpretation struct {
	Command     *apdu.ApduCommand
	Interpreter ApduCommandInterpreter
	ClaIntp     *ByteInterpretations
	InsIntp     *ByteInterpretations
	P1Intp      *ByteInterpretations
	P2Intp      *ByteInterpretations
	P3Intp      *ByteInterpretations
	DataIntp    *DataInterpretations
}

func NewCommandInterpretation(command *apdu.ApduCommand, interpreter ApduCommandInterpreter) CommandInterpretation {
	return CommandInterpretation{
		Command:     command,
		Interpreter: interpreter,
		ClaIntp:     &ByteInterpretations{},
		InsIntp:     &ByteInterpretations{},
		P1Intp:      &ByteInterpretations{},
		P2Intp:      &ByteInterpretations{},
		P3Intp:      &ByteInterpretations{},
	}
}

type ResponseInterpretation struct {
	// TODO
}
