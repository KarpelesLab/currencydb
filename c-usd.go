package currencydb

var USD = &Currency{
	ISO:            "USD",
	Country:        "US",
	Countries:      []string{"AS", "BQ", "EC", "FM", "GU", "HT", "IO", "MH", "MP", "PA", "PR", "PW", "SV", "TC", "TL", "UM", "US", "VG", "VI"},
	Name:           "US Dollar",
	Symbol:         "$",
	Decimals:       2,
	SymbolPosition: Before,
}
