package interpreter

type InterpreterEngine struct {
	Schema              Schema
	CommandInterpreters []ApduCommandInterpreter
}
type Interpretation struct {
}
