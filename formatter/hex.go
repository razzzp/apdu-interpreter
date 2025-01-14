package formatter

const (
	hextable = "0123456789abcdef"
)

// copied and modified from:
//
//	https://cs.opensource.google/go/go/+/refs/tags/go1.23.4:src/encoding/hex/hex.go
func EncodeWithSpace(dst, src []byte, spaceRune rune) int {
	j := 0
	for i, v := range src {
		dst[j] = hextable[v>>4]
		dst[j+1] = hextable[v&0x0f]
		// don't add last space
		if i != len(src)-1 {
			dst[j+2] = byte(spaceRune)
		}
		j += 3
	}
	return len(src) * 3
}

func EncodeStringWithSpace(src []byte) string {

	dst := make([]byte, len(src)*3-1)
	_ = EncodeWithSpace(dst, src, ' ')
	return string(dst)
}
