package interpreter

import "github.com/razzzp/apdu-interpreter/apdu"

type ApduInterpreter struct {
	CommandInterpreter  ApduCommandInterpreter
	ResponseInterpreter ApduReponseInterpreter
}

type apduCommandInterpreter struct {
	ClaMatcher      []ByteInterpreter
	InsMatcher      []ByteInterpreter
	P1Matcher       []ByteInterpreter
	P2Matcher       []ByteInterpreter
	P3Matcher       []ByteInterpreter
	DataInterpreter any
	LeMatcher       ByteInterpreter
}

func (aci *apduCommandInterpreter) MatchesCla(apdu apdu.ApduCommand) bool {
	for _, matcher := range aci.ClaMatcher {
		if matcher.Matches(apdu.Cla) {
			return true
		}
	}
	return false
}
func (aci *apduCommandInterpreter) MatchesIns(apdu apdu.ApduCommand) bool {
	for _, matcher := range aci.InsMatcher {
		if matcher.Matches(apdu.Ins) {
			return true
		}
	}
	return false
}
func (aci *apduCommandInterpreter) Matches(apdu *apdu.ApduCommand) bool {
	panic("not implemented") // TODO: Implement
}

func (aci *apduCommandInterpreter) Interpret(apdu *apdu.ApduCommand) (Interpretation, error) {
	panic("not implemented") // TODO: Implement
}
