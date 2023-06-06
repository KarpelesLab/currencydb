package currencydb

// The ngultrum is currently pegged to the Indian rupee at parity.

var BTN = &Currency{
	ISO:            "BTN",
	Type:           Main,
	Country:        "BT",
	Countries:      []string{"BT"},
	Name:           "Ngultrum",
	Symbol:         "Nu.",
	Decimals:       2,
	SymbolPosition: Before,
	Equivalent:     INR,
}
