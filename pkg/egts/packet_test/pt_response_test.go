package packet_test

import (
	"encoding/hex"
	"log"
	"testing"

	"EGTS_PROTOCOL/pkg/egts/packet"
)

var (
	PTResponseCheckIncome = []string{"00000006000000580101000300000000"}
)

func TestPTResponseDecoding(t *testing.T) {
	for i := range PTResponseCheckIncome {
		pkgHex := PTResponseCheckIncome[i]
		pkgBytes, err := hex.DecodeString(pkgHex)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		subr := packet.PTResponse{}
		err = subr.Decode(pkgBytes)
		if err != nil {
			log.Println(pkgBytes)
			t.Errorf("Error: %s", err.Error())
		}
		hexed, err := subr.Encode()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if hex.EncodeToString(hexed) != PTResponseCheckIncome[i] {
			t.Errorf("Have to be %s, but got %s", PTResponseCheckIncome[i], hex.EncodeToString(hexed))
		}
	}
}
