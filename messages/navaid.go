package messages

// Aid to Navigation reports may mark things like buoys or lighthouses.
//
// This message is used by an aids to navigation (AtoN) AIS station. It is
// generally transmitted autonomously at a rate of once every three minutes and
// should not occupy more than two slots.
//
// https://www.navcen.uscg.gov/?pageName=AISMessage21
type NavigationAid struct {
	Type   uint8  `bits:"0:6"`
	Repeat uint8  `bits:"6:2"`
	MMSI   uint32 `bits:"8:30"`

	AidType AidType `bits:"38:5"`
	Name    string  `bits:"43:120"`

	RawLocation struct {
		// If true, the location given is accurate to less than 10m
		Accuracy  bool    `bits:"163:1"`
		Longitude float32 `bits:"164:28" divisor:"600000"`
		Latitude  float32 `bits:"192:27" divisor:"600000"`

		// From later in the stream

		Fix Fix `bits:"249:4"`
	}

	Geometry struct {
		Bow       uint16 `bits:"219:9"`
		Stern     uint16 `bits:"228:9"`
		Port      uint8  `bits:"237:6"`
		Starboard uint8  `bits:"243:6"`
	}

	Seconds uint8 `bits:"253:6"`

	OffPosition bool `bits:"259:1"`

	// Reserved
	// RAIM

	Virtual bool `bits:"269:1"`

	// AsignedModeFlag
	// Spare
	// Name xtn
}

func (na NavigationAid) Location() Location {
	return Location{
		Longitude: na.RawLocation.Longitude,
		Latitude:  na.RawLocation.Latitude,
		Fix:       na.RawLocation.Fix,
	}
}

type AidType uint8
