package currencydb

import "golang.org/x/text/currency"

var GBP = &Currency{
	ISO:            "GBP",
	Type:           Main,
	Country:        "GB",
	Name:           "British Pound",
	Symbol:         "£",
	Decimals:       2,
	SymbolPosition: Before,
	unit:           currency.GBP,
	unitSet:        true,
}
