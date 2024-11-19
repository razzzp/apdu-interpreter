package schema

type CommandDefinition struct {
	Name        string
	Decsription string
	Cla         []ByteDefinition
	Ins         []ByteDefinition
	P1          []ByteDefinition
	P2          []ByteDefinition
	P3          []ByteDefinition
	Data        string
	Le          string
}

type CommonDefition struct {
	Description string
}

type ByteDefinition struct {
	BitPattern          *BitPatternDefinition
	BytePattern         *BytePatternDefinition
	SingleBitDefinition *SingleBitDefinition
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
type BytePatternDefinition struct {
	Description string
	Pattern     string
}

type SchemaDefinition struct {
	Name               string
	Group              string
	Version            string
	Description        string
	Labels             []string
	CommandDefinitions []CommandDefinition `yaml:"commandDefinitions"`
}
