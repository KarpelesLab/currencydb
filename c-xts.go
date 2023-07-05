package currencydb

import "golang.org/x/text/currency"

var XTS = &Currency{
	ISO:            "XTS",
	Type:           Other,
	Name:           "Test",
	Symbol:         "XTS",
	Decimals:       2,
	SymbolPosition: After,
	unit:           currency.XTS,
	unitSet:        true,
}
