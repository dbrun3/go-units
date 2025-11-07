package units

var (
	Area = UnitOptionQuantity("area")

	// metric
	SquareMeter      = Square(Meter)
	SquareKilometer  = Square(KiloMeter)
	SquareCentimeter = Square(CentiMeter)
	SquareMillimeter = Square(MilliMeter)
	Hectare          = NewUnit("hectare", "ha", Area, SI)

	// imperial/us - using imperial length units from length_units.go
	SquareInch = Square(Inch)
	SquareFoot = Square(Foot)
	SquareYard = Square(Yard)
	SquareMile = Square(Mile)
	Acre       = NewUnit("acre", "ac", UnitOptionAliases("acres", "a"), Area, US)
)

func init() {
	// metric
	NewRatioConversion(SquareCentimeter, SquareMeter, 0.0001)
	NewRatioConversion(SquareMillimeter, SquareMeter, 0.000001)
	NewRatioConversion(SquareKilometer, SquareMeter, 1000000.0)
	NewRatioConversion(Hectare, SquareMeter, 10000.0)

	// us
	NewRatioConversion(SquareInch, SquareMeter, 0.00064516)
	NewRatioConversion(SquareFoot, SquareInch, 144.0)
	NewRatioConversion(SquareYard, SquareFoot, 9.0)
	NewRatioConversion(SquareMile, SquareMeter, 2589988.110336)
	NewRatioConversion(Acre, SquareFoot, 43560.0)
}
