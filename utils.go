package gobikeindex

import "strings"

const (
	twoDigits = 2
)

/*
	Perform quick validations to avoid unecessary HTTP requests
*/

// ValidateStolenness validates entry for stolenness
func ValidateStolenness(s string) bool {
	switch strings.ToLower(s) {
	case "", "stolen", "non", "proximity", "all":
		return true
	default:
		return false
	}
}

// IsState2Digits validates that the US state name is in 2 digit format
func IsState2Digits(s string) bool {
	return len(s) == twoDigits
}
