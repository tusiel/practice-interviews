package palindrome

import "testing"

func TestReverse(t *testing.T) {
	testCases := []struct {
		Input  string
		Output string
	}{
		{Input: "This is a sentence", Output: "ecnetnes a si sihT"},
		{Input: "WORD", Output: "DROW"},
	}

	for _, tc := range testCases {
		if res := reverse(tc.Input); res != tc.Output {
			t.Errorf("Expected %[1]s to return %[2]s, but got %[3]s", tc.Input, tc.Output, res)
		}
	}
}

func TestReverseWordOrder(t *testing.T) {
	testCases := []struct {
		Input  string
		Output string
	}{
		{Input: "This is a sentence", Output: "sentence a is This"},
		{Input: "We should consider reversing this decision", Output: "decision this reversing consider should We"},
	}

	for _, tc := range testCases {
		if res := reverseWordOrder(tc.Input); res != tc.Output {
			t.Errorf("Expected %[1]s to return %[2]s, but got %[3]s", tc.Input, tc.Output, res)
		}
	}
}

func TestSingleWords(t *testing.T) {
	testCases := []struct {
		Word   string
		Result bool
	}{
		{Word: "Mom", Result: true},
		{Word: "redivider", Result: true},
		{Word: "deified", Result: true},
		{Word: "civiC", Result: true},
		{Word: "radar", Result: true},
		{Word: "leVEl", Result: true},
		{Word: "rotor", Result: true},
		{Word: "kAyak", Result: true},
		{Word: "reviver", Result: true},
		{Word: "racecar", Result: true},
		{Word: "reDer", Result: true},
		{Word: "madam", Result: true},
		{Word: "Refer", Result: true},

		{Word: "word", Result: false},
		{Word: "UPPER", Result: false},
		{Word: "MiXturE", Result: false},
	}

	for _, tc := range testCases {
		if res := IsPalindrome(tc.Word); res != tc.Result {
			t.Errorf("Expected %[1]s to be %[2]t, but got %[3]t", tc.Word, tc.Result, res)
		}
	}
}

func TestSentences(t *testing.T) {
	testCases := []struct {
		Sentence string
		Result   bool
	}{
		// Short tests
		{Sentence: "A nut for a jar of tuna.", Result: true},
		{Sentence: "Al lets Della call Ed “Stella.", Result: true},
		{Sentence: "Amore, Roma.", Result: true},
		{Sentence: "Borrow or rob?", Result: true},
		{Sentence: "Taco cat", Result: true},
		{Sentence: "Was it a car or a cat I saw?", Result: true},
		{Sentence: "Ed, I saw Harpo Marx ram Oprah W. aside.", Result: true},
		{Sentence: "Madam, in Eden, I’m Adam.", Result: true},
		{Sentence: "Murder for a jar of red rum.", Result: true},
		{Sentence: "Oozy rat in a sanitary zoo.", Result: true},
		{Sentence: "Yo, banana boy!", Result: true},
		{Sentence: "UFO tofu?", Result: true},
		// Long tests
		{Sentence: "Are we not pure? “No, sir!” Panama’s moody Noriega brags. “It is garbage!” Irony dooms a man—a prisoner up to new era.", Result: true},
		{Sentence: "Dennis, Nell, Edna, Leon, Nedra, Anita, Rolf, Nora, Alice, Carol, Leo, Jane, Reed, Dena, Dale, Basil, Rae, Penny, Lana, Dave, Denny, Lena, Ida, Bernadette, Ben, Ray, Lila, Nina, Jo, Ira, Mara, Sara, Mario, Jan, Ina, Lily, Arne, Bette, Dan, Reba, Diane, Lynn, Ed, Eva, Dana, Lynne, Pearl, Isabel, Ada, Ned, Dee, Rena, Joel, Lora, Cecil, Aaron, Flora, Tina, Arden, Noel, and Ellen sinned.", Result: true},

		// Palindromes can also be by words, instead of characters..
		{Sentence: "King, are you glad you are king?", Result: true},

		{Sentence: "This is not a valid palindrome", Result: false},
		{Sentence: "Borrow or robc?", Result: false},
	}

	for _, tc := range testCases {
		if res := IsPalindrome(tc.Sentence); res != tc.Result {
			t.Errorf("Expected %[1]s to be %[2]t, but got %[3]t", tc.Sentence, tc.Result, res)
		}
	}
}
