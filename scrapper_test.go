package cypherhunterscrapper

import (
	"fmt"
	"log"
	"testing"
)

func TestGetInvestorsAllBadUrl(t *testing.T) {
	coinUrl := "http://bad_url"
	_, err := GetInvestorsAll(coinUrl)

	if err == nil {
		t.Errorf("expected error on %v", coinUrl)
	}
}

func TestGetInvestorsAllNoPage(t *testing.T) {
	coinUrl := "https://www.cypherhunter.com/en/p/non-existent-endpoint/"
	_, err := GetInvestorsAll(coinUrl)

	if err == nil {
		t.Errorf("expected error on %v", coinUrl)
	}
}

func TestGetInvestorsAllNoInvestors(t *testing.T) {
	coinUrl := "https://www.cypherhunter.com/en/p/bitcoin/"
	_, err := GetInvestorsAll(coinUrl)

	if err == nil {
		t.Errorf("expected error on %v", coinUrl)
	}
}

func TestGetInvestorsAllOk(t *testing.T) {
	coinUrl := "https://www.cypherhunter.com/en/p/ethereum/"
	investors, _ := GetInvestorsAll(coinUrl)

	l := len(investors)
	maxL := 10
	minL := 5

	if l > maxL {
		t.Errorf("got len %v expected less than %v", l, maxL)
	}
	if l < minL {
		t.Errorf("got len %v expected more than %v", l, minL)
	}

	expectedIn := "Electric Capital"
	contains := func(s []string, e string) bool {
		for _, a := range s {
			if a == e {
				return true
			}
		}
		return false
	}
	if !contains(investors, expectedIn) {
		t.Errorf("expected %v in investors", expectedIn)
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
		{[]string{}, []Investor{}},
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
		}, {
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
