package messages

//
type VesselState uint8

const (
	UnderWay VesselState = iota
	AtAnchor
	NotUnderCommand
	RestrictedManeuverability
	ConstrainedByDraught
	Moored
	Aground
	Fishing
	UnderWaySailing

	Undefined VesselState = VesselState(15)
)
