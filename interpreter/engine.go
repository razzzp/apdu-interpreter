package interpreter

type InterpreterEngine struct {
	Schema              Schema
	CommandInterpreters []ApduCommandInterpreter
}
