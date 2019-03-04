package ch01

import "testing"

func TestIsOneAway(t *testing.T) {
	s1, s2 := "pale", "ple"
	if !IsOneAway(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "ple", "pale"
	if !IsOneAway(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "pales", "pale"
	if !IsOneAway(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "pale", "pales"
	if !IsOneAway(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "pale", "bale"
	if !IsOneAway(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "pale", "bake"
	if IsOneAway(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" are more than one edit away", s1, s2)
	}

	s1, s2 = "bak", "pale"
	if IsOneAway(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" are more than one edit away", s1, s2)
	}
}
