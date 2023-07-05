package currencydb

import "golang.org/x/text/currency"

var CAD = &Currency{
	ISO:            "CAD",
	Type:           Main,
	Country:        "CA",
	Name:           "Canadian Dollar",
	Symbol:         "CA$",
	Decimals:       2,
	SymbolPosition: Before,
	unit:           currency.CAD,
	unitSet:        true,
}
