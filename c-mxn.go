package currencydb

var MXN = &Currency{
	ISO:            "MXN",
	Type:           Main,
	Country:        "MX",
	Countries:      []string{"MX"},
	Name:           "MEXICAN PESO",
	Symbol:         "MXN",
	Decimals:       2,
	SymbolPosition: After,
}
