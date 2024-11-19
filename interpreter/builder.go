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
	if def.SingleBitDefinition != nil {
		return SingleBitDefinition(def.SingleBitDefinition.BitNumber, def.SingleBitDefinition.ZeroIsOn, def.SingleBitDefinition.Description)
	}
	if def.BitPattern != nil {
		return BitPattern(def.BitPattern.Pattern, def.BitPattern.Description)
	}
	if def.BytePattern != nil {
		return BytePattern(def.BytePattern.Pattern, def.BytePattern.Description)
	}
	return nil, errors.New("unknown definition")
}

func buildByteIntpsToList(
	defs []schema.ByteDefinition,
	list *[]ByteIntp,
) error {
	for _, def := range defs {
		intp, err := BuildByteIntp(def)
		if err != nil {
			return err
		}
		newList := append(*list, intp)
		*list = newList
	}
	return nil
}

func (ib *InterpreterBuilder) BuildCommandIntp(commandDef *schema.CommandDefinition) (*ApduCommandInterpreter, error) {
	apduIntp := ApduCommandInterpreter{
		Name:        commandDef.Name,
		Description: commandDef.Decsription,
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
		commandIntp, err := ib.BuildCommandIntp(&commandDef)
		if err != nil {
			ib.eCollector.AppendError(err.Error())
			continue
		}
		ib.result.CommandInterpreters = append(ib.result.CommandInterpreters, *commandIntp)

	}
	return ib.result
}
