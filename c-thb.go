package currencydb

import "golang.org/x/text/currency"

var THB = &Currency{
	ISO:            "THB",
	Type:           Main,
	Country:        "TH",
	Name:           "Thai Baht",
	Symbol:         "฿",
	Decimals:       2,
	SymbolPosition: Before,
	unit:           currency.THB,
	unitSet:        true,
}
