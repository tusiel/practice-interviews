package palindrome

import (
	"strings"
)

func IsPalindrome(s string) bool {
	splitString := strings.Split(s, " ")
	splitString = removeCharactersInSlice(splitString)
	s = strings.ToLower(s)
	s = removeCharacters(s)
	if len(splitString) == 1 {
		return isStringPalindrome(s)
	}
	return isStringPalindrome(s) || isSlicePalindrome(splitString)
}

func isStringPalindrome(s string) bool {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		if s[i] != s[opp] {
			return false
		}
	}
	return true
}

func isSlicePalindrome(s []string) bool {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		if s[i] != s[opp] {
			return false
		}
	}
	return true
}


func removeCharacters(s string) (ret string) {
	for _, c := range s {
		if c > 96 && c < 123 {
			ret = ret + string(c)
		}
	}
	return ret
}

func removeCharactersInSlice(s []string) []string {
	for i := range s {
		s[i] = removeCharacters(strings.ToLower(s[i]))
	}
	return s
}