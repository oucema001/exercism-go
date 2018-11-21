package isogram

import (
	"unicode"
)

//IsIsogram : checks if word is isogram
func IsIsogram(input string) bool {
	runes := make(map[rune]bool)
	for _, c := range input {
		if c == '-' || c == ' ' {
			continue
		}
		if runes[c]  {
			return false
		}
		c = unicode.ToUpper(c)
		runes[c] = true
	}
	return true
}
