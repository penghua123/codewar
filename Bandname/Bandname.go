package Bandname

import (
	"regexp"
	"strings"
)

func BandNameGenerator(word string) string {
	//word = strings.ToLower(word)
	reg := regexp.MustCompile(`(\w)(\S*)(\w)`)
	result := reg.FindStringSubmatch(word)
	if result[1] == result[3] {
		return strings.Title(result[1] + result[2] + result[3] + result[2] + result[3])
	} else {
		return "The " + strings.Title(word)
	}
	return "Error! There are non character in the string"
}

func BandNameGenerator2(word string) string {
	//word = strings.ToLower(word)
	if word[0] == word[len(word)-1] {
		return strings.Title(word + word[1:])
	} else {
		return "The " + strings.Title(word)
	}
	return "Error! There are non character in the string"
}
