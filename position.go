package currencydb

type Position int

const (
	PositionInvalid Position = iota
	Before
	After
)

func (p Position) String() string {
	switch p {
	case Before:
		return "Before"
	case After:
		return "After"
	default:
		return "PositionInvalid"
	}
}
