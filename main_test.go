package wordal

import (
	"crypto/rand"
	"encoding/binary"
	"testing"
)

func TestWordalOneByte(t *testing.T) {
	bytes := make([]byte, 1)
	_, err := rand.Read(bytes)
	failIf(err, t)
	t.Logf("uint16: %s", bytes)

	wordal := Wordal(bytes)
	t.Logf(wordal.String())
}

func TestWordalTwoBytes(t *testing.T) {
	bytes := make([]byte, 2)
	_, err := rand.Read(bytes)
	failIf(err, t)
	t.Logf("uint16: %v", binary.BigEndian.Uint16(bytes))

	wordal := Wordal(bytes)
	t.Logf(wordal.String())
}

func TestWordalFourBytes(t *testing.T) {
	bytes := make([]byte, 4)
	_, err := rand.Read(bytes)
	failIf(err, t)
	t.Logf("uint16: %v", binary.BigEndian.Uint16(bytes))

	wordal := Wordal(bytes)
	t.Logf(wordal.String())
}

func TestWordalSevenBytes(t *testing.T) {
	bytes := make([]byte, 7)
	_, err := rand.Read(bytes)
	failIf(err, t)
	t.Logf("uint16: %v", binary.BigEndian.Uint16(bytes))

	wordal := Wordal(bytes)
	t.Logf(wordal.String())
}

func failIf(e error, t *testing.T) {
	if e != nil {
		t.Fail()
	}
}
