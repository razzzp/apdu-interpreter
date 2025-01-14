package interpreter

import "github.com/razzzp/apdu-interpreter/apdu"

type dataInterpreter struct {
	Criteria     CommandInterpreter
	Interpreters []ByteArrayInterpreter
}

func (di *dataInterpreter) Interpret(apdu *apdu.ApduCommand) (*DataInterpretations, error) {
	panic("not implemented") // TODO: Implement
}
