package ch01

import "testing"

func TestUrlifyn(t *testing.T) {
	actual, expected := URLifyn("slug url", 8), "slug%20url"
	if actual != expected {
		t.Errorf("Expected \"%s\" after ULRification, got \"%s\"", expected, actual)
	}
}

func TestUrlify(t *testing.T) {
	actual, expected := URLify("slug url"), "slug%20url"
	if actual != expected {
		t.Errorf("Expected \"%s\" after ULRification, got \"%s\"", expected, actual)
	}
}
