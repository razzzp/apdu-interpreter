package interpreter

type ByteInterpretation struct {
	Intps []any
}

func (bi *ByteInterpretation) Add(intp any) {
	bi.Intps = append(bi.Intps, intp)
}

type CommandInterpretation struct {
	Command string
	ClaIntp ByteIntp
	InsIntp ByteIntp
	P1Intp  ByteIntp
	P2Intp  ByteIntp
	P3Intp  ByteIntp
}
