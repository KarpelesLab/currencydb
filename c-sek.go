package currencydb

import "golang.org/x/text/currency"

var SEK = &Currency{
	ISO:            "SEK",
	Type:           Main,
	Country:        "SE",
	Name:           "Sweden krona",
	Symbol:         "Kr",
	Decimals:       0,
	SymbolPosition: After,
	unit:           currency.SEK,
	unitSet:        true,
}
