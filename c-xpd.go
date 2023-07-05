package currencydb

import "golang.org/x/text/currency"

var XPD = &Currency{
	ISO:            "XPD",
	Type:           PreciousMetal,
	Name:           "Palladium",
	Symbol:         "XPD",
	Decimals:       5,
	SymbolPosition: After,
	unit:           currency.XPD,
	unitSet:        true,
}
