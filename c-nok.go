package currencydb

var NOK = &Currency{
	ISO:            "NOK",
	Country:        "NO",
	Countries:      []string{"BV", "NO", "SJ"},
	Name:           "Norwegian krone",
	Symbol:         "Kr",
	Decimals:       2,
	SymbolPosition: After,
}
