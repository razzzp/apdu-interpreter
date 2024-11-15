package interpreter

import (
	"apdu-interpreter/schema"
	"apdu-interpreter/utils"
)

type InterpreterBuilder struct {
	eCollector utils.ErrorCollector
	result     *InterpreterEngine
}

type Schema struct {
	Name        string
	Description string
	Group       string
	Labels      []string
}

func (ib *InterpreterBuilder) BuildCommandIntp(commandDef *schema.CommandDefinition) (*ApduCommandInterpreter, error) {

	return nil, nil
}

func (ib *InterpreterBuilder) Build(schema schema.SchemaDefinition) *InterpreterEngine {
	ib.result = &InterpreterEngine{
		Schema: Schema{
			Name:        schema.Name,
			Group:       schema.Group,
			Description: schema.Description,
			Labels:      schema.Labels,
		},
	}
	for _, commandDef := range schema.CommandDefinitions {
		commandIntp, err := ib.BuildCommandIntp(&commandDef)
		if err != nil {
			ib.eCollector.AppendError(err.Error())
			continue
		}
		ib.result.CommandInterpreters = append(ib.result.CommandInterpreters, *commandIntp)

	}
	return ib.result
}
