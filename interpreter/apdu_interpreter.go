package interpreter

import (
	"github.com/razzzp/apdu-interpreter/apdu"
)

type apduCommandInterpreter struct {
	ClaMatcher      []ByteInterpreter
	InsMatcher      []ByteInterpreter
	P1Matcher       []ByteInterpreter
	P2Matcher       []ByteInterpreter
	P3Matcher       []ByteInterpreter
	DataInterpreter any
	LeMatcher       ByteInterpreter
}

func (aci *apduCommandInterpreter) MatchesCla(apdu *apdu.ApduCommand) bool {
	for _, matcher := range aci.ClaMatcher {
		if matcher.Matches(apdu.Cla) {
			return true
		}
	}
	return false
}
func (aci *apduCommandInterpreter) MatchesIns(apdu *apdu.ApduCommand) bool {
	for _, matcher := range aci.InsMatcher {
		if matcher.Matches(apdu.Ins) {
			return true
		}
	}
	return false
}
func (aci *apduCommandInterpreter) Matches(apdu *apdu.ApduCommand) bool {
	return aci.MatchesCla(apdu) && aci.MatchesIns(apdu)
}

func (aci *apduCommandInterpreter) Interpret(apdu *apdu.ApduCommand) (*CommandInterpretation, error) {
	if !aci.Matches(apdu) {
		return nil, nil
	}

	result := NewCommandInterpretation()
	for _, matcher := range aci.ClaMatcher {
		err := matcher.Interpret(result.ClaIntp, apdu.Cla)
		if err != nil {
			return nil, err
		}
	}
	for _, matcher := range aci.InsMatcher {
		err := matcher.Interpret(result.InsIntp, apdu.Ins)
		if err != nil {
			return nil, err
		}
	}
	for _, matcher := range aci.P1Matcher {
		err := matcher.Interpret(result.P1Intp, apdu.P1)
		if err != nil {
			return nil, err
		}
	}
	for _, matcher := range aci.P2Matcher {
		err := matcher.Interpret(result.P2Intp, apdu.P2)
		if err != nil {
			return nil, err
		}
	}

	if apdu.P3 != nil {
		for _, matcher := range aci.P3Matcher {
			err := matcher.Interpret(result.P3Intp, *apdu.P3)
			if err != nil {
				return nil, err
			}
		}
	}

	return &result, nil
}
