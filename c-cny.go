package currencydb

import "golang.org/x/text/currency"

var CNY = &Currency{
	ISO:            "CNY",
	Type:           Main,
	Country:        "CN",
	Name:           "Yuan Renminbi",
	Symbol:         "¥",
	Decimals:       2,
	SymbolPosition: After,
	unit:           currency.CNY,
	unitSet:        true,
}
