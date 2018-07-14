package messages

/* EPFD, or Electronic Position Fixing Device, is the generic name for a thing
 * like GPS, but not tied to a specific technology. This usaully defines the
 * technology that was used to get the lat/lon fix.
 *
 * We call this "fix" since it makes a lot more sense to a non-technical
 * library user. */
type Fix uint8

const (
	// Default option, undefined.
	UndefinedFix Fix = iota

	// GPS is a global radionavigation system run by the US Air Force.
	GPSFix

	// GLONASS is a Russian state-run company that provides a radionavigation
	// satellite constilation.
	GLONASSFix

	// Combined GPS and GLONASS.
	CombinedGPSGLONASSFix

	// A 1950's LF radionavigation system that uses fixed land based beacons
	// to aquire a location fix. It was used mostly by militaries at the time.
	LoranCFix

	// The Russian version of LORAN-C.
	ChaykaFix

	//
	IntegratedNavigationSystemFix

	// Position was Surveyed and hardcoded. This may be set for stuff like
	// beacons.
	SurveyedFix

	// Galileo is a EU run satellite constilation that provides radionavigation.
	// This is the EU version of GPS or GLONASS.
	GalileoFix

	// This is a bug, but a lot of units emit this rather than 0.
	UnknownFix Fix = Fix(15)
)
