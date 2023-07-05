package currencydb

import "golang.org/x/text/currency"

var MXN = &Currency{
	ISO:            "MXN",
	Type:           Main,
	Country:        "MX",
	Countries:      []string{"MX"},
	Name:           "MEXICAN PESO",
	Symbol:         "MXN",
	Decimals:       2,
	SymbolPosition: After,
	unit:           currency.MXN,
	unitSet:        true,
}
