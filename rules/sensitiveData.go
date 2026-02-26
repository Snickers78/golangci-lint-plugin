package rules

import (
	"regexp"
)

var (
	emailRegexp = regexp.MustCompile(`(?i)[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}`)
	ccRegexp    = regexp.MustCompile(`\b(?:\d[ -]*?){13,19}\b`)
	ipRegexp    = regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)
)

func containsSensitiveData(s string) (bool, string) {

	if emailRegexp.MatchString(s) {
		return true, "contains email address"
	}

	if ccRegexp.MatchString(s) {
		return true, "contains possible credit card number"
	}

	if ipRegexp.MatchString(s) {
		return true, "contains IP address"
	}

	for _, re := range customSensitiveRegexps {
		if re.MatchString(s) {
			return true, "matches custom sensitive pattern " + re.String()
		}
	}

	return false, ""
}
