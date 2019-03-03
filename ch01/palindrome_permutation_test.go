package ch01

import "testing"

func TestIsPalindromePermutation(t *testing.T) {
	s := "taco cat"
	if !IsPalindromePermutation(s) {
		t.Errorf("\"%s\" expected to be marked as a palindrome permutation", s)
	}
}
