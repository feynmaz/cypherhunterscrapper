package cypherhunterscrapper

import (
	"fmt"
	"sort"
	"strings"

	"github.com/anaskhan96/soup"
)

// GetInvestorsAll gives list of all investors of the coin described on the page available by coinUrl.
//
// Returns array of string with investor names and possibly nil error
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

// GetInvestorsExceptional finds top investors in list of investors names.
// 
// Returns possibly empty list of top investors of type Investor
func GetInvestorsExceptional(investors []string) []Investor {
	if len(investors) == 0 {
		return []Investor{}
	}
	var topInvestors []Investor

	for _, name := range investors {
		i, ok := InListOfTop(name)
		if ok {
			topInvestors = append(topInvestors, i)
		}
	}

	if len(topInvestors) > 0 {
		sort.Slice(topInvestors, func(i, j int) bool { return topInvestors[i].TierId < topInvestors[j].TierId })
	}
	return topInvestors
}
