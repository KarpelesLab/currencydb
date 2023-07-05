package currencydb

import "golang.org/x/text/currency"

var RUB = &Currency{
	ISO:            "RUB",
	Type:           Main,
	Country:        "RU",
	Name:           "Russian Ruble",
	Symbol:         "RUB",
	Decimals:       2,
	SymbolPosition: After,
	unit:           currency.RUB,
	unitSet:        true,
}
