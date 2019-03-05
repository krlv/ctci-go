package ch01

import (
	"testing"
)

func TestCompress(t *testing.T) {
	example := "aabcccccaaa"
	expected := "a2b1c5a3"
	actual := Compress(example)
	if actual != expected {
		t.Errorf("Expected to compress string \"%s\" to \"%s\", got \"%s\"", example, expected, actual)
	}

	example = "abca"
	actual = Compress(example)
	if actual != example {
		t.Errorf("Expected to compress string \"%s\" to \"%s\", got \"%s\"", example, example, actual)
	}
}
