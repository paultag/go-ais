package messages_test

import (
	"log"
	"testing"

	"pault.ag/go/ais/armor"
	"pault.ag/go/ais/messages"
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

func TestUnpack(t *testing.T) {
	sixbytes, err := armor.Decode("14eG;o@034o8sd<L9i:a;WF>062D")
	isok(t, err)

	slice, err := sixbit.Decode(sixbytes)
	isok(t, err)

	nav := messages.Position{}
	isok(t, messages.Unmarshal(slice, &nav))

	assert(t, nav.Header.Type == 0x01, "Unpack message type isn't 1")
	assert(t, nav.Header.MMSI == 316001245, "MMSI is jacked")

	var lon float32 = -123.877748
	var lat float32 = 49.200283

	assert(t, nav.RawLocation.Longitude == lon, "long is wrong")
	assert(t, nav.RawLocation.Latitude == lat, "lat is wrong")
	assert(t, nav.Speed == 19.6, "Speed over ground is wrong")
	assert(t, nav.Course == 235, "Course over ground is wrong")
	assert(t, nav.Heading == 235, "Heading sucks")
}
