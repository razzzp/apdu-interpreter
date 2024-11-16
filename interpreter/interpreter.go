package interpreter

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
