package cypherhunterscrapper

import (
	"fmt"
	"log"
	"testing"
)

func TestErrorsGetInvestorsAll(t *testing.T) {
	// TableDriver https://github.com/golang/go/wiki/TableDrivenTests
	var testCases = []struct {
		in        string
		errorText string
	}{
		{"http://bad_url", "couldn't perform GET request to http://bad_url"},
		{"https://www.cypherhunter.com/en/p/bitcoins/", "that page doesn't exist"},
		{"https://www.cypherhunter.com/en/p/bitcoin/", "no information about investors for Bitcoin"},
	}

	for _, c := range testCases {
		_, err := GetInvestorsAll(c.in)

		if err == nil {
			t.Errorf("expected error on input %v", c.in)
		}

		if err.Error() != c.errorText {
			t.Errorf("got %v, expected %v", err.Error(), c.errorText)
		}
	}

}

func TestGetInvestorsAll(t *testing.T) {

	contains := func(s []string, e string) bool {
		for _, a := range s {
			if a == e {
				return true
			}
		}
		return false
	}

	in := "https://www.cypherhunter.com/en/p/ethereum/"
	investors, _ := GetInvestorsAll(in)

	if len(investors) < 5 {
		t.Errorf("got %v, expected >= 5", len(investors))
	}

	if len(investors) > 10 {
		t.Errorf("got %v, expected <= 10", len(investors))
	}

	if !contains(investors, "Electric Capital") {
		t.Error("expected Electric Capital in investors")
	}
}

func ExampleGetInvestorsAll() {
	coinUrl := "https://www.cypherhunter.com/en/p/ethereum/"
	investors, err := GetInvestorsAll(coinUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", investors)
	// Output: []string{"Wavemaker Genesis", "KR1", "Electric Capital", "Breyer Capital", "8 Decimal Capital", "PANTERA Capital"}
}

func TestGetInvestorsExceptional(t *testing.T) {
	var testCases = []struct {
		in  []string
		out []Investor
	}{
		{[]string{""}, []Investor{}},
		{[]string{"random one", "random two"}, []Investor{}},
		{[]string{"ANDREESSEN HOROWITZ", "random"}, []Investor{{
				Name:   "ANDREESSEN HOROWITZ",
				Link:   "https://www.cypherhunter.com/en/p/andreessen-horowitz/",
				Tier:   "Premier",
				TierId: 0,
			},
		}},
		{[]string{"ANDREESSEN HOROWITZ", "SOLANA"}, []Investor{{
				Name:   "ANDREESSEN HOROWITZ",
				Link:   "https://www.cypherhunter.com/en/p/andreessen-horowitz/",
				Tier:   "Premier",
				TierId: 0,
			},{
				Name:   "SOLANA",
				Link:   "https://www.cypherhunter.com/en/p/solana/",
				Tier:   "First",
				TierId: 1,
			},
		}},
		
	}

	areEqual := func(a, b []Investor) bool {
		if len(a) == 0 && len(b) == 0 {
			return true
		}

		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i].Name != b[i].Name {
				return false
			}
		}
		return true
	}

	for _, c := range testCases {
		list := GetInvestorsExceptional(c.in)

		if !areEqual(c.out, list) {
			t.Errorf("got %v expected %v", list, c.out)
		}

	}
}

func ExampleGetInvestorsExceptional() {
	investorsAll := []string{"PANTERA Capital", "Random Investor", "GBV Capital"}
	investorsExceprional := GetInvestorsExceptional(investorsAll)

	fmt.Printf("%#v", investorsExceprional)
	// Output: []cypherhunterscrapper.Investor{cypherhunterscrapper.Investor{Name:"PANTERA Capital", Link:"https://www.cypherhunter.com/en/p/pantera-capital/", Tier:"Premier", TierId:0}, cypherhunterscrapper.Investor{Name:"GBV Capital", Link:"https://www.cypherhunter.com/en/p/gbv-capital/", Tier:"First", TierId:1}}
}
