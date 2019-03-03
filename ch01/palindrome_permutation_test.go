package ch01

import "testing"

func TestIsPalindromePermutation(t *testing.T) {
	s := "taco cat"
	if !IsPalindromePermutation(s) {
		t.Errorf("\"%s\" expected to be marked as a palindrome permutation", s)
	}

	s = "tacos cat"
	if IsPalindromePermutation(s) {
		t.Errorf("\"%s\" expected not to be marked as a palindrome permutation", s)
	}
}

func TestIsPalindromePermutationSimplified(t *testing.T) {
	s := "taco cat"
	if !IsPalindromePermutationSimplified(s) {
		t.Errorf("\"%s\" expected to be marked as a palindrome permutation", s)
	}

	s = "tacos cat"
	if IsPalindromePermutationSimplified(s) {
		t.Errorf("\"%s\" expected not to be marked as a palindrome permutation", s)
	}
}
