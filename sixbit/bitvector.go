//
//
package sixbit

import (
	"strings"
)

// A bitslice is an array of bits. This allows us to slice and re-interpret
// the bits, regardless of how they got to us.
//
// For AIS, we get data in 6-bit bytes, but we need to slice them up without
// regard to any particular alignment.
//
// This is a memory hog (it's a slice of bools), and not very fast, but it's
// going to be a good place to start optimizing from, once it's all working.
type BitSlice struct {
	bits []bool
}

// Append a single bit to the bit slice.
func (bv *BitSlice) Append(bit bool) {
	bv.bits = append(bv.bits, bit)
}

// Slice the bit slice from bit offset `from`, for `length` bits. That slice
// is returned as another BitSlice, which can be interpreted from there.
func (bv *BitSlice) Slice(from, length uint) *BitSlice {
	return &BitSlice{
		bits: bv.bits[from : from+length],
	}
}

func (bv *BitSlice) StringTrim() string {
	return strings.TrimSpace(bv.String())
}

func (bv *BitSlice) String() string {
	ret := []byte{}
	for i := 0; i < len(bv.bits); i = i + 6 {
		nibble := bv.Slice(uint(i), 6).Uint()
		if nibble == 0 {
			/* is this right? are nulls in the string valid? */
			break
		}

		if nibble < 32 {
			nibble += 64
		}
		ret = append(ret, byte(nibble))
	}
	return string(ret)
}

// Read the bit slice as a uint. This assumes the 0th index is the MSB.
func (bv *BitSlice) Uint() uint {
	ret := uint(0)
	for i := range bv.bits {
		if bv.bits[i] {
			ret += (0x01 << uint((len(bv.bits)-1)-i))
		}
	}
	return ret
}

func (bv *BitSlice) Bool() bool {
	return bv.Uint() != 0
}

func (bv *BitSlice) Uint64() uint64 {
	return uint64(bv.Uint())
}

func (bv *BitSlice) Int64() int64 {
	bits := uint(len(bv.bits))
	x := int64(bv.Uint64())

	if x >= (0x01 << (bits - 1)) {
		x -= (0x01 << bits)
	}
	return x
}

// Append a byte to the list, MSB first.
func (bv *BitSlice) AppendByte(b byte, length uint) {
	for i := int(length - 1); i >= 0; i-- {
		bv.Append((b & (0x01 << uint(i))) != 0)
	}
}

// Append a slice of bytes.
func (bv *BitSlice) AppendBytes(data []byte, length uint) {
	for _, el := range data {
		bv.AppendByte(el, length)
	}
}

// Allocate a new BitSlice.
func NewBitSlice() BitSlice {
	return BitSlice{bits: []bool{}}
}
