//go:build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/KarpelesLab/currencydb"
)

func main() {
	// gather all currencies & list countries
	countryList := make(map[string]*currencydb.Currency)

	for _, cur := range currencydb.All {
		if cur.Type != currencydb.Main {
			continue
		}
		if cur.Countries == nil && cur.Country != "" {
			if fnd, ok := countryList[cur.Country]; ok {
				log.Printf("CONFLICT: country %s got currencies %s and %s", cur.Country, fnd.ISO, cur.ISO)
			}
			countryList[cur.Country] = cur
		} else {
			for _, cc := range cur.Countries {
				if fnd, ok := countryList[cc]; ok {
					log.Printf("CONFLICT: country %s got currencies %s and %s", cc, fnd.ISO, cur.ISO)
				}
				countryList[cc] = cur
			}
		}
	}

	sortedCountries := make([]string, 0, len(countryList))
	for k := range countryList {
		sortedCountries = append(sortedCountries, k)
	}
	sort.Strings(sortedCountries)

	// output file
	f, err := os.Create("countries.go")
	if err != nil {
		log.Printf("failed to gen countries.go: %s", err)
		os.Exit(1)
	}

	fmt.Fprintf(f, "package currencydb\n\n")
	fmt.Fprintf(f, "//go:generate go run gen.go\n\n")
	fmt.Fprintf(f, "var Country = map[string]*Currency{\n")
	for _, id := range sortedCountries {
		fmt.Fprintf(f, "\t%q: %s,\n", id, countryList[id].ISO)
	}
	fmt.Fprintf(f, "}\n")
}
