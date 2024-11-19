package interpreter

import "github.com/razzzp/apdu-interpreter/apdu"

type Interpretation interface {
	Add(intp any)
}

type ByteIntp interface {
	Matches(b byte) bool
	Interpret(i Interpretation, b byte) error
}

type Interpreter interface {
	Matches(apdu *apdu.ApduCommand) bool
	Interpret(apdu *apdu.ApduCommand) (Interpretation, error)
}
