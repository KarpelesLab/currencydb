package currencydb

type Type int

const (
	InvalidType Type = iota
	Main
	PreciousMetal
	Retired
	Digital
	Private
	Other
)

func (t Type) String() string {
	switch t {
	case Main:
		return "Main"
	case PreciousMetal:
		return "PreciousMetal"
	case Other:
		return "Other"
	default:
		return "InvalidType"
	}
}

type Currency struct {
	ISO            string // 3 letters
	Type           Type
	Country        string // 2 letters
	Countries      []string
	Name           string
	Symbol         string
	Decimals       int
	SymbolPosition Position  // Before or After
	Equivalent     *Currency // if pegged to another currency
}
