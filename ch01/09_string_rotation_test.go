package ch01

import "testing"

func TestIsRotation(t *testing.T) {
	s1 := "walterwhite"
	s2 := "erwhitewalt"
	if !IsRotation(s1, s2) {
		t.Errorf("Expected \"%s\" to be rotation of \"%s\"", s1, s2)
	}

	s1 = "walterwhite"
	s2 = "erwalt"
	if IsRotation(s1, s2) {
		t.Errorf("\"%s\" is not rotation of \"%s\"", s1, s2)
	}
}
