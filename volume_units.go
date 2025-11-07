package units

var (
	Volume = UnitOptionQuantity("volume")

	// metric
	Liter           = NewUnit("liter", "l", Volume, SI, UnitOptionAliases("litre"))
	ExaLiter        = Exa(Liter)
	PetaLiter       = Peta(Liter)
	TeraLiter       = Tera(Liter)
	GigaLiter       = Giga(Liter)
	MegaLiter       = Mega(Liter)
	KiloLiter       = Kilo(Liter)
	HectoLiter      = Hecto(Liter)
	DecaLiter       = Deca(Liter)
	DeciLiter       = Deci(Liter)
	CentiLiter      = Centi(Liter)
	MilliLiter      = Milli(Liter)
	MicroLiter      = Micro(Liter)
	NanoLiter       = Nano(Liter)
	PicoLiter       = Pico(Liter)
	FemtoLiter      = Femto(Liter)
	AttoLiter       = Atto(Liter)
	CubicCentimeter = Cube(CentiMeter, UnitOptionAliases("cc"))
	CubicMeter      = Cube(Meter)

	// imperial
	ImperialQuart      = NewUnit("imperial quart", "uk qt", UnitOptionAliases("UK quart", "UK qt", "imperial qt", "British quart"), Volume, BI)
	ImperialPint       = NewUnit("imperial pint", "uk pt", UnitOptionAliases("UK pint", "UK pt", "imperial pt", "British pint"), Volume, BI)
	ImperialGallon     = NewUnit("imperial gallon", "uk gal", UnitOptionAliases("UK gallon", "UK gal", "imperial gal", "British gallon"), Volume, BI)
	ImperialFluidOunce = NewUnit("imperial fluid ounce", "uk fl oz", UnitOptionAliases("UK fl oz", "UK fluid ounce", "imperial fl oz", "British fl oz"), Volume, BI)
	ImperialBushel     = NewUnit("imperial bushel", "uk bu", UnitOptionAliases("UK bushel", "UK bu", "imperial bu", "British bushel"), Volume, BI)
	ImperialPeck       = NewUnit("imperial peck", "uk pk", UnitOptionAliases("UK peck", "UK pk", "imperial pk", "British peck"), Volume, BI)

	// us
	Quart      = NewUnit("quart", "qt", UnitOptionAliases("quarts", "US quart", "U.S. quart", "US qt", "U.S. qt", "liquid quart", "fluid quart", "fl qt"), Volume, US)
	Pint       = NewUnit("pint", "pt", UnitOptionAliases("pints", "US pint", "U.S. pint", "US pt", "U.S. pt", "liquid pint", "fluid pint", "fl pt"), Volume, US)
	Gallon     = NewUnit("gallon", "gal", UnitOptionAliases("gallons", "U.S. gallon", "US gallon", "U.S. gal", "US gal", "fluid gallon", "fl gal", "liquid gallon"), Volume, US)
	FluidOunce = NewUnit("fluid ounce", "fl oz", UnitOptionAliases("floz", "fluid ounces", "US fl oz", "U.S. fl oz", "US fluid ounce", "U.S. fluid ounce"), Volume, US)
	Bushel     = NewUnit("bushel", "bu", UnitOptionAliases("bushels", "US bushel", "U.S. bushel", "US bu", "U.S. bu", "dry bushel"), Volume, US)
	Peck       = NewUnit("peck", "pk", UnitOptionAliases("pecks", "US peck", "U.S. peck", "US pk", "U.S. pk", "dry peck"), Volume, US)
	DryGallon  = NewUnit("dry gallon", "dry gal", UnitOptionAliases("dry gallons", "US dry gallon", "U.S. dry gallon", "US dry gal", "U.S. dry gal"), Volume, US)
	DryQuart   = NewUnit("dry quart", "dry qt", UnitOptionAliases("dry quarts", "US dry quart", "U.S. dry quart", "US dry qt", "U.S. dry qt"), Volume, US)
	DryPint    = NewUnit("dry pint", "dry pt", UnitOptionAliases("dry pints", "US dry pint", "U.S. dry pint", "US dry pt", "U.S. dry pt"), Volume, US)
	CubicInch  = Cube(Inch)
	CubicFoot  = Cube(Foot)
	Cord       = NewUnit("cord", "cd", Volume, US)
)

func init() {
	// metric
	NewRatioConversion(CubicCentimeter, MilliLiter, 1.0)
	NewRatioConversion(CubicMeter, Liter, 1000.0)

	// imperial
	NewRatioConversion(ImperialQuart, Liter, 1.1365225)
	NewRatioConversion(ImperialPint, Liter, 0.56826125)
	NewRatioConversion(ImperialGallon, Liter, 4.54609)
	NewRatioConversion(ImperialFluidOunce, MilliLiter, 28.4130625)
	NewRatioConversion(ImperialBushel, Liter, 36.36872)
	NewRatioConversion(ImperialPeck, ImperialBushel, 0.25)

	// us
	NewRatioConversion(Quart, Liter, 0.946352946)
	NewRatioConversion(Pint, Liter, 0.473176473)
	NewRatioConversion(Gallon, Liter, 3.785411784)
	NewRatioConversion(FluidOunce, MilliLiter, 29.5735295625)
	NewRatioConversion(Bushel, CubicInch, 2150.42)
	NewRatioConversion(Peck, Bushel, 0.25)
	NewRatioConversion(DryGallon, Bushel, 0.125)
	NewRatioConversion(DryQuart, DryGallon, 0.25)
	NewRatioConversion(DryPint, DryQuart, 0.5)
	NewRatioConversion(CubicInch, MilliLiter, 16.387064)
	NewRatioConversion(CubicFoot, Liter, 28.316846592)
	NewRatioConversion(Cord, CubicFoot, 128.0)
}
