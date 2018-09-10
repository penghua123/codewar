package Bandname

import (
	"regexp"
	"strings"
)

func BandNameGenerator(word string) string {
	word = strings.ToLower(word)
	reg := regexp.MustCompile(`(\w)(\w*)(\w)`)
	result := reg.FindStringSubmatch(word)
	if result[1] == result[3] {
		return strings.ToUpper(result[1]) + result[2] + result[3] + result[2] + result[3]
	} else {
		return "The " + strings.ToUpper(result[1]) + result[2] + result[3]
	}
	return "Error! There are non character in the string"
}