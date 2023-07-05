package currencydb

import "golang.org/x/text/currency"

var KRW = &Currency{
	ISO:            "KRW",
	Type:           Main,
	Country:        "KR",
	Name:           "South Korean won",
	Symbol:         "₩",
	Decimals:       0,
	SymbolPosition: Before,
	unit:           currency.KRW,
	unitSet:        true,
}
