package messages

type Header struct {
	Type   uint8  `bits:"0:6"`
	Repeat uint8  `bits:"6:2"`
	MMSI   uint32 `bits:"8:30"`
}

// Any type that has a latitude or longitude.
type Locatable interface {
	// Return a standard Location struct from the raw representation
	// of a Location from the message itself.
	//
	// It's almost certenly better to access a location through this
	// message rather than using RawLocation. RawLocation is needed
	// to properly set the struct tags for those fields.
	Location() Location
}

// Common Location type.
type Location struct {
	Longitude float32
	Latitude  float32
	Fix       Fix
}
