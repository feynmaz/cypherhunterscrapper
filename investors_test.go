package cypherhunterscrapper

import (
        "fmt"
        "testing"
)

func Test_InListOfTop(t *testing.T) {
        var testCases = []struct {
                in       string
                outValue Investor
                outOk    bool
        }{
                {"random-name", Investor{}, false},
                {"PANTERA Capital", Investor{
                        Name:   "PANTERA Capital",
                        Link:   "https://www.cypherhunter.com/en/p/pantera-capital/",
                        Tier:   "Premier",
                        TierId: 0,
                }, true},
        }

        for _, c := range testCases {
                value, ok := InListOfTop(c.in)

                if c.outOk != ok {
                        t.Errorf("got %v, expected %v", ok, c.outOk)
                }

                if c.outValue != value {
                        t.Errorf("got %v, expected %v", value, c.outValue)
                }
        }
}

func ExampleInListOfTop() {
        name := "Polychain Capital"
        investor, ok := InListOfTop(name)
        if ok {
                fmt.Printf("%#v", investor)
        }
        // Output: cypherhunterscrapper.Investor{Name:"Polychain Capital", Link:"https://www.cypherhunter.com/en/p/polychain-capital/", Tier:"Premier", TierId:0}
}
