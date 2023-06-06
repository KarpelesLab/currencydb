package currencydb

var ILS = &Currency{
	ISO:            "ILS",
	Type:           Main,
	Country:        "IL",
	Countries:      []string{"IL", "PS"},
	Name:           "New Israeli Shekel",
	Symbol:         "₪",
	Decimals:       2,
	SymbolPosition: Before,
}
