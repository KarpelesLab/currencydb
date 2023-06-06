package currencydb

//go:generate go run gen.go

var Country = map[string]*Currency{
	"AD": EUR,
	"AE": AED,
	"AF": AFN,
	"AL": ALL,
	"AS": USD,
	"AT": EUR,
	"AU": AUD,
	"AX": EUR,
	"BE": EUR,
	"BG": BGN,
	"BL": EUR,
	"BQ": USD,
	"BR": BRL,
	"BT": INR,
	"CA": CAD,
	"CC": AUD,
	"CH": CHF,
	"CK": NZD,
	"CL": CLP,
	"CN": CNY,
	"CX": AUD,
	"CY": EUR,
	"DE": EUR,
	"DK": DKK,
	"EC": USD,
	"EE": EUR,
	"ES": EUR,
	"FI": EUR,
	"FM": USD,
	"FO": DKK,
	"FR": EUR,
	"GB": GBP,
	"GF": EUR,
	"GL": DKK,
	"GP": EUR,
	"GR": EUR,
	"GU": USD,
	"HK": HKD,
	"HM": AUD,
	"HR": EUR,
	"HT": USD,
	"IE": EUR,
	"IL": ILS,
	"IN": INR,
	"IO": USD,
	"IT": EUR,
	"JP": JPY,
	"KI": AUD,
	"KP": KPW,
	"KR": KRW,
	"LI": CHF,
	"LT": EUR,
	"LU": EUR,
	"LV": EUR,
	"MC": EUR,
	"ME": EUR,
	"MF": EUR,
	"MH": USD,
	"MP": USD,
	"MQ": EUR,
	"MT": EUR,
	"NF": AUD,
	"NL": EUR,
	"NO": NOK,
	"NR": AUD,
	"NU": NZD,
	"NZ": NZD,
	"PA": USD,
	"PL": PLN,
	"PM": EUR,
	"PN": NZD,
	"PR": USD,
	"PS": ILS,
	"PT": EUR,
	"PW": USD,
	"RE": EUR,
	"RU": RUB,
	"SE": SEK,
	"SG": SGD,
	"SI": EUR,
	"SK": EUR,
	"SM": EUR,
	"SV": USD,
	"TC": USD,
	"TH": THB,
	"TK": NZD,
	"TL": USD,
	"TV": AUD,
	"UM": USD,
	"US": USD,
	"VA": EUR,
	"VG": USD,
	"VI": USD,
	"XK": EUR,
	"XZ": CZK,
	"YT": EUR,
}
