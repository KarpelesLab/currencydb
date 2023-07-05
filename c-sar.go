package currencydb

import "golang.org/x/text/currency"

var SAR = &Currency{
	ISO:            "SAR",
	Type:           Main,
	Country:        "SA",
	Countries:      []string{"SA"},
	Name:           "SAUDI RIYAL",
	Symbol:         "SAR",
	Decimals:       2,
	SymbolPosition: After,
	unit:           currency.SAR,
	unitSet:        true,
}
