package cypherhunterscrapper

// Investor is how the cypherhunterscrapper package stores an information about the investor
type Investor struct {
	Name   string
	Link   string
	Tier   string
	TierId int
}

// InListOfTop gives Investor if passed name is in list of top investors. Is prefered to be used in Comma Ok Idiom.
// 
// Returns Investor and bool meaning that the name is on the list of top investors
func InListOfTop(name string) (Investor, bool) {
	for _, i := range topInvestors {
		if i.Name == name {
			return i, true
		}
	}
	return Investor{}, false
}


// topInvestors is the list of top invstors
var topInvestors = []Investor{
	{
		Name:   "PANTERA Capital",
		Link:   "https://www.cypherhunter.com/en/p/pantera-capital/",
		Tier:   "Premier",
		TierId: 0,
	}, {
		Name:   "Fenbushi Capital",
		Link:   "https://www.cypherhunter.com/en/p/fenbushi-capital/",
		Tier:   "Premier",
		TierId: 0,
	}, {
		Name:   "Fenbushi Digital",
		Link:   "https://www.cypherhunter.com/en/p/fenbushi-digital/",
		Tier:   "Premier",
		TierId: 0,
	}, {
		Name:   "FBG Capital",
		Link:   "https://www.cypherhunter.com/en/p/fbg-capital/",
		Tier:   "Premier",
		TierId: 0,
	}, {
		Name:   "Arrington XRP Capital",
		Link:   "https://www.cypherhunter.com/en/p/arrington-xrp-capital/",
		Tier:   "Premier",
		TierId: 0,
	}, {
		Name:   "Polychain Capital",
		Link:   "https://www.cypherhunter.com/en/p/polychain-capital/",
		Tier:   "Premier",
		TierId: 0,
	}, {
		Name:   "FTX",
		Link:   "https://www.cypherhunter.com/en/p/ftx/",
		Tier:   "Premier",
		TierId: 0,
	}, {
		Name:   "GBIC",
		Link:   "https://www.cypherhunter.com/en/p/gbic/",
		Tier:   "Premier",
		TierId: 0,
	}, {
		Name:   "ANDREESSEN HOROWITZ",
		Link:   "https://www.cypherhunter.com/en/p/andreessen-horowitz/",
		Tier:   "Premier",
		TierId: 0,
	}, {
		Name:   "Binance Labs",
		Link:   "https://www.cypherhunter.com/en/p/binance-labs/",
		Tier:   "Premier",
		TierId: 0,
	}, {
		Name:   "Coinbase Ventures",
		Link:   "https://www.cypherhunter.com/en/p/coinbase-ventures/",
		Tier:   "Premier",
		TierId: 0,
	},

	{
		Name:   "KuCoin",
		Link:   "https://www.cypherhunter.com/en/p/kucoin/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "NGC Ventures",
		Link:   "https://www.cypherhunter.com/en/p/ngc-ventures/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "OKX",
		Link:   "https://www.cypherhunter.com/en/p/okx/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "Huobi Ventures",
		Link:   "https://www.cypherhunter.com/en/p/huobi-ventures/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "Huobi Group",
		Link:   "https://www.cypherhunter.com/en/p/huobi-group/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "Huobi Labs",
		Link:   "https://www.cypherhunter.com/en/p/huobi-labs/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "SOLANA",
		Link:   "https://www.cypherhunter.com/en/p/solana/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "Animoca Brands",
		Link:   "https://www.cypherhunter.com/en/p/animoca-brands/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "Kenetic Capital",
		Link:   "https://www.cypherhunter.com/en/p/kenetic-capital/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "Alameda Research",
		Link:   "https://www.cypherhunter.com/en/p/alameda-research/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "GBV Capital",
		Link:   "https://www.cypherhunter.com/en/p/gbv-capital/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "Spark Digital Capital",
		Link:   "https://www.cypherhunter.com/en/p/spark-digital-capital/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "Three Arrows Capital",
		Link:   "https://www.cypherhunter.com/en/p/three-arrows-capital/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "3Commas Capital",
		Link:   "https://www.cypherhunter.com/en/p/3commas-capital/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "LD Capital",
		Link:   "https://www.cypherhunter.com/en/p/ld-capital/",
		Tier:   "First",
		TierId: 1,
	}, {
		Name:   "SkyVision Capital",
		Link:   "https://www.cypherhunter.com/en/p/skyvision-capital/",
		Tier:   "First",
		TierId: 1,
	},
}
