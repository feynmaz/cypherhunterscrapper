package cypherhunterscrapper

import (
	"fmt"
	"log"
	"testing"
)

func TestNewCoinRequester(t *testing.T) {
	// TableDriver https://github.com/golang/go/wiki/TableDrivenTests
	var tests = []struct {
		in      string
		url     string
		isError bool
	}{
		{" ", "", true},
		{"https://stackoverflow.com/", "", true},
		{"  https://www.cypherhunter.com/en/p/solana/   ", "https://www.cypherhunter.com/en/p/solana/", false},
		{"https://www.cypherhunter.com/en/p/solana/", "https://www.cypherhunter.com/en/p/solana/", false},
	}

	for _, tt := range tests {
		cr, err := NewCoinRequester(tt.in)
		isError := err != nil

		if cr.coinUrl != tt.url {
			t.Errorf("got %v, want %v", cr.coinUrl, tt.url)
		}

		if isError != tt.isError {
			t.Errorf("got %v, want %v", isError, tt.isError)
		}
	}
}

func ExampleNewCoinRequester() {
	rawUrl := "  https://www.cypherhunter.com/en/p/solana/   "
	cr, err := NewCoinRequester(rawUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", cr)
	// Output: cypherhunterscrapper.CoinRequester{coinUrl:"https://www.cypherhunter.com/en/p/solana/"}
}
