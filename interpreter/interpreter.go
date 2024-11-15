package interpreter

import "apdu-interpreter/apdu"

type Interpreter interface {
	Matches(apdu apdu.ApduCommand) bool
	Interpret(apdu apdu.ApduCommand) (Interpretation, error)
}

type ApduCommandInterpreter struct {
	Name            string
	Description     string
	ClaMatcher      Matcher
	InsMatcher      Matcher
	P1Matcher       Matcher
	P2Matcher       Matcher
	P3Matcher       Matcher
	DataInterpreter any
	LeMatcher       Matcher
}
