package interpreter

import "github.com/razzzp/apdu-interpreter/apdu"

type ApduCommandInterpreter struct {
	Name            string
	Description     string
	ClaMatcher      []ByteIntp
	InsMatcher      []ByteIntp
	P1Matcher       []ByteIntp
	P2Matcher       []ByteIntp
	P3Matcher       []ByteIntp
	DataInterpreter any
	LeMatcher       ByteIntp
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
