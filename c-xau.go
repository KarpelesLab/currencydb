package currencydb

import "golang.org/x/text/currency"

var XAU = &Currency{
	ISO:            "XAU",
	Type:           PreciousMetal,
	Name:           "Gold",
	Symbol:         "XAU",
	Decimals:       5,
	SymbolPosition: After,
	unit:           currency.XAU,
	unitSet:        true,
}
