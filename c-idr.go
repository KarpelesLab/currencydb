package currencydb

import "golang.org/x/text/currency"

var IDR = &Currency{
	ISO:            "IDR",
	Type:           Main,
	Country:        "ID",
	Countries:      []string{"ID"},
	Name:           "Rupiah",
	Symbol:         "Rp",
	Decimals:       2,
	SymbolPosition: Before,
	unit:           currency.IDR,
	unitSet:        true,
}
