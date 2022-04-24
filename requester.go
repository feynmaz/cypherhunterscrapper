package cypherhunterscrapper

import (
	"errors"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
)

// Requester is the interface that requests cypherhunter.com using github.com/anaskhan96/soup
type Requester interface {
	Request() (string, error)
}

// CoinRequester is the struct containing coinUrl to be requested
type CoinRequester struct {
	coinUrl string
}

// NewCoinRequester gives new value of type CoinRequester. Before creating CoinRequester checks if input url is valid cypherhunter.com url.
//
// Returns new CoinRequester with url and nil error if url is valid else empty CoinRequester and error
func NewCoinRequester(coinUrl string) (CoinRequester, error) {
	url := strings.TrimSpace(coinUrl)
	if len(url) != 0 {
		template := `^\s*(https:\/\/www.cypherhunter.com\/en\/p\/).+(\/)?\s*$`
		re, _ := regexp.Compile(template)

		if re.Match([]byte(url)) {
			return CoinRequester{coinUrl: url}, nil
		}
		return CoinRequester{}, errors.New(`input does not match pattern "https://www.cypherhunter.com/en/p/coin_name/"`)
	}
	return CoinRequester{}, errors.New(`input is empty`)

}

// Implements function Request() of Requester interface
func (c CoinRequester) Request() (string, error) {
	return soup.Get(c.coinUrl)
}


