package wordal

import (
	"encoding/binary"
	"fmt"
	"strings"
)

// Wordal type
type Wordal []byte

func (w Wordal) String() string {
	var results []string
	var indices []uint16

	for remaining := len(w); remaining > 0; {
		i := len(w) - remaining
		switch remaining {
		case 0:
			break
		case 1:
			padded := []byte{0, w[i]}
			indices = append(indices, binary.BigEndian.Uint16(padded))
			remaining--
		default:
			indices = append(indices, binary.BigEndian.Uint16(w[i:i+2]))
			remaining -= 2
		}
	}

	for i := range indices {
		wordal := Dict[indices[i]]
		fmt.Printf("i: %v, wordal: %v\n", i, wordal)
		results = append(results, wordal)
	}

	return strings.Join(results, " ")
}
