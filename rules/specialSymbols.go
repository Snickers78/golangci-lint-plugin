package rules

import "unicode"

func containsSpecialSymbolsOrEmoji(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			continue
		}
		return true
	}
	return false
}
