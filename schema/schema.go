package schema

type CommandDefinition struct {
	Cla  string
	Ins  string
	P1   string
	P2   string
	P3   string
	Data string
	Le   string
}

type Schema struct {
	Name               string
	Description        string
	Labels             []string
	CommandDefinitions []CommandDefinition `yaml:"commandDefinitions"`
}
