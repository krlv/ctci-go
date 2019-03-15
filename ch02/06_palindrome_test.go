package ch02

import "testing"

func TestIsPalindrome(t *testing.T) {
	node := New(1)
	node.AppendToTail(2)
	node.AppendToTail(3)
	node.AppendToTail(4)

	if IsPalindrome(node) {
		t.Error("Expected the linked list to be not a polindrome")
	}

	node = New(1)
	node.AppendToTail(2)
	node.AppendToTail(3)
	node.AppendToTail(2)
	node.AppendToTail(1)

	if !IsPalindrome(node) {
		t.Error("Expected the linked list to be a polindrome")
	}
}
