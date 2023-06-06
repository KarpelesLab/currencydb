package currencydb

var EGP = &Currency{
	ISO:            "EGP",
	Type:           Main,
	Country:        "EG",
	Countries:      []string{"EG"},
	Name:           "EGYPTIAN POUND",
	Symbol:         "EGP",
	Decimals:       2,
	SymbolPosition: After,
}
