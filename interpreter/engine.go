package interpreter

type InterpreterEngine struct {
	Schema           Schema
	ApduInterpreters []*ApduInterpreter
}
