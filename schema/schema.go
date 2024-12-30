package schema

type CommandDefinition struct {
	Cla  []ByteDefinition
	Ins  []ByteDefinition
	P1   []ByteDefinition
	P2   []ByteDefinition
	P3   []ByteDefinition
	Data string
	Le   string
}

type ResponseDefinition struct {
	// TODO
	Data any
	SW1  any
	SW2  any
}

type CommandResponseDefinition struct {
	Name        string
	Description string
	Command     CommandDefinition
	Response    *ResponseDefinition
}

type ByteDefinition struct {
	BitPattern          *BitPatternDefinition   `yaml:"bitPattern"`
	BytePatterns        *BytePatternsDefinition `yaml:"bytePatterns"`
	SingleBitDefinition *SingleBitDefinition    `yaml:"singleBit"`
}

type SingleBitDefinition struct {
	Description string
	BitNumber   int
	ZeroIsOn    bool
}

type BitPatternDefinition struct {
	Description string
	Pattern     string
}
type BytePatternsDefinition struct {
	Description string
	Patterns    []string
}

type SchemaDefinition struct {
	Name        string
	Group       string
	Version     string
	Description string
	Labels      []string
	Spec        []CommandResponseDefinition
}
