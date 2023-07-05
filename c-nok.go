package currencydb

import "golang.org/x/text/currency"

var NOK = &Currency{
	ISO:            "NOK",
	Type:           Main,
	Country:        "NO",
	Countries:      []string{"BV", "NO", "SJ"},
	Name:           "Norwegian krone",
	Symbol:         "Kr",
	Decimals:       2,
	SymbolPosition: After,
	unit:           currency.NOK,
	unitSet:        true,
}
