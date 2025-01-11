package bytearray

import "github.com/razzzp/apdu-interpreter/interpreter"

// length-value interpreter
type LvInterpreter struct {
	Label string
}

func (li *LvInterpreter) Interpret(i interpreter.Interpretations, b []byte, startIdx int) (int, error) {
	panic("not implemented") // TODO: Implement
}
