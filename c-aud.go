package currencydb

import "golang.org/x/text/currency"

var AUD = &Currency{
	ISO:            "AUD",
	Type:           Main,
	Country:        "AU",
	Countries:      []string{"AU", "CC", "CX", "HM", "KI", "NF", "NR", "TV"},
	Name:           "Australian Dollar",
	Symbol:         "AU$",
	Decimals:       2,
	SymbolPosition: Before,
	unit:           currency.AUD,
	unitSet:        true,
}
