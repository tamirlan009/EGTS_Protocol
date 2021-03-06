package subrecord

import (
	"encoding/hex"
	"testing"

	"EGTS_PROTOCOL/pkg/egts/subrecord"
)

var (
	SRExPosDataRecordCheckIncome = []string{"18110000"}
)

func TestSRExPosDataRecordDecoding(t *testing.T) {
	for i := range SRExPosDataRecordCheckIncome {
		pkgHex := SRExPosDataRecordCheckIncome[i]
		pkgBytes, err := hex.DecodeString(pkgHex)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		subr := subrecord.SRExPosDataRecord{}
		err = subr.Decode(pkgBytes)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		hexed, err := subr.Encode()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if hex.EncodeToString(hexed) != SRExPosDataRecordCheckIncome[i] {
			t.Errorf("Have to be %s, but got %s", SRExPosDataRecordCheckIncome[i], hex.EncodeToString(hexed))
		}
	}
}
