package currencydb

import "golang.org/x/text/currency"

var DKK = &Currency{
	ISO:            "DKK",
	Type:           Main,
	Country:        "DK",
	Countries:      []string{"DK", "FO", "GL"},
	Name:           "Danish kroner",
	Symbol:         "Kr",
	Decimals:       2,
	SymbolPosition: After,
	unit:           currency.DKK,
	unitSet:        true,
}
