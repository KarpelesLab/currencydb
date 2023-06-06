[![GoDoc](https://godoc.org/github.com/KarpelesLab/currencydb?status.svg)](https://godoc.org/github.com/KarpelesLab/currencydb)

# Currency database

A simple list of all known currencies.

Currencies can be accessed in multiple ways:

* Through the currency name if hardcoded, for example `currencydb.EUR`
* Via the `currencydb.All` map (lookup by string), for example `cur := currencydb.All[isoCode]`
* Via the `currencydb.Country` map (lookup by country code), for example `cur := currencydb.Country["US"]` would return `currencydb.USD`
