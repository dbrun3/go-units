package units

var (
	Energy = UnitOptionQuantity("energy")

	// metric
	Joule      = NewUnit("joule", "J", Energy)
	KiloJoule  = Kilo(Joule)
	MegaJoule  = Mega(Joule)
	GigaJoule  = Giga(Joule)
	TeraJoule  = Tera(Joule)
	PetaJoule  = Peta(Joule)
	ExaJoule   = Exa(Joule)
	ZettaJoule = Zetta(Joule)
	YottaJoule = Yotta(Joule)
	MilliJoule = Milli(Joule)
	MicroJoule = Micro(Joule)
	NanoJoule  = Nano(Joule)
	PicoJoule  = Pico(Joule)
	FemtoJoule = Femto(Joule)
	AttoJoule  = Atto(Joule)

	WattHour     = NewUnit("watt-hour", "Wh", Energy)
	KiloWattHour = Kilo(WattHour)
	GigaWattHour = Giga(WattHour)
	MegaWattHour = Mega(WattHour)
	TeraWattHour = Tera(WattHour)
	PetaWattHour = Peta(WattHour)

	// other
	ElectronVolt = NewUnit("electronvolt", "eV", Energy)
	Calorie      = NewUnit("calorie", "cal", Energy)
	BTU          = NewUnit("british thermal unit", "BTU", Energy)
	KiloBTU      = Kilo(BTU)
	MMBTU        = NewUnit("million british thermal unit", "MMBTU", Energy)
)

func init() {
	// Non-metric to Metric conversions
	NewRatioConversion(ElectronVolt, Joule, 1.60218e-19)
	NewRatioConversion(Calorie, Joule, 4.184)
	NewRatioConversion(WattHour, Joule, 3600)
	NewRatioConversion(KiloWattHour, Joule, 3600*1e3)
	NewRatioConversion(MegaWattHour, Joule, 3600*1e6)
	NewRatioConversion(GigaWattHour, Joule, 3600*1e9)
	NewRatioConversion(TeraWattHour, Joule, 3600*1e12)
	NewRatioConversion(PetaWattHour, Joule, 3600*1e15)

	NewRatioConversion(BTU, Joule, 1054.80)
	NewRatioConversion(MMBTU, BTU, 1e6)
}
