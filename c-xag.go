package currencydb

import "golang.org/x/text/currency"

var XAG = &Currency{
	ISO:            "XAG",
	Type:           PreciousMetal,
	Name:           "Silver",
	Symbol:         "XAG",
	Decimals:       5,
	SymbolPosition: After,
	unit:           currency.XAG,
	unitSet:        true,
}
