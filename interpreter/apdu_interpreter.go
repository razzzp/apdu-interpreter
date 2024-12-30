package interpreter

import "github.com/razzzp/apdu-interpreter/apdu"

type ApduCommandInterpreter struct {
	Name            string
	Description     string
	ClaMatcher      []ByteInterpreter
	InsMatcher      []ByteInterpreter
	P1Matcher       []ByteInterpreter
	P2Matcher       []ByteInterpreter
	P3Matcher       []ByteInterpreter
	DataInterpreter any
	LeMatcher       ByteInterpreter
}

func (aci *ApduCommandInterpreter) MatchesCla(apdu apdu.ApduCommand) bool {
	for _, matcher := range aci.ClaMatcher {
		if matcher.Matches(apdu.Cla) {
			return true
		}
	}
	return false
}
func (aci *ApduCommandInterpreter) MatchesIns(apdu apdu.ApduCommand) bool {
	for _, matcher := range aci.InsMatcher {
		if matcher.Matches(apdu.Ins) {
			return true
		}
	}
	return false
}

func (aci *ApduCommandInterpreter) Interpret(apdu apdu.ApduCommand) (Interpretation, error) {
	panic("not implemented") // TODO: Implement
}
