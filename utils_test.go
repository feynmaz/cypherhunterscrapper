package cypherhunterscrapper

import (
	"fmt"
	"log"
	"testing"
)

func TestValidateUrl(t *testing.T) {
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
		url, err := ValidateUrl(tt.in)
		isError := err != nil

		if url != tt.url {
			t.Errorf("got %v, want %v", url, tt.url)
		}

		if isError != tt.isError {
			t.Errorf("got %v, want %v", isError, tt.isError)
		}
	}
}

func ExampleValidateUrl() {
	rawUrl := "  https://www.cypherhunter.com/en/p/solana/   "
	cypherhunerUrl, err := ValidateUrl(rawUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cypherhunerUrl)
	// Output: https://www.cypherhunter.com/en/p/solana/
}
