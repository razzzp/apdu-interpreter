package interpreter

import "github.com/razzzp/apdu-interpreter/apdu"

// used to match a command for data interpretation
type CriteriaMatcher struct {
	ClaMatchers []ByteInterpreter
	// ins shouldn't be needed since only one possible for a command
	P1Matchers []ByteInterpreter
	P2Matchers []ByteInterpreter
}

func byteMatches(b byte, bMatchers []ByteInterpreter) bool {
	if len(bMatchers) == 0 {
		return true
	}

	for _, m := range bMatchers {
		if m.Matches(b) {
			return true
		}
	}

	return false
}

func (cm *CriteriaMatcher) Matches(apdu *apdu.ApduCommand) bool {
	result := byteMatches(apdu.Cla, cm.ClaMatchers)
	result = result && byteMatches(apdu.P1, cm.P1Matchers)
	result = result && byteMatches(apdu.P2, cm.P2Matchers)
	return result
}

func (cm *CriteriaMatcher) Interpret(apdu *apdu.ApduCommand) (*CommandInterpretation, error) {
	panic("not implemented") // TODO: Implement
}
