package apdu

type ApduParser interface {
	GetNextCommandResponse() (*ApduCommandResponse, error)
}
