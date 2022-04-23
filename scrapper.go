package cypherhunterscrapper

import (
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
)

// Investor is how the cypherhunterscrapper package stores an information about the investor
type Investor struct {
	Name   string
	Link   string
	Tier   string
	TierId int
}

// GetInvestorsAll shows list of all investors of the coin described on the page available by coinUrl
// Returns array of string with investor names and error that can be nil
func GetInvestorsAll(coinUrl string) ([]string, error) {

	resp, err := soup.Get(coinUrl)
	if err != nil {
		return nil, err
	}

	if strings.Contains(resp, "Oops! That page doesn't exist") {
		return nil, fmt.Errorf("that page doesn't exist")

	}

	doc := soup.HTMLParse(resp)

	coinFullName := doc.Find("h1").Text()

	var divInvestors soup.Root
	headers := doc.FindAll("h2")
	for _, h := range headers {
		t := h.Text()
		if strings.Contains(t, "Investors") {
			divInvestors = h.FindNextSibling()
			break
		}
	}

	if divInvestors == (soup.Root{}) {
		return nil, fmt.Errorf("no information about investors for %v", coinFullName)

	} else {
		aa := divInvestors.FindAll("a")
		var investorsAll = make([]string, len(aa))

		for i, a := range aa {
			investorName := a.Attrs()["title"]
			investorsAll[i] = investorName

		}
		return investorsAll, nil
	}
}
