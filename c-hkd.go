package currencydb

import "golang.org/x/text/currency"

var HKD = &Currency{
	ISO:            "HKD",
	Type:           Main,
	Country:        "HK",
	Name:           "Kong Kong Dollar",
	Symbol:         "HK$",
	Decimals:       2,
	SymbolPosition: Before,
	unit:           currency.HKD,
	unitSet:        true,
}
