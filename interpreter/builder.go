package interpreter

import (
	"log"

	"github.com/razzzp/apdu-interpreter/schema"
)

type InterpreterEngineBuilder struct {
	result *InterpreterEngine
}

type Schema struct {
	Name        string
	Description string
	Group       string
	Labels      []string
}

func appendByteInterpreter(def schema.ByteDefinition, interpreters []ByteInterpreter) ([]ByteInterpreter, error) {
	if def.SingleBit != nil {
		intp, err := SingleBitDefinition(def.SingleBit.BitNumber, def.SingleBit.ZeroIsOn, def.SingleBit.Description)
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

	if def.BytePattern != nil {
		intp, err := BytePattern(def.BytePattern.Pattern, def.BytePattern.Description)
		if err != nil {
			return nil, err
		}
		return append(interpreters, intp), nil
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

func (ib *InterpreterEngineBuilder) BuildCommandInterpreter(def *schema.CommandDefinition) (*apduCommandInterpreter, error) {
	apduIntp := apduCommandInterpreter{}

	err := buildByteIntpsToList(def.Cla, &apduIntp.ClaMatcher)
	if err != nil {
		return nil, err
	}

	err = buildByteIntpsToList(def.Ins, &apduIntp.InsMatcher)
	if err != nil {
		return nil, err
	}

	err = buildByteIntpsToList(def.P1, &apduIntp.P1Matcher)
	if err != nil {
		return nil, err
	}
	err = buildByteIntpsToList(def.P2, &apduIntp.P2Matcher)
	if err != nil {
		return nil, err
	}
	err = buildByteIntpsToList(def.P3, &apduIntp.P3Matcher)
	if err != nil {
		return nil, err
	}

	return &apduIntp, nil
}

func (ib *InterpreterEngineBuilder) BuildSchema(schemaDef *schema.SchemaDefinition) *InterpreterEngine {
	if ib.result == nil {
		ib.result = &InterpreterEngine{
			Schemas: []*schema.SchemaDefinition{},
		}
	}

	ib.result.Schemas = append(ib.result.Schemas, schemaDef)

	for _, specDef := range schemaDef.Spec {
		commandIntp, err := ib.BuildCommandInterpreter(&specDef.Command)
		if err != nil {
			log.Printf("[Error] Failed to build interpreter for command spec `%s`: %v", specDef.Name, err)
			continue
		}
		apduIntp := ApduInterpreter{
			SchemaDef:          schemaDef,
			CommandResponseDef: &specDef,
			CommandInterpreter: commandIntp,
		}
		ib.result.ApduInterpreters = append(ib.result.ApduInterpreters, &apduIntp)

	}
	return ib.result
}
