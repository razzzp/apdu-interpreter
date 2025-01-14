package bytearray

import (
	"fmt"

	"github.com/razzzp/apdu-interpreter/formatter"
	"github.com/razzzp/apdu-interpreter/interpreter"
)

// length-value interpreter
type LvInterpreter struct {
	Label string
}

func (li *LvInterpreter) Interpret(i interpreter.Interpretations, b []byte, startIdx int) (int, error) {
	// check out of bounds
	if len(b) <= startIdx {
		return startIdx, fmt.Errorf("start index out of bounds idx=%d length=%d", startIdx, len(b))
	}

	// get length bytes
	length := b[startIdx]
	startIdx++
	endIdx := min(len(b), startIdx+int(length))
	i.Add(fmt.Sprintf("%s: %s", li.Label, formatter.EncodeStringWithSpace(b[startIdx:endIdx])))
	return endIdx, nil
}
