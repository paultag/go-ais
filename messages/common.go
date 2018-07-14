package messages

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
