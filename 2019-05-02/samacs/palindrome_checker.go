package palindrome_checker

import (
	"regexp"
	"strings"
)

// IsPalindrome checks if the given sentence is a palindrome.
func IsPalindrome(sentence string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9\\s+]+")
	if err != nil {
		panic(err)
	}
	sentence = strings.ToLower(reg.ReplaceAllString(sentence, ""))
	words := strings.Split(sentence, " ")
	length := len(words)
	if length > 1 {
		if words[0] == words[length-1] {
			for i := 1; i < length-1; i++ {
				if i == length-i-1 {
					return true
				}
				if words[i] != words[length-i-1] {
					return false
				}
			}
			return true
		}
	}
	return wordIsPalindrome(sentence)
	// The solution below builds a table of occurrences.
	// Although the time complexity is the same this solution is
	// slightly slower than the `wordIsPalindrome` function.
	// This solution uses a little bit more space also.

	//odd := 0
	//table := make(map[rune]int)
	//a := 'a'
	//z := 'z'
	//for _, char := range sentence {
	//	n := rune(-1)
	//	if char >= a && char <= z {
	//		n = char - a
	//	}
	//	if n != -1 {
	//		table[char]++
	//		if table[char]%2 == 1 {
	//			odd++
	//		} else {
	//			odd--
	//		}
	//	}
	//}
	//return odd <= 1
}

func wordIsPalindrome(word string) bool {
	bits := 0
	a := 'a'
	z := 'z'
	for _, char := range word {
		n := rune(-1)
		if char >= a && char <= z {
			n = char - a
		}
		if n >= 0 {
			mask := 1 << uint(n)
			if (bits & mask) == 0 {
				bits |= mask
			} else {
				bits &= ^mask
			}
		}
	}
	if bits == 0 {
		return true
	}
	return bits&(bits-1) == 0
}
