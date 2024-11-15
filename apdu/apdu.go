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
