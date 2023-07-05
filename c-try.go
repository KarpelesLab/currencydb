package currencydb

import "golang.org/x/text/currency"

var TRY = &Currency{
	ISO:            "TRY",
	Type:           Main,
	Country:        "TR",
	Countries:      []string{"TR"},
	Name:           "Turkish Lira",
	Symbol:         "TRY",
	Decimals:       2,
	SymbolPosition: After,
	unit:           currency.TRY,
	unitSet:        true,
}
