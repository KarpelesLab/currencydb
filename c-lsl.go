package currencydb

// The Lesotho loti (LSL) is the national currency of the African Kingdom of Lesotho. The LSL is pegged to the South African Rand (ZAR) at one-to-one, and Lesotho is part of the regional Common Monetary Area. As a result, rand are also frequently found in Lesotho and often accepted as legal tender in exchange.

var LSL = &Currency{
	ISO:            "LSL",
	Type:           Main,
	Country:        "LS",
	Countries:      []string{"LS"},
	Name:           "LOTI",
	Symbol:         "LSL",
	Decimals:       2,
	SymbolPosition: After,
	Equivalent:     ZAR,
}
