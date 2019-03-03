package ch01

import "testing"

func TestCheckPermutationBySort(t *testing.T) {
	s1, s2 := "dog", "god"
	if !CheckPermutationBySort(s1, s2) {
		t.Errorf("Expected strings \"%s\" and \"%s\" to be permutation of each other", s1, s2)
	}

	s1, s2 = "man", "god"
	if CheckPermutationBySort(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" aren't permutation of each other", s1, s2)
	}

	s1, s2 = "man", "human"
	if CheckPermutationBySort(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" aren't permutation of each other", s1, s2)
	}
}

func TestCheckPermutationBySet(t *testing.T) {
	s1, s2 := "dog", "god"
	if !CheckPermutationBySet(s1, s2) {
		t.Errorf("Expected strings \"%s\" and \"%s\" to be permutation of each other", s1, s2)
	}

	s1, s2 = "man", "god"
	if CheckPermutationBySet(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" aren't permutation of each other", s1, s2)
	}

	s1, s2 = "man", "human"
	if CheckPermutationBySet(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" aren't permutation of each other", s1, s2)
	}
}
