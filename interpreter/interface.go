package interpreter

import "apdu-interpreter/apdu"

type ByteIntp interface {
	Matches(b byte) bool
	Interpret(b byte) (map[string]any, error)
}

type Interpreter interface {
	Matches(apdu apdu.ApduCommand) bool
	Interpret(apdu apdu.ApduCommand) (Interpretation, error)
}
