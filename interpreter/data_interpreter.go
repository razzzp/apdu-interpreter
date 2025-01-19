package interpreter

import (
	"log"

	"github.com/razzzp/apdu-interpreter/apdu"
)

type dataInterpreter struct {
	Criteria     CommandInterpreter
	Interpreters []ByteArrayInterpreter
}

func (di *dataInterpreter) Interpret(apdu *apdu.ApduCommand) (*DataInterpretations, error) {
	result := DataInterpretations{}
	if di.Criteria.Matches(apdu) {

		curIdx := 0
		for _, intp := range di.Interpreters {
			var err error
			curIdx, err = intp.Interpret(&result, apdu.Data, curIdx)
			if err != nil {
				log.Printf("Error parsing data: %v", err)
			}
		}
		return &result, nil
	}
	return &result, nil
}
