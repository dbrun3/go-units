package units

// Square creates a squared area unit from a base length unit, automatically generating
// aliases with ² and 2 suffixes, and registering conversions.
func Square(base Unit, o ...UnitOption) Unit {
	name := "square " + base.Name
	symbol := "sq " + base.Symbol

	// inherit system from base unit by default
	opts := []UnitOption{}
	if base.system != "" {
		opts = append(opts, UnitOptionSystem(base.system))
	}

	// create squared aliases
	aliases := []string{
		base.Symbol + "²",
		base.Symbol + "2",
	}

	// also add squared versions of base aliases
	for _, alias := range base.aliases {
		aliases = append(aliases, alias+"²")
		aliases = append(aliases, alias+"2")
	}

	// aliases and plural
	opts = append(opts, UnitOptionAliases(aliases...), UnitOptionPlural("square "+base.plural))

	// append any supplemental options
	opts = append(opts, o...)

	// create the unit with Area quantity
	opts = append(opts, UnitOptionQuantity("area"))
	u := NewUnit(name, symbol, opts...)

	return u
}
