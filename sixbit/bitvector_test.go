package sixbit_test

import (
	"strings"
	"testing"

	"pault.ag/go/ais/sixbit"
)

func TestMMSI(t *testing.T) {
	bs := sixbit.NewBitSlice()
	sixByteString := []byte{
		0x1, 0x4, 0x2d, 0x17, 0x1c, 0xa, 0x10, 0x0, 0x0, 0x0, 0x37,
		0x8, 0x37, 0x21, 0x28, 0x1c, 0x1d, 0x32, 0x1f, 0x2b, 0x30,
		0x35, 0x17, 0x10, 0x0, 0x8, 0x18, 0x1b,
	}
	bs.AppendBytes(sixByteString, 6)
	val := bs.Slice(8, 30).Uint()
	assert(t, val == 316005417, "decoded MMSI is wrong")
}

func TestString(t *testing.T) {
	slice, err := sixbit.Decode([]byte("HELLO?!HELLO!?"))
	isok(t, err)

	assert(t, strings.Compare(slice.String(), "HELLO?!HELLO!?") == 0, "Round trip borked")
}
