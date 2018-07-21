package messages

import (
	"pault.ag/go/ais/sixbit"
)

// Equipment that supports Message 24 part A shall transmit once every 6 min
// alternating between channels.
//
// Message 24 Part A may be used by any AIS station to associate a MMSI with a name.
//
// Message 24 Part A and Part B should be transmitted once every 6 min by Class
// B “CS” and Class B “SO” shipborne mobile equipment. The message consists of
// two parts. Message 24B should be transmitted within 1 min following Message
// 24A.
//
// When the parameter value of dimension of ship/reference for position or type
// of electronic position fixing device is changed, Class-B :CS” and Class-B
// “SO” should transmit Message 24B.
//
// When requesting the transmission of a Message 24 from a Class B “CS” or
// Class B “SO”, the AIS station should respond with part A and part B.
//
// When requesting the transmission of a Message 24 from a Class A, the AIS
// station should respond with part B, which may contain the vendor ID only.
//
// https://www.navcen.uscg.gov/?pageName=AISMessagesB
type StaticData struct {
	// Always Type 24.
	Header Header

	PartNumber uint8 `bits:"38:2"`

	PartA StaticDataPartA
	PartB StaticDataPartB
}

func (sd StaticData) GetHeader() Header {
	return sd.Header
}

type StaticDataPartA struct {
	Name string `bits:"40:120"`
}

type StaticDataPartB struct {
	ShipType      ShipType `bits:"40:8"`
	VendorID      string   `bits:"48:18"`
	UnitModelCode uint8    `bits:"66:4"`
	Serial        uint32   `bits:"70:20"`
	Callsign      string   `bits:"90:42"`

	Geometry struct {
		Bow       uint16 `bits:"132:9"`
		Stern     uint16 `bits:"141:9"`
		Port      uint8  `bits:"150:6"`
		Starboard uint8  `bits:"156:6"`
	}

	MothershipMMSI uint32 `bits:"132:30"`
}

/* Sweet jesus take the wheel on this one.
 *
 * In order to deal with the fact that my unmarshaling code is super not
 * smart at all, I've had to get around the fact that when I have a "part a"
 * and "part b" that are made up of the same bits, but the meaning and slicing
 * changes based on a value of a field.
 *
 * One (maybe better) option is to define a way to make Unmarshal avoid
 * recursing into nested structs, and force unpacking. It would also be handy
 * if I could Unmarshel into an Unmarshallable type without having it call the
 * function I usually want to Unmarshal in.
 *
 * As a result, I'm just defining a new type that will avoid the Unmarshallable
 * interface, so I can unpack the values, then I can just cast it over to
 * the real one. */
type _partA StaticDataPartA
type _partB StaticDataPartB

func (p *StaticDataPartA) UnmarshalBits(bits *sixbit.BitSlice) error {
	// Also known as the StaticData.PartNumber
	if bits.Slice(38, 2).Uint() != 0 {
		return nil
	}

	// Unmarshal into "myself" then cast to a public type, then dump it
	// over myself.
	ret := _partA{}
	if err := Unmarshal(bits, &ret); err != nil {
		return err
	}
	*p = StaticDataPartA(ret)
	return nil
}

//
func (p *StaticDataPartB) UnmarshalBits(bits *sixbit.BitSlice) error {
	// Also known as the StaticData.PartNumber
	if bits.Slice(38, 2).Uint() != 1 {
		return nil
	}

	// Unmarshal into "myself" then cast to a public type, then dump it
	// over myself.
	ret := _partB{}
	if err := Unmarshal(bits, &ret); err != nil {
		return err
	}
	*p = StaticDataPartB(ret)
	return nil
}

//

type ShipType uint8
