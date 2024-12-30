package interpreter

import (
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

func appendByteInterpreter(def schema.ByteDefinition, interpreters []ByteInterpreter) ([]ByteInterpreter, error) {
	if def.SingleBitDefinition != nil {
		intp, err := SingleBitDefinition(def.SingleBitDefinition.BitNumber, def.SingleBitDefinition.ZeroIsOn, def.SingleBitDefinition.Description)
		if err != nil {
			return nil, err
		}
		return append(interpreters, intp), nil
	}
	if def.BitPattern != nil {
		intp, err := BitPattern(def.BitPattern.Pattern, def.BitPattern.Description)
		if err != nil {
			return nil, err
		}
		return append(interpreters, intp), nil
	}
	if def.BytePatterns != nil {
		intps, err := BytePatterns(def.BytePatterns.Patterns, def.BytePatterns.Description)
		if err != nil {
			return nil, err
		}
		return append(interpreters, intps...), nil
	}
	return interpreters, nil
}

func buildByteIntpsToList(
	defs []schema.ByteDefinition,
	list *[]ByteInterpreter,
) error {
	for _, def := range defs {
		newList, err := appendByteInterpreter(def, *list)
		if err != nil {
			return err
		}
		*list = newList
	}
	return nil
}

func (ib *InterpreterBuilder) BuildCommandInterpreter(commandDef *schema.CommandDefinition) (*ApduCommandInterpreter, error) {
	apduIntp := ApduCommandInterpreter{
		Name:        commandDef.Name,
		Description: commandDef.Description,
	}
	err := buildByteIntpsToList(commandDef.Cla, &apduIntp.ClaMatcher)
	if err != nil {
		return nil, err
	}
	err = buildByteIntpsToList(commandDef.Ins, &apduIntp.InsMatcher)
	if err != nil {
		return nil, err
	}
	err = buildByteIntpsToList(commandDef.P1, &apduIntp.P1Matcher)
	if err != nil {
		return nil, err
	}
	err = buildByteIntpsToList(commandDef.P2, &apduIntp.P2Matcher)
	if err != nil {
		return nil, err
	}
	err = buildByteIntpsToList(commandDef.P3, &apduIntp.P3Matcher)
	if err != nil {
		return nil, err
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
		commandIntp, err := ib.BuildCommandInterpreter(&commandDef)
		if err != nil {
			ib.eCollector.AppendError(err.Error())
			continue
		}
		ib.result.CommandInterpreters = append(ib.result.CommandInterpreters, *commandIntp)

	}
	return ib.result
}
