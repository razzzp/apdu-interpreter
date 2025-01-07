package apdu

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"unicode"
)

const HEX_CHARS = "0123456789abcdefABCDEF"
const WHITESPACE_CHARS = " \t\r"
const (
	TOKEN_HEX     = iota
	TOKEN_NEWLINE = iota
)

// not used
const (
	STATE_CLA   = iota
	STATE_INS   = iota
	STATE_P1    = iota
	STATE_P2    = iota
	STATE_P3    = iota
	STATE_CDATA = iota
	STATE_LE    = iota
	STATE_RDATA = iota
	STATE_SW1   = iota
	STATE_SW2   = iota
)

type ApduParser interface {
	GetNextCommandResponse() (*ApduCommandResponse, error)
}

type token struct {
	TokenType uint
	Value     byte
}

type apduLogParser struct {
	reader io.RuneReader
	pos    uint
	line   uint
	col    uint
}

func NewApduLogParser(reader io.RuneReader) apduLogParser {
	return apduLogParser{
		reader: reader,
		line:   1,
		col:    1,
	}
}

func (alp *apduLogParser) readRune() (rune, int, error) {
	alp.pos++
	alp.col++
	return alp.reader.ReadRune()
}

func isHexadecimalChar(r rune) bool {
	if unicode.Is(unicode.Hex_Digit, r) {
		return true
	} else {
		return false
	}
}

func isWhitespaceChar(r rune) bool {
	return unicode.IsSpace(r)
}

func isNewline(r rune) bool {
	if r == '\n' {
		return true
	} else {
		return false
	}
}

func (alp *apduLogParser) ReadNextToken() (*token, error) {
	hexArr := []rune{}
	for {
		// read next char
		r, _, err := alp.readRune()
		if err != nil {
			return nil, err
		}
		if isHexadecimalChar(r) {
			hexArr = append(hexArr, r)
			if len(hexArr) == 2 {
				// already formed hex
				hexAsByte, err := hex.DecodeString(string(hexArr))
				if err != nil {
					return nil, fmt.Errorf("invalid hex: %s", err.Error())
				}
				return &token{
					TokenType: TOKEN_HEX,
					Value:     hexAsByte[0],
				}, nil
			}
		} else if isNewline(r) {
			// err odd numbered chars, warn but continue
			if len(hexArr) == 1 {
				log.Printf("Odd numbered chars at line: %d, col: %d", alp.line, alp.col-1)
			}
			// inc line num and reset col
			alp.line++
			alp.col = 1
			return &token{
				TokenType: TOKEN_NEWLINE,
			}, nil
		} else if isWhitespaceChar(r) {
			//skip
		} else {
			// error
			return nil, fmt.Errorf("unknown char '%c' at line: %d, col: %d", r, alp.line, alp.col-1)
		}
	}
}

func (alp *apduLogParser) ReadLine() (result []byte, err error) {
	result = []byte{}
	for {
		tok, err := alp.ReadNextToken()
		if err != nil {
			if err.Error() == "EOF" && len(result) != 0 {
				return result, nil
			} else {
				return nil, err
			}
		}

		if tok.TokenType == TOKEN_HEX {
			result = append(result, tok.Value)
		} else if tok.TokenType == TOKEN_NEWLINE {
			// if blank line continue next line
			if len(result) == 0 {
				continue
			}
			return result, nil
		} else {
			log.Printf("Unknown token type %v", tok)
		}
	}
}

func (alp *apduLogParser) GetNextCommandResponse() (*ApduCommandResponse, error) {
	result := ApduCommandResponse{
		Command: &ApduCommand{},
	}

	// read command

	lineArr, err := alp.ReadLine()
	if err != nil {
		if err.Error() != "EOF" {
			return nil, fmt.Errorf("line %d: failed to read command: %s", alp.line-1, err.Error())
		} else {
			return nil, err
		}
	}

	lineLength := len(lineArr)
	if lineLength < 4 {
		return nil, fmt.Errorf("line %d: invalid apdu command, less than 4 bytes", alp.line-1)
	}

	// set command
	result.Command.Cla = lineArr[0]
	result.Command.Ins = lineArr[1]
	result.Command.P1 = lineArr[2]
	result.Command.P2 = lineArr[3]
	if lineLength > 4 {
		result.Command.P3 = &lineArr[4]
		// check data and length
		if int(*result.Command.P3) != 0 {
			if lineLength != 5+int(*result.Command.P3) && lineLength != 6+int(*result.Command.P3) {
				log.Printf("Warning: line %d: mismatched P3 and data length, expected: %d, found: %d", alp.line-1, int(*result.Command.P3), lineLength-5)
				result.Command.Data = lineArr[5:]
			} else {
				result.Command.Data = lineArr[5 : 5+int(*result.Command.P3)]
			}
		}

		// le
		if lineLength > 5+int(*result.Command.P3) {
			result.Command.Le = &lineArr[5+int(*result.Command.P3)]
		}
	}

	// read response
	lineArr, err = alp.ReadLine()
	if err != nil {
		log.Printf("Warning: line %d: failed to read response line: %s", alp.line-1, err.Error())
		return &result, nil
	}

	lineLength = len(lineArr)
	if lineLength >= 2 {
		result.Response = &ApduResponse{}
		if lineLength == 2 {
			result.Response.SW1 = lineArr[0]
			result.Response.SW2 = lineArr[1]
			return &result, nil
		} else {
			result.Response.Data = lineArr[:lineLength-2]
			result.Response.SW1 = lineArr[lineLength-2]
			result.Response.SW2 = lineArr[lineLength-1]
			return &result, nil
		}
	} else {
		log.Printf("Warning: line %d: Invalid response, less than 2 bytes", alp.line-1)
		return &result, nil
	}
}
