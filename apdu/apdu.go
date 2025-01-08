package apdu

type ApduCommand struct {
	Cla  byte
	Ins  byte
	P1   byte
	P2   byte
	P3   *byte
	Data []byte
	Le   *byte
}

func (ac *ApduCommand) AsBytes() []byte {
	result := make([]byte, 4)
	result[0] = ac.Cla
	result[1] = ac.Ins
	result[2] = ac.P1
	result[3] = ac.P2
	if ac.P3 != nil {
		result = append(result, *ac.P3)
	}
	if ac.Data != nil {
		result = append(result, ac.Data...)
	}
	if ac.Le != nil {
		result = append(result, *ac.Le)
	}
	return result
}

type ApduResponse struct {
	Data []byte
	SW1  byte
	SW2  byte
}

func (ar *ApduResponse) AsBytes() []byte {
	result := []byte{}
	result = append(result, ar.Data...)
	result = append(result, ar.SW1)
	result = append(result, ar.SW2)

	return result
}

type ApduCommandResponse struct {
	Command  *ApduCommand
	Response *ApduResponse
}
