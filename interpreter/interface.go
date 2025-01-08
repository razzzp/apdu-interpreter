package interpreter

import (
	"github.com/razzzp/apdu-interpreter/apdu"
	"github.com/razzzp/apdu-interpreter/schema"
)

type Interpretations interface {
	Add(intp any)
}

type ByteInterpreter interface {
	Matches(b byte) bool
	Interpret(i Interpretations, b byte) error
}

type CommandInterpreter interface {
	Matches(apdu *apdu.ApduCommand) bool
	Interpret(apdu *apdu.ApduCommand) (*CommandInterpretation, error)
}

type ReponseInterpreter interface {
	Matches(apdu *apdu.ApduCommand) bool
	Interpret(apdu *apdu.ApduCommand) (*CommandInterpretation, error)
}

type ApduInterpreter struct {
	SchemaDef           *schema.SchemaDefinition
	CommandResponseDef  *schema.CommandResponseDefinition
	CommandInterpreter  CommandInterpreter
	ResponseInterpreter ReponseInterpreter
}
