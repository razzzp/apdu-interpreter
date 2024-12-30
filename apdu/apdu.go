package apdu

type ApduCommand struct {
	Cla  byte
	Ins  byte
	P1   byte
	P2   byte
	P3   byte
	Data []byte
	Le   *byte
}

type ApduResponse struct {
	Data []byte
	SW1  byte
	SW2  byte
}

type ApduCommandResponse struct {
	Command  *ApduCommand
	Response *ApduResponse
}
