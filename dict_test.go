package wordal

import (
	"testing"
)

func TestDict(t *testing.T) {
	if Dict[55917] != "testing" {
		t.Error("Word 'testing' not at expected index in array.  Is Dict corrupt?")
	}
	if ReverseDict["testing"] != 55917 {
		t.Error("Unexpected value of 'testing' in map.  Is ReverseDict corrupt?")
	}
}
