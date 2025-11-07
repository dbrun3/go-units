package units

// Cube creates a cubic volume unit from a base length unit, automatically generating
// aliases with ³ and 3 suffixes.
func Cube(base Unit, o ...UnitOption) Unit {
	name := "cubic " + base.Name
	symbol := base.Symbol + "3"

	// inherit system from base unit by default
	opts := []UnitOption{}
	if base.system != "" {
		opts = append(opts, UnitOptionSystem(base.system))
	}

	// create cubic aliases
	aliases := []string{
		"cu " + base.Symbol,
		base.Symbol + "³",
	}

	// also add cubic versions of base aliases
	for _, alias := range base.aliases {
		aliases = append(aliases, "cu "+alias)
		aliases = append(aliases, alias+"³")
		aliases = append(aliases, alias+"3")
	}

	// aliases and plural
	opts = append(opts, UnitOptionAliases(aliases...), UnitOptionPlural("cubic "+base.plural))

	// append any supplemental options
	opts = append(opts, o...)

	// create the unit with Volume quantity
	opts = append(opts, UnitOptionQuantity("volume"))
	u := NewUnit(name, symbol, opts...)

	return u
}
