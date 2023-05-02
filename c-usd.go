package currencydb

var USD = &Currency{
	ISO:            "USD",
	Country:        "US",
	Name:           "US Dollar",
	Symbol:         "$",
	Decimals:       2,
	SymbolPosition: Before,
}
