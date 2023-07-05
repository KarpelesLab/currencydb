package currencydb

import "golang.org/x/text/currency"

var XPT = &Currency{
	ISO:            "XPT",
	Type:           PreciousMetal,
	Name:           "Platinum",
	Symbol:         "XPT",
	Decimals:       5,
	SymbolPosition: After,
	unit:           currency.XPT,
	unitSet:        true,
}
