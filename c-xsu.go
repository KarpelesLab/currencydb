package currencydb

// https://en.wikipedia.org/wiki/SUCRE

var XSU = &Currency{
	ISO:            "XSU",
	Type:           Other,
	Countries:      []string{"BO", "CU", "EC", "NI", "VE"},
	Name:           "SUCRE (unit of account)",
	Symbol:         "XSU",
	Decimals:       0,
	SymbolPosition: After,
}
