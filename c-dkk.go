package currencydb

var DKK = &Currency{
	ISO:            "DKK",
	Country:        "DK",
	Countries:      []string{"DK", "FO", "GL"},
	Name:           "Danish kroner",
	Symbol:         "Kr",
	Decimals:       2,
	SymbolPosition: After,
}
