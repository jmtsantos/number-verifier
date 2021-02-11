package util

import (
	"regexp"
)

// NumberParser parses a number in the "%s %s" format
func NumberParser(formatString string) string {
	r := regexp.MustCompile(`([0-9]+)`)
	return r.FindStringSubmatch(formatString)[1]
}
