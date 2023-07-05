package currencydb

import "golang.org/x/text/currency"

var TWD = &Currency{
	ISO:            "TWD",
	Type:           Main,
	Country:        "TW",
	Countries:      []string{"TW"},
	Name:           "New Taiwan Dollar",
	Symbol:         "NT$",
	Decimals:       2,
	SymbolPosition: Before,
	unit:           currency.TWD,
	unitSet:        true,
}
