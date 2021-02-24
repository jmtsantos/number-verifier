// https://github.com/rainycape/countries

package util

import (
	"strings"

	"github.com/rainycape/countries"
)

// ParseCountry parses a country
func ParseCountry(countryStr string) (string, error) {
	var (
		country *countries.Country
		err     error
	)

	if strings.ToLower(countryStr) == "unitedkingdom" {
		countryStr = "United Kingdom"
	}

	if country, err = countries.Parse(countryStr); err != nil {
		return "", err
	}

	return country.Name, err
}
