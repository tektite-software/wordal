package wordal

import (
	"encoding/json"
	"hash/crc32"
	"testing"
)

func TestDict(t *testing.T) {
	dictJSONData, _ := json.Marshal(Dict)
	dictChecksum := crc32.ChecksumIEEE(dictJSONData)
	t.Logf("Dict checksum: %X (%d)\n", dictChecksum, dictChecksum)
	if dictChecksum != 3663095265 {
		t.Errorf("Dict checksum should be DA5665E1, but received %X", dictChecksum)
	}

	rdictJSONData, _ := json.Marshal(ReverseDict)
	rdictChecksum := crc32.ChecksumIEEE(rdictJSONData)
	t.Logf("ReverseDict checksum: %X (%d)\n", rdictChecksum, rdictChecksum)
	if rdictChecksum != 3250507868 {
		t.Errorf("ReverseDict checksum should be C1BED05C, but received %X", rdictChecksum)
	}
}
