package solutions

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	var nonAlphaNumRegEx = regexp.MustCompile(`[^a-zA-Z0-9]`)
	cleanString := nonAlphaNumRegEx.ReplaceAllString(s, "")
	cleanString = strings.ReplaceAll(cleanString, " ", "")
	cleanString = strings.ToLower(cleanString)
	reverseCleanString := reverseString(cleanString)
	return cleanString == reverseCleanString
}

func reverseString(s string) string {
	var result = ""
	for _, str := range s {
		result = string(str) + result
	}
	return result
}
