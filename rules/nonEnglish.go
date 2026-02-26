package rules

import "unicode"

func containsNonEnglishLetters(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.In(r, unicode.Latin) {
			return true
		}
	}
	return false
}
