package messages

// https://www.navcen.uscg.gov/?pageName=AISMessagesB
// Less detail than type 1-3
type ClassBPosition struct {
	// Always Type 18
	Header Header

	// Regional reserved

	Speed float32 `bits:"46:10" divisor:"10"`

	RawLocation struct {
		// If true, the location given is accurate to less than 10m
		Accuracy  bool    `bits:"56:1"`
		Longitude float32 `bits:"57:28" divisor:"600000"`
		Latitude  float32 `bits:"85:27" divisor:"600000"`
	}

	Course  uint16 `bits:"112:12" divisor:"10"`
	Heading uint16 `bits:"124:9"`

	Seconds uint8 `bits:"133:6"`

	// Regional reserved
	// CS Unit
	// Display flag
	// DSC Flag
	// Band flag
	// Message 22 flag
}

func (cbp ClassBPosition) GetHeader() Header {
	return cbp.Header
}

func (cbp ClassBPosition) GetLocation() Location {
	return Location{
		Longitude: cbp.RawLocation.Longitude,
		Latitude:  cbp.RawLocation.Latitude,
	}
}
