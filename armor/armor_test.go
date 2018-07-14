package armor_test

import (
	"log"
	"strings"
	"testing"

	"pault.ag/go/ais/armor"
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

//
func TestArmor(t *testing.T) {
	compareTo := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05}
	data, err := armor.Decode("012345")
	isok(t, err)

	assert(t, len(data) == len(compareTo), "lengths don't match")

	for i, el := range compareTo {
		assert(t, data[i] == el, "decoded armor doesn't match")
	}
}

//
func TestRoundTrip(t *testing.T) {
	testString := "Hack the Planet!"
	testBytes := []byte(testString)

	armorString, err := armor.Encode(testBytes)
	isok(t, err)

	roundTripBytes, err := armor.Decode(armorString)
	isok(t, err)

	assert(
		t,
		strings.Compare(string(roundTripBytes), testString) == 0,
		"String got mangled during round trip through Encode/Decode",
	)

}
