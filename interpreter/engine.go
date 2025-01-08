package interpreter

import (
	"log"

	"github.com/razzzp/apdu-interpreter/apdu"
	"github.com/razzzp/apdu-interpreter/schema"
)

type InterpreterEngine struct {
	Schemas                    []*schema.SchemaDefinition
	ApduInterpreters           []*ApduInterpreter
	CommonResponseInterpreters []ResponseInterpreter
	Parser                     apdu.ApduParser
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
			CommandResponse: apdu,
		}
		result = append(result, interpretation)
		for _, interpreter := range ie.ApduInterpreters {
			// interpret command
			cmdIntp, err := interpreter.CommandInterpreter.Interpret(apdu.Command)
			if err != nil {
				log.Printf("Error: interpreting command: %v", err)
			}
			// doesn't match
			if cmdIntp == nil {
				// skip
				continue
			}

			// create new interpretation
			cmdRspInterpretation := &CommandResponseInterpretation{
				Interpreter: interpreter,
				CommandIntp: cmdIntp,
			}
			interpretation.Interpretations = append(interpretation.Interpretations, cmdRspInterpretation)

			// check response
			if apdu.Response == nil {
				break
			}
			// interpret response TODO
			if interpreter.ResponseInterpreter != nil && interpreter.ResponseInterpreter.Matches(apdu.Response) {
				respIntp, err := interpreter.ResponseInterpreter.Interpret(apdu.Response)
				if err != nil {
					log.Printf("Error: interpreting response: %v", err)
				}

				cmdRspInterpretation.ResponseIntp = respIntp

			}
			if cmdRspInterpretation.ResponseIntp == nil {
				// try common interpreters
				for _, cIntpr := range ie.CommonResponseInterpreters {
					if cIntpr.Matches(apdu.Response) {
						respIntp, err := cIntpr.Interpret(apdu.Response)
						if err != nil {
							log.Printf("Error: interpreting response: %v", err)
						}

						cmdRspInterpretation.ResponseIntp = respIntp

					}
				}
			}

			break
		}
	}
	return result
}
