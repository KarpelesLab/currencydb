package currencydb

import "golang.org/x/text/currency"

var ZAR = &Currency{
	ISO:            "ZAR",
	Type:           Main,
	Country:        "ZA",
	Countries:      []string{"ZA"},
	Name:           "Rand",
	Symbol:         "ZAR",
	Decimals:       2,
	SymbolPosition: After,
	unit:           currency.ZAR,
	unitSet:        true,
}
