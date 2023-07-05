package currencydb

import "golang.org/x/text/currency"

var NZD = &Currency{
	ISO:            "NZD",
	Type:           Main,
	Country:        "NZ",
	Countries:      []string{"CK", "NU", "NZ", "PN", "TK"},
	Name:           "New Zealand Dollar",
	Symbol:         "NZ$",
	Decimals:       2,
	SymbolPosition: Before,
	unit:           currency.NZD,
	unitSet:        true,
}
