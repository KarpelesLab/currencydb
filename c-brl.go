package currencydb

import "golang.org/x/text/currency"

var BRL = &Currency{
	ISO:            "BRL",
	Type:           Main,
	Country:        "BR",
	Name:           "Real",
	Symbol:         "R$",
	Decimals:       2,
	SymbolPosition: Before,
	unit:           currency.BRL,
	unitSet:        true,
}
