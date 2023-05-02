package currencydb

var EUR = &Currency{
	ISO:            "EUR",
	Country:        "EU", // not a real country
	Name:           "Euro",
	Symbol:         "€",
	Decimals:       2,
	SymbolPosition: After,
}
