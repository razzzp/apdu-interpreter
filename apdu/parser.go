package apdu

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/razzzp/apdu-interpreter/utils"
)

const HEX_CHARS = "0123456789abcdefABCDEF"
const WHITESPACE_CHARS = " \t\r"
const TOKEN_HEX = 0
const TOKEN_NEWLINE = 1

type ApduParser interface {
	GetNextCommandResponse() (*ApduCommandResponse, error)
}

type token struct {
	TokenType uint
	Value     string
}

type apduLogParser struct {
	reader bufio.Reader
	ec     utils.ErrorCollector
	pos    uint
	line   uint
	col    uint
}

func NewApduLogParser(reader bufio.Reader) apduLogParser {
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
	if strings.IndexRune(HEX_CHARS, r) > 0 {
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

func (alp *apduLogParser) readNextToken() (*token, error) {
	hex := []rune{}
	for {
		// read next char
		r, _, err := alp.readRune()
		if err != nil {
			return nil, err
		}
		if isHexadecimalChar(r) {
			hex = append(hex, r)
			if len(hex) == 2 {
				// already formed hex
				return &token{
					TokenType: TOKEN_HEX,
					Value:     string(r),
				}, nil
			}
		} else if isNewline(r) {
			// err odd numbered chars, warn but continue
			if len(hex) == 1 {
				log.Printf("Odd numbered chars at line: %d, col: %d", alp.line, alp.col)
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
			return nil, fmt.Errorf("unknown char '%c' at line: %d, col: %d", r, alp.line, alp.col)
		}
	}
}

func (alp *apduLogParser) GetNextCommandResponse() (*ApduCommandResponse, error) {
	for {
		tok, err := alp.readNextToken()
		if err != nil {
			return nil, err
		}
	}
}
