package cypherhunterscrapper

import "testing"

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

	if !contains(investors, "Electric Capital") {
		t.Error("expected Electric Capital in investors")
	}

}
