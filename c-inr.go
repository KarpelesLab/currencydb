package currencydb

import "golang.org/x/text/currency"

var INR = &Currency{
	ISO:            "INR",
	Type:           Main,
	Country:        "IN",
	Countries:      []string{"IN"},
	Name:           "Ruppee",
	Symbol:         "INR",
	Decimals:       2,
	SymbolPosition: After,
	unit:           currency.INR,
	unitSet:        true,
}
