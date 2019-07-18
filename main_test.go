package wordal

import (
	"crypto/rand"
	"encoding/binary"
	"testing"
)

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

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Recovered in TestWordalSevenBytes; Panic received as expected.")
		}
	}()

	t.Logf(wordal.String())
	t.Error("Did not panic as expected")
}

func TestDecode(t *testing.T) {
	bytes := make([]byte, 2)
	decoded, err := Decode("aardvark")
	if err != nil {
		t.Error(err.Error())
	}
	if string(decoded) != string(bytes) {
		t.Errorf("Unable to decode 'aardvark'.  Expected %v, received %v", bytes, decoded)
	}
}

func TestFullCycle(t *testing.T) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	failIf(err, t)
	t.Logf("%v", bytes)

	wordal := Wordal(bytes)
	t.Logf(wordal.String())

	decoded, err := Decode(wordal.String())
	if err != nil {
		t.Error(err.Error())
	}
	if string(decoded) != string(bytes) {
		t.Errorf("Unable to decode '%v'.  Expected %v, received %v", wordal, bytes, decoded)
	}
}

func failIf(e error, t *testing.T) {
	if e != nil {
		t.Fail()
	}
}
