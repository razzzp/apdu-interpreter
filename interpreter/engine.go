package interpreter

import (
	"log"

	"github.com/razzzp/apdu-interpreter/apdu"
	"github.com/razzzp/apdu-interpreter/schema"
)

type InterpreterEngine struct {
	Schemas          []*schema.SchemaDefinition
	ApduInterpreters []*ApduInterpreter
	Parser           apdu.ApduParser
}

func (ie *InterpreterEngine) Interpret() []*ApduInterpretation {
	result := []*ApduInterpretation{}
	for {
		apdu, err := ie.Parser.GetNextCommandResponse()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Printf("Error: reading command response pair: %v", err)
			continue
		}

		// create empty interpretation
		var interpretation *ApduInterpretation = &ApduInterpretation{
			Command: &CommandInterpretation{
				Command: apdu.Command,
			},
		}
		result = append(result, interpretation)
		for _, interpreter := range ie.ApduInterpreters {
			// interpret command
			cmdIntp, err := interpreter.CommandInterpreter.Interpret(apdu.Command)
			if err != nil {
				log.Printf("Error: reading interpreting command: %v", err)
			}
			// doesn't match
			if cmdIntp == nil {
				// skip
				continue
			}

			interpretation.ApduInterpreter = interpreter
			// replace empty interpretation with filled
			interpretation.Command = cmdIntp
			// check response
			if apdu.Response == nil {
				break
			}
			// interpret response TODO
			// respIntp, err := interpreter.ResponseInterpreter.Interpret
			break
		}
	}
	return result
}
