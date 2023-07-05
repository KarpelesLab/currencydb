package currencydb

import "golang.org/x/text/currency"

type Type int

const (
	InvalidType Type = iota
	Main
	PreciousMetal
	Retired
	Digital
	Private
	Other
)

func (t Type) String() string {
	switch t {
	case Main:
		return "Main"
	case PreciousMetal:
		return "PreciousMetal"
	case Retired:
		return "Retired"
	case Digital:
		return "Digital"
	case Private:
		return "Private"
	case Other:
		return "Other"
	default:
		return "InvalidType"
	}
}

type Currency struct {
	ISO            string // 3 letters
	Type           Type
	Country        string // 2 letters
	Countries      []string
	Name           string
	Symbol         string
	Decimals       int
	SymbolPosition Position  // Before or After
	Equivalent     *Currency // if pegged to another currency
	unit           currency.Unit
	unitSet        bool
}

func (c *Currency) Unit() currency.Unit {
	if c.unitSet {
		return c.unit
	}
	i, err := currency.ParseISO(c.ISO)
	if err != nil {
		i = currency.XXX
	}
	c.unit = i
	c.unitSet = true
	return i
}
