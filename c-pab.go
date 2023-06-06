package currencydb

// The PAB is pegged at 1:1 to the U.S. dollar

var PAB = &Currency{
	ISO:            "PAB",
	Type:           Main,
	Country:        "PA",
	Countries:      []string{"PA"},
	Name:           "BALBOA",
	Symbol:         "PAB",
	Decimals:       2,
	SymbolPosition: After,
	Equivalent:     USD,
}
