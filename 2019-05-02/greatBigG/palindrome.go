package palindrome

import (
	"log"
	"regexp"
	"strings"
)

// reverse takes a string as input, reverses it
// and returns it
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// reverseWord order takes a string as input, reverses the order
// of the words and returns it.
func reverseWordOrder(input string) string {
	s := strings.Split(input, " ")

	last := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[last-i] = s[last-i], s[i]
	}

	return strings.Join(s, " ")
}

// IsPalindrome accepts an input string and returns whether it is a
// palindrome or not.
func IsPalindrome(input string) bool {
	removeAll, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}

	keepSpaces, err := regexp.Compile("[^a-zA-Z\\s]+")
	if err != nil {
		log.Fatal(err)
	}

	pra := strings.ToLower(removeAll.ReplaceAllString(input, ""))
	pks := strings.ToLower(keepSpaces.ReplaceAllString(input, ""))

	if s := strings.Split(input, " "); len(s) > 1 {
		return pra == reverse(pra) || pks == reverseWordOrder(pks)
	}

	return pra == reverse(pra)
}
