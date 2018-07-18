package messages

import (
	"fmt"
)

//
type Header struct {
	// Type of the message. This is a number between 0 and 31. This number
	// indicates what the next bits mean to the client.
	Type uint8 `bits:"0:6"`

	// This indicates a value similar to an IP TTL value. If this value is
	// non-zero, the sender requests that you decrement this value by one,
	// and re-send this using your equipment. This allows boats to "mesh"
	// and relay messages much further than any one boat would be able to.
	//
	// Most programs importing this library will be reading only, so in
	// practice this value is not used, but it may be useful for other
	// purposes (maybe deduplication?)
	Repeat uint8 `bits:"6:2"`

	// MMSI, or Maritime Mobile Service Identity, is the globally unique
	// identifier tied to a specific radio, which is usually tied 1-to-1
	// to a Boat.
	//
	// If you'd like to try and map this to a boat, you can check the
	// FCC Database (http://wireless2.fcc.gov/UlsApp/UlsSearch/searchShip.jsp),
	// however most IDs aren't in this database.
	MMSI uint32 `bits:"8:30"`
}

//
func (h Header) String() string {
	return fmt.Sprintf("type=%d mmsi=%d (repeat=%d)", h.Type, h.MMSI, h.Repeat)
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
