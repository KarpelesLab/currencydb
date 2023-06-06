package currencydb

var NZD = &Currency{
	ISO:            "NZD",
	Country:        "NZ",
	Countries:      []string{"CK", "NU", "NZ", "PN", "TK"},
	Name:           "New Zealand Dollar",
	Symbol:         "NZ$",
	Decimals:       2,
	SymbolPosition: Before,
}
