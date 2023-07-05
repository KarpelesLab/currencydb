package currencydb

import "golang.org/x/text/currency"

var CHF = &Currency{
	ISO:            "CHF",
	Type:           Main,
	Country:        "CH",
	Countries:      []string{"CH", "LI"},
	Name:           "Swiss Franc",
	Symbol:         "CHF",
	Decimals:       2,
	SymbolPosition: After,
	unit:           currency.CHF,
	unitSet:        true,
}
