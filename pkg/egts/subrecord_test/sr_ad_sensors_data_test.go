package subrecord

import (
	"encoding/hex"
	"testing"

	"EGTS_PROTOCOL/pkg/egts/subrecord"
)

var (
	SRAdSensorsDataCheckIncome = []string{"000007000000000000000000"}
)

func TestSRAdSensorsDataDecoding(t *testing.T) {
	for i := range SRAdSensorsDataCheckIncome {
		pkgHex := SRAdSensorsDataCheckIncome[i]
		pkgBytes, err := hex.DecodeString(pkgHex)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		subr := subrecord.SRAdSensorsData{}
		err = subr.Decode(pkgBytes)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		hexed, err := subr.Encode()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if hex.EncodeToString(hexed) != SRAdSensorsDataCheckIncome[i] {
			t.Errorf("Have to be %s, but got %s", SRAdSensorsDataCheckIncome[i], hex.EncodeToString(hexed))
		}
	}
}
