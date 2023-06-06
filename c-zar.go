package currencydb

var ZAR = &Currency{
	ISO:            "ZAR",
	Type:           Main,
	Country:        "LS",
	Countries:      []string{"LS", "NA", "ZA"},
	Name:           "RAND",
	Symbol:         "ZAR",
	Decimals:       2,
	SymbolPosition: After,
}
