package interpreter

import "apdu-interpreter/apdu"

type Interpreter interface {
	Matches(apdu apdu.ApduCommand) bool
	Interpret(apdu apdu.ApduCommand) (Interpretation, error)
}

type ApduCommandInterpreter struct {
}

type Interpretation struct {
}
