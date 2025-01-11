package interpreter

import "github.com/razzzp/apdu-interpreter/apdu"

type dataInterpreter struct {
	Criteria     ApduInterpreter
	Interpreters []ByteArrayInterpreter
}

func (di *dataInterpreter) Matches(response *apdu.ApduCommand) bool {
	panic("not implemented") // TODO: Implement
}

func (di *dataInterpreter) Interpret(apdu *apdu.ApduCommand) (*DataInterpretations, error) {
	panic("not implemented") // TODO: Implement
}
