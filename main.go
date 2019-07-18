package wordal

import (
	"encoding/binary"
	"strings"
)

// Wordal type
type Wordal []byte

func (w Wordal) String() string {
	var results []string
	var indices []uint16

	if len(w)%2 == 1 {
		panic("Wordal contains uneven number of bytes!  Accepts byte pairs only.")
	}

	for remaining := len(w); remaining > 0; {
		i := len(w) - remaining
		indices = append(indices, binary.BigEndian.Uint16(w[i:i+2]))
		remaining -= 2
	}

	for i := range indices {
		wordal := Dict[indices[i]]
		results = append(results, wordal)
	}

	return strings.Join(results, " ")
}

// Decode takes a string and looks up the words in the wordal dictionary
func Decode(input string) ([]byte, error) {
	var bytes []byte

	splitStrings := strings.Split(strings.ToLower(input), " ")

	for s := range splitStrings {
		bits, ok := ReverseDict[splitStrings[s]]
		if !ok {
			return make([]byte, 0), &InvalidWordalError{splitStrings[s]}
		}
		b := make([]byte, 2)
		binary.BigEndian.PutUint16(b, bits)
		bytes = append(bytes, b[0])
		bytes = append(bytes, b[1])
	}

	return bytes, nil
}

// InvalidWordalError represents failure case on Decode
type InvalidWordalError struct {
	String string
}

func (e *InvalidWordalError) Error() string {
	return strings.Join([]string{"The following string could not be looked up in the wordal dict:", e.String}, " ")
}
