package units

var (
	Mass = UnitOptionQuantity("mass")

	// metric
	Gram        = NewUnit("gram", "g", Mass)
	ExaGram     = Exa(Gram)
	PetaGram    = Peta(Gram)
	TeraGram    = Tera(Gram)
	GigaGram    = Giga(Gram)
	MegaGram    = Mega(Gram)
	KiloGram    = Kilo(Gram)
	HectoGram   = Hecto(Gram)
	DecaGram    = Deca(Gram)
	DeciGram    = Deci(Gram)
	CentiGram   = Centi(Gram)
	MilliGram   = Milli(Gram)
	MicroGram   = Micro(Gram)
	NanoGram    = Nano(Gram)
	PicoGram    = Pico(Gram)
	FemtoGram   = Femto(Gram)
	AttoGram    = Atto(Gram)
	MetricTonne = NewUnit("metric tonne", "MT", UnitOptionAliases("metric ton", "tonne"), Mass, SI)

	// imperial
	Grain        = NewUnit("grain", "gr", Mass, BI)
	Drachm       = NewUnit("drachm", "dr", Mass, BI)
	Ounce        = NewUnit("ounce", "oz", Mass, BI)
	BritishPound = NewUnit("british pound", "uk lb", Mass, BI)
	Stone        = NewUnit("stone", "st", Mass, BI)
	LongTon      = NewUnit("long ton", "uk t", UnitOptionAliases("lt"), Mass, BI)
	Slug         = NewUnit("slug", "", Mass, BI)
	TroyOunce    = NewUnit("troy ounce", "ozt", UnitOptionAliases("oz t", "apoz", "troy oz", "fine troy ounce", "t. oz"), Mass, BI)

	// us
	Pound = NewUnit("pound", "lb", UnitOptionAliases("lbs", "pounds", "US pound", "US lb", "U.S. pound", "U.S. lb"), Mass, US)
	Cwt   = NewUnit("hundredweight", "cwt", UnitOptionAliases("US cwt", "U.S. cwt", "US hundredweight", "U.S. hundredweight", "short hundredweight", "sh cwt"), Mass, US)
	Ton   = NewUnit("ton", "t", UnitOptionAliases("tons", "US ton", "U.S. ton", "short ton", "sh tn", "sh ton"), Mass, US)

	// other
	Arroba  = NewUnit("arroba", "arroba", UnitOptionAliases("@", "arr"), Mass)
	Sack    = NewUnit("sack", "sc", Mass)
	Quintal = NewUnit("quintal", "q", UnitOptionAliases("quintals"), Mass)
)

func init() {

	// metric
	NewRatioConversion(MetricTonne, KiloGram, 1000.0)

	// imperial
	NewRatioConversion(Grain, Gram, 0.06479891)
	NewRatioConversion(Drachm, Gram, 1.7718451953125)
	NewRatioConversion(Ounce, Gram, 28.349523125)
	NewRatioConversion(BritishPound, Gram, 453.59237)
	NewRatioConversion(Stone, Gram, 6350.29318)
	NewRatioConversion(LongTon, Gram, 1016046.9088)
	NewRatioConversion(Slug, Gram, 14593.90294)
	NewRatioConversion(TroyOunce, Gram, 31.10352185)

	// us
	NewRatioConversion(Pound, Gram, 453.59237)
	NewRatioConversion(Cwt, Pound, 100.0)
	NewRatioConversion(Ton, Pound, 2000.0)

	// other
	NewRatioConversion(Arroba, KiloGram, 11.502)
	NewRatioConversion(Sack, KiloGram, 60.0)
	NewRatioConversion(Quintal, KiloGram, 100.0)
}
