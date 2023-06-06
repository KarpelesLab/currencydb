package currencydb

var HUF = &Currency{
	ISO:            "HUF",
	Type:           Main,
	Country:        "HU",
	Countries:      []string{"HU"},
	Name:           "FORINT",
	Symbol:         "HUF",
	Decimals:       2,
	SymbolPosition: After,
}
