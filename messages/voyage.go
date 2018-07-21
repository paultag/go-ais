package messages

// In addition, the Class A AIS unit broadcasts the following information every
// 6 minutes. Should only be used by Class A shipborne and SAR aircraft AIS
// stations when reporting static or voyage related data
//
// https://www.navcen.uscg.gov/?pageName=AISMessagesAStatic
type Voyage struct {
	// Always Type 5.
	Header Header

	// AISVersion uint8    `bits:"38:2"`

	IMO      IMONumber `bits:"40:30"`
	Callsign string    `bits:"70:42"`
	Name     string    `bits:"112:120"`
	ShipType ShipType  `bits:"232:8"`

	Geometry struct {
		Bow       uint16 `bits:"240:9"`
		Stern     uint16 `bits:"249:9"`
		Port      uint8  `bits:"258:6"`
		Starboard uint8  `bits:"264:6"`
	}

	RawLocation struct {
		Fix Fix `bits:"270:4"`
	}

	ETA struct {
		Month  uint8 `bits:"274:4"`
		Day    uint8 `bits:"278:5"`
		Hour   uint8 `bits:"283:5"`
		Minute uint8 `bits:"288:6"`
	}

	Draught     float32 `bits:"294:8" divisor:"10"`
	Destination string  `bits:"302:120"`
}

func (v Voyage) GetHeader() Header {
	return v.Header
}

// TODO: Add in Checksum and tests for IMO Numbers.
type IMONumber uint32
