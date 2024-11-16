package interpreter

import (
	"errors"

	"github.com/razzzp/apdu-interpreter/schema"
	"github.com/razzzp/apdu-interpreter/utils"
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

func BuildByteIntp(def schema.ByteDefinition) (ByteIntp, error) {
	if def.BitPattern != nil {
		return BitPattern(def.BitPattern.Pattern, def.BitPattern.Description)
	}
	return nil, errors.New("unknown definition")
}

func (ib *InterpreterBuilder) BuildCommandIntp(commandDef *schema.CommandDefinition) (*ApduCommandInterpreter, error) {
	apduIntp := ApduCommandInterpreter{
		Name:        commandDef.Name,
		Description: commandDef.Decsription,
	}
	for _, def := range commandDef.Cla {
		intp, err := BuildByteIntp(def)
		if err != nil {
			return nil, err
		}
		apduIntp.ClaMatcher = append(apduIntp.ClaMatcher, intp)
	}

	return &apduIntp, nil
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
