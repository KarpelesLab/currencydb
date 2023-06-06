package currencydb

var INR = &Currency{
	ISO:            "INR",
	Type:           Main,
	Country:        "IN",
	Countries:      []string{"BT", "IN"},
	Name:           "Ruppee",
	Symbol:         "INR",
	Decimals:       2,
	SymbolPosition: After,
}
