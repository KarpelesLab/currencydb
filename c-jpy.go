package currencydb

import "golang.org/x/text/currency"

var JPY = &Currency{
	ISO:            "JPY",
	Type:           Main,
	Country:        "JP",
	Name:           "Japanese Yen",
	Symbol:         "¥",
	Decimals:       0,
	SymbolPosition: Before,
	unit:           currency.JPY,
	unitSet:        true,
}
