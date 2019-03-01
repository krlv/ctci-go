package ch01

import "testing"

func TestSortString(t *testing.T) {
	actual, expected := SortString("cabd"), "abcd"
	if actual != expected {
		t.Errorf("Expected \"%s\" string, got \"%s\"", actual, expected)
	}
}
