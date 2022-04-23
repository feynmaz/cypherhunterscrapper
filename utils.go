package cypherhunterscrapper

import (
	"errors"
	"regexp"
	"strings"
)

// ValidateUrl checks if input url is valid cypherhunter.com url and trims it using https://pkg.go.dev/strings#TrimSpace.
//
// Returns trimmed url and nil error if url is valid else nil and error
func ValidateUrl(rawUrl string) (string, error) {
	url := strings.TrimSpace(rawUrl)
	if len(url) != 0 {
		template := `^\s*(https:\/\/www.cypherhunter.com\/en\/p\/).+(\/)?\s*$`
		re, _ := regexp.Compile(template)

		if re.Match([]byte(url)) {
			return url, nil
		}
		return "", errors.New(`input does not match pattern "https://www.cypherhunter.com/en/p/coin_name/"`)
	}
	return "", errors.New(`input is empty`)
}
