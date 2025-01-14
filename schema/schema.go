package schema

type CommandDefinition struct {
	Cla  []ByteDefinition
	Ins  []ByteDefinition
	P1   []ByteDefinition
	P2   []ByteDefinition
	P3   []ByteDefinition
	Data []DataDefinition
	Le   string
}

type ResponseDefinition struct {
	Description string
	Sw1         string
	Sw2         string
	// TODO
	// Data        any
}

type CommandResponseDefinition struct {
	Name        string
	Description string
	Command     CommandDefinition
	Response    *ResponseDefinition
}

type ByteDefinition struct {
	BitPattern   *BitPatternDefinition   `yaml:"bitPattern"`
	BytePatterns *BytePatternsDefinition `yaml:"bytePatterns"`
	BytePattern  *BytePatternDefinition  `yaml:"bytePattern"`
	SingleBit    *SingleBitDefinition    `yaml:"singleBit"`
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

type BytePatternDefinition struct {
	Description string
	Pattern     string
}

type SchemaDefinition struct {
	Name        string
	Group       string
	Version     string
	Description string
	Labels      []string
	Spec        []CommandResponseDefinition
	Common      CommonDefinitions
}

type CommonDefinitions struct {
	Responses []ResponseDefinition
}

type DataDefinition struct {
	When *GroupDefinition
}

type CriteriaDefinition struct {
	Cla []ByteDefinition
	Ins []ByteDefinition
	P1  []ByteDefinition
	P2  []ByteDefinition
}

type GroupDefinition struct {
	Label        string
	Criteria     CriteriaDefinition
	Interpreters []ByteArrayDefinition
}

type ByteArrayDefinition struct {
	LengthValue *LengthValueDefinition `yaml:"lengthValue"`
}

type LengthValueDefinition struct {
	Label string
}
