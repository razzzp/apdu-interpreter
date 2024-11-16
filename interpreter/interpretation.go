package interpreter

type ByteInterpretation struct {
	Intps []any
}

func (bi *ByteInterpretation) Add(intp any) {
	bi.Intps = append(bi.Intps, intp)
}
