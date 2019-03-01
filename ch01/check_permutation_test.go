package ch01

import "testing"

func TestCheckPermutation(t *testing.T) {
	s1, s2 := "dog", "god"
	if !CheckPermutation(s1, s2) {
		t.Errorf("Expected strings \"%s\" and \"%s\" to be permutation of each other", s1, s2)
	}

	s1, s2 = "man", "god"
	if CheckPermutation(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" aren't permutation of each other", s1, s2)
	}

	s1, s2 = "man", "human"
	if CheckPermutation(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" aren't permutation of each other", s1, s2)
	}
}
