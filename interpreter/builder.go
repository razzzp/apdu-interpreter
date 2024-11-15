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
	intp := ApduCommandInterpreter{
		Name:        commandDef.Name,
		Description: commandDef.Decsription,
	}
	claMatcher, err := NewByteMatcher(commandDef.Cla, "")
	if err != nil {
		return nil, err
	}
	intp.ClaMatcher = claMatcher

	insMatcher, err := NewByteMatcher(commandDef.Ins, "")
	if err != nil {
		return nil, err
	}
	intp.InsMatcher = insMatcher

	// p1Matcher, err := NewByteMatcher(commandDef.P1, "")
	// if err != nil {
	// 	return nil, err
	// }
	// intp.P1Matcher = p1Matcher

	// p2Matcher, err := NewByteMatcher(commandDef.P2, "")
	// if err != nil {
	// 	return nil, err
	// }
	// intp.P2Matcher = p2Matcher

	// p3Matcher, err := NewByteMatcher(commandDef.P3, "")
	// if err != nil {
	// 	return nil, err
	// }
	// intp.P3Matcher = p3Matcher
	return &intp, nil
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
