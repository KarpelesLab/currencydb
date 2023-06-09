package currencydb

import "golang.org/x/text/currency"

var USD = &Currency{
	ISO:            "USD",
	Type:           Main,
	Country:        "US",
	Countries:      []string{"AS", "BQ", "EC", "FM", "GU", "IO", "MH", "MP", "PR", "PW", "SV", "TC", "TL", "UM", "US", "VG", "VI"},
	Name:           "US Dollar",
	Symbol:         "$",
	Decimals:       2,
	SymbolPosition: Before,
	unit:           currency.USD,
	unitSet:        true,
}
