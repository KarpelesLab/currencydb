package currencydb

import "golang.org/x/text/currency"

var PLN = &Currency{
	ISO:            "PLN",
	Type:           Main,
	Country:        "PL",
	Name:           "Polish złoty",
	Symbol:         "zł",
	Decimals:       2,
	SymbolPosition: After,
	unit:           currency.PLN,
	unitSet:        true,
}
