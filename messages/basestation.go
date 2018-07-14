package messages

import (
	"time"
)

// Type 4

// A Base Station Report is used for reporting UTC time and date and, at the
// same time, position.
//
// It is also is used by AIS stations for determining if
// it is within 120 NM for response to Messages 20 and 23.
//
// https://www.navcen.uscg.gov/?pageName=AIS_Base_Station_Report
type BaseStation struct {
	Type   uint8  `bits:"0:6"`
	Repeat uint8  `bits:"6:2"`
	MMSI   uint32 `bits:"8:30"`

	RawTime struct {
		Year   uint16 `bits:"38:14"`
		Month  uint8  `bits:"52:4"`
		Day    uint8  `bits:"56:5"`
		Hour   uint8  `bits:"61:5"`
		Minute uint8  `bits:"66:6"`
		Second uint8  `bits:"72:6"`
	}

	RawLocation struct {
		// If true, the location given is accurate to less than 10m
		Accuracy  bool    `bits:"78:1"`
		Longitude float32 `bits:"79:28" divisor:"600000"`
		Latitude  float32 `bits:"107:27" divisor:"600000"`
		Fix       Fix     `bits:"134:4"`
	}

	// Spare
	// RAIM
	// SOTDMA
}

func (bs BaseStation) Time() time.Time {
	return time.Date(
		int(bs.RawTime.Year),
		time.Month(bs.RawTime.Month),
		int(bs.RawTime.Day),
		int(bs.RawTime.Hour),
		int(bs.RawTime.Minute),
		int(bs.RawTime.Second),
		0,
		time.UTC,
	)
}

func (bs BaseStation) Location() Location {
	return Location{
		Longitude: bs.RawLocation.Longitude,
		Latitude:  bs.RawLocation.Latitude,
		Fix:       bs.RawLocation.Fix,
	}
}
