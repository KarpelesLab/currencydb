package currencydb

var GBP = &Currency{
	ISO:            "GBP",
	Type:           Main,
	Country:        "GB",
	Name:           "British Pound",
	Symbol:         "£",
	Decimals:       2,
	SymbolPosition: Before,
}
