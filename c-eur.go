package currencydb

import "golang.org/x/text/currency"

var EUR = &Currency{
	ISO:            "EUR", // not a real country
	Type:           Main,
	Country:        "EU",
	Countries:      []string{"AD", "AT", "AX", "BE", "BL", "CY", "DE", "EE", "ES", "FI", "FR", "GF", "GP", "GR", "HR", "IE", "IT", "LT", "LU", "LV", "MC", "ME", "MF", "MQ", "MT", "NL", "PM", "PT", "RE", "SI", "SK", "SM", "VA", "XK", "YT"},
	Name:           "Euro",
	Symbol:         "€",
	Decimals:       2,
	SymbolPosition: After,
	unit:           currency.EUR,
	unitSet:        true,
}
