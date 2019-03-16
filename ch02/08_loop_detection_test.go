package ch02

import "testing"

func TestGetLoopHead(t *testing.T) {
	node := New(1)
	tail := node
	for i := 1; i < 5; i++ {
		tail.next = New(i + 1)
		tail = tail.next
	}

	if GetLoopHead(node) != nil {
		t.Errorf("List doesn't have a loop inside\n")
	}

	head := New(6)
	loop := head
	for i := 6; i < 10; i++ {
		loop.next = New(i + 1)
		loop = loop.next
	}

	loop.next = head
	tail.next = head

	actual := GetLoopHead(node)
	if actual != head {
		t.Errorf("Expected %v to be the head of the loop, got %v\n", head, actual)
	}
}
