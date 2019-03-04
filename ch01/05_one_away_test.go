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

func TestIsOneAwayLoop(t *testing.T) {
	s1, s2 := "pale", "ple"
	if !IsOneAwayLoop(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "ple", "pale"
	if !IsOneAwayLoop(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "pales", "pale"
	if !IsOneAwayLoop(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "pale", "pales"
	if !IsOneAwayLoop(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "pale", "bale"
	if !IsOneAwayLoop(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "pale", "bake"
	if IsOneAwayLoop(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" are more than one edit away", s1, s2)
	}

	s1, s2 = "bak", "pale"
	if IsOneAwayLoop(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" are more than one edit away", s1, s2)
	}
}

func TestIsOneAwayClosure(t *testing.T) {
	s1, s2 := "pale", "ple"
	if !IsOneAwayClosure(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "ple", "pale"
	if !IsOneAwayClosure(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "pales", "pale"
	if !IsOneAwayClosure(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "pale", "pales"
	if !IsOneAwayClosure(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "pale", "bale"
	if !IsOneAwayClosure(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" considered to be one edit away", s1, s2)
	}

	s1, s2 = "pale", "bake"
	if IsOneAwayClosure(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" are more than one edit away", s1, s2)
	}

	s1, s2 = "bak", "pale"
	if IsOneAwayClosure(s1, s2) {
		t.Errorf("Strings \"%s\" and \"%s\" are more than one edit away", s1, s2)
	}
}
