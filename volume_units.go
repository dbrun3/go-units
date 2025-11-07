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
	ImperialQuart      = NewUnit("imperial quart", "uk qt", Volume, BI)
	ImperialPint       = NewUnit("imperial pint", "uk pt", Volume, BI)
	ImperialGallon     = NewUnit("imperial gallon", "uk gal", Volume, BI)
	ImperialFluidOunce = NewUnit("imperial fluid ounce", "uk fl oz", Volume, BI)
	ImperialBushel     = NewUnit("imperial bushel", "uk bu", Volume, BI)
	ImperialPeck       = NewUnit("imperial peck", "uk pk", Volume, BI)

	// us
	Quart      = NewUnit("quart", "qt", Volume, US)
	Pint       = NewUnit("pint", "pt", Volume, US)
	Gallon     = NewUnit("gallon", "gal", Volume, US)
	FluidOunce = NewUnit("fluid ounce", "fl oz", Volume, US, UnitOptionAliases("floz"))
	Bushel     = NewUnit("bushel", "bu", Volume, US)
	Peck       = NewUnit("peck", "pk", Volume, US)
	CubicInch  = Cube(Inch)
	CubicFoot  = Cube(Foot)
	Cord       = NewUnit("cord", "cord", UnitOptionAliases("cd"), Volume, US)
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
	NewRatioConversion(CubicInch, MilliLiter, 16.387064)
	NewRatioConversion(CubicFoot, Liter, 28.316846592)
	NewRatioConversion(Cord, CubicFoot, 128.0)
}
