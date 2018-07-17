package messages

import ()

// Type 1, 2, 3

// A Class A AIS unit broadcasts the following information every 2 to 10
// seconds while underway, and every 3 minutes while at anchor at a power level
// of 12.5 watts.
//
// https://www.navcen.uscg.gov/?pageName=AISMessagesA
type Position struct {
	// Valid types are 1, 2 or 3.
	Header     Header
	State      VesselState `bits:"38:4"`
	RateOfTurn int8        `bits:"42:8"`
	Speed      float32     `bits:"50:10" divisor:"10"`

	RawLocation struct {
		// If true, the location given is accurate to less than 10m
		Accuracy  bool    `bits:"60:1"`
		Longitude float32 `bits:"61:28" divisor:"600000"`
		Latitude  float32 `bits:"89:27" divisor:"600000"`
	}

	Course  uint16 `bits:"116:12" divisor:"10"`
	Heading uint16 `bits:"128:9"`

	Seconds uint8 `bits:"137:6"`

	// 0 = Default
	// 1 = No special maneuver
	// 2 = Special maneuver (such as regional passing arrangement)
	ManeuverIndicator uint8 `bits:"143:2"`

	// Spare
	// RAIM
	// Radio Status
}

func (p Position) Location() Location {
	return Location{
		Longitude: p.RawLocation.Longitude,
		Latitude:  p.RawLocation.Latitude,
	}
}
