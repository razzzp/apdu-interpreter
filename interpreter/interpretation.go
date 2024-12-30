package interpreter

import "github.com/razzzp/apdu-interpreter/apdu"

type ByteInterpretation struct {
	Intps []any
}

func (bi *ByteInterpretation) Add(intp any) {
	bi.Intps = append(bi.Intps, intp)
}

type CommandInterpretation struct {
	Command     *apdu.ApduCommand
	Interpreter *ApduCommandInterpreter
	ClaIntp     *ByteInterpretation
	InsIntp     *ByteInterpretation
	P1Intp      *ByteInterpretation
	P2Intp      *ByteInterpretation
	P3Intp      *ByteInterpretation
}
