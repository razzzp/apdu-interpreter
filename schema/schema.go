package schema

type CommandDefinition struct {
	Name        string
	Decsription string
	Cla         string
	Ins         string
	P1          string
	P2          string
	P3          string
	Data        string
	Le          string
}

type SchemaDefinition struct {
	Name               string
	Group              string
	Version            string
	Description        string
	Labels             []string
	CommandDefinitions []CommandDefinition `yaml:"commandDefinitions"`
}
