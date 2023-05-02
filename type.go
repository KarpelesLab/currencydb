package currencydb

type Currency struct {
	ISO            string // 3 letters
	Country        string // 2 letters
	Name           string
	Symbol         string
	Decimals       int
	SymbolPosition Position // Before or After
}
