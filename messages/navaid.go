package messages

// Aid to Navigation reports may mark things like buoys or lighthouses.
//
// This message is used by an aids to navigation (AtoN) AIS station. It is
// generally transmitted autonomously at a rate of once every three minutes and
// should not occupy more than two slots.
//
// https://www.navcen.uscg.gov/?pageName=AISMessage21
type NavigationAid struct {
	// Always Type 21
	Header Header

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

	RawGeometry struct {
		Bow       uint16 `bits:"219:9"`
		Stern     uint16 `bits:"228:9"`
		Port      uint8  `bits:"237:6"`
		Starboard uint8  `bits:"243:6"`
	}

	Seconds uint8 `bits:"253:6"`

	OffPosition bool `bits:"259:1"`

	// Reserved
	// RAIM

	// Is the Aid to Navigation broadcasting the AIS track itself, or is
	// there a virtual site (maybe ashore) transmiting the location of this
	// Aid.
	Virtual bool `bits:"269:1"`

	// AsignedModeFlag
	// Spare
	// Name xtn
}

func (na NavigationAid) GetGeometry() Geometry {
	return Geometry{
		Bow:       na.RawGeometry.Bow,
		Stern:     na.RawGeometry.Stern,
		Port:      na.RawGeometry.Port,
		Starboard: na.RawGeometry.Starboard,
	}
}

func (na NavigationAid) GetName() Name {
	return Name(na.Name)
}

func (na NavigationAid) GetHeader() Header {
	return na.Header
}

func (na NavigationAid) GetLocation() Location {
	return Location{
		Longitude: na.RawLocation.Longitude,
		Latitude:  na.RawLocation.Latitude,
		Fix:       na.RawLocation.Fix,
	}
}

type AidType uint8
