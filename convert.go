//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/KarpelesLab/currencydb"
)

func main() {
	err := doit()
	if err != nil {
		log.Printf("Failed to convert: %s", err)
		os.Exit(1)
	}
}

func doit() error {
	in, err := os.Open("...")
	if err != nil {
		return err
	}
	defer in.Close()
	inb := bufio.NewReader(in)

	flds, err := inb.ReadString('\n')
	if err != nil {
		return err
	}
	_ = flds

	currencies := make(map[string]*currencydb.Currency)

	for {
		t, err := inb.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		nfo := strings.Split(strings.TrimSpace(t), "\t")
		//log.Printf("nfo = %+v", nfo)
		code := nfo[2]

		cur, ok := currencies[code]
		if ok {
			// add country only
			cur.Countries = append(cur.Countries, nfo[5])
			continue
		}

		dec, err := strconv.Atoi(nfo[4])
		if err != nil {
			return err
		}

		if len(nfo) < 6 {
			// no country
			cur = &currencydb.Currency{
				ISO:            code,
				Type:           currencydb.Other,
				Name:           nfo[3],
				Symbol:         code,
				Decimals:       dec,
				SymbolPosition: currencydb.After,
			}
			currencies[code] = cur
			continue
		}

		cur = &currencydb.Currency{
			ISO:            code,
			Type:           currencydb.Main,
			Country:        nfo[5],
			Countries:      []string{nfo[5]},
			Name:           nfo[3],
			Symbol:         code,
			Decimals:       dec,
			SymbolPosition: currencydb.After,
		}
		currencies[code] = cur
	}

	// for each currency
	allCurs := make([]string, 0, len(currencies))
	for _, cur := range currencies {
		allCurs = append(allCurs, cur.ISO)

		log.Printf("checking currency %s", cur.ISO)
		ldata, ok := currencydb.All[cur.ISO]
		if !ok {
			log.Printf("creating missing currency %s", cur.ISO)
			if err := writeCurrency(cur); err != nil {
				return err
			}
			continue
		}
		if ldata.Type == currencydb.InvalidType {
			ldata.Type = cur.Type
			if err := writeCurrency(ldata); err != nil {
				return err
			}
		}
	}

	sort.Strings(allCurs)

	// generate list.go
	return writeListGo(allCurs)
}

func writeCurrency(cur *currencydb.Currency) error {
	fn := "c-" + strings.ToLower(cur.ISO) + ".go"
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "package currencydb\n\n")
	fmt.Fprintf(f, "var %s = &Currency{\n", cur.ISO)
	fmt.Fprintf(f, "\tISO: %q,\n", cur.ISO)
	fmt.Fprintf(f, "\tType: %s,\n", cur.Type)
	if cur.Country != "" {
		fmt.Fprintf(f, "\tCountry: %q,\n", cur.Country)
	}
	if cur.Countries != nil {
		fmt.Fprintf(f, "\tCountries: %#v,\n", cur.Countries)
	}
	fmt.Fprintf(f, "\tName: %q,\n", cur.Name)
	fmt.Fprintf(f, "\tSymbol: %q,\n", cur.Symbol)
	fmt.Fprintf(f, "\tDecimals: %d,\n", cur.Decimals)
	fmt.Fprintf(f, "\tSymbolPosition: %s,\n", cur.SymbolPosition)
	fmt.Fprintf(f, "}\n")
	return nil
}

func writeListGo(l []string) error {
	f, err := os.Create("list.go")
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "package currencydb\n\n")
	fmt.Fprintf(f, "var All = map[string]*Currency{\n")
	for _, c := range l {
		fmt.Fprintf(f, "\t%q: %s,\n", c, c)
	}
	fmt.Fprintf(f, "}\n")
	return nil
}
