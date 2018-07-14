package sixbit_test

import (
	"log"
	"testing"

	"pault.ag/go/ais/armor"
	"pault.ag/go/ais/sixbit"
)

func isok(t *testing.T, err error) {
	if err != nil {
		log.Printf("Error! Error is not nil! - %s\n", err)
		t.FailNow()
	}
}

func notok(t *testing.T, err error) {
	if err == nil {
		log.Printf("Error! Error is nil!\n")
		t.FailNow()
	}
}

func assert(t *testing.T, expr bool, why string) {
	if !expr {
		log.Printf("Assertion failed: %s", why)
		t.FailNow()
	}
}

func TestPayloadDecode(t *testing.T) {
	sixbytes, err := armor.Decode("14eG;o@034o8sd<L9i:a;WF>062D")
	isok(t, err)

	slice, err := sixbit.Decode(sixbytes)
	isok(t, err)

	type_ := slice.Slice(0, 6).Uint()
	assert(t, type_ == 0x01, "message type isn't one")

	mmsi := slice.Slice(8, 30).Uint()
	assert(t, mmsi == 316001245, "decoded mmsi wrong")
}
