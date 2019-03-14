package ch02

import (
	"testing"
)

func TestAppendToTail(t *testing.T) {
	node := New(1)
	node.AppendToTail(2)
	node.AppendToTail(3)

	n := node
	for i, d := range []int{1, 2, 3} {
		if n.data != d {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, d, n.data)
		}
		n = n.next
	}
}

func TestDeleteNode(t *testing.T) {
	node := New(1)
	node.AppendToTail(2)
	node.AppendToTail(3)

	node.DeleteNode(2)
	n := node
	for i, d := range []int{1, 3} {
		if n.data != d {
			t.Errorf("%dth node expected to have data %v, got %v\n", i+1, d, n.data)
		}
		n = n.next
	}

	node.DeleteNode(1)
	n = node
	for i, d := range []int{3} {
		if n.data != d {
			t.Errorf("%dth node expected to have data %v, got %v\n", i+1, d, n.data)
		}
		n = n.next
	}

	node.DeleteNode(3)
	n = node
	if n.data != 0 || n.next != nil {
		t.Errorf("Expected to have empty linked list, got %v\n", n)
	}
}

func TestLen(t *testing.T) {
	node := New(1)
	node.AppendToTail(2)
	node.AppendToTail(3)

	actual := node.Len()
	expected := 3
	if actual != expected {
		t.Errorf("List length expected to be %d, got %d", expected, actual)
	}

	node = New(1)
	actual = node.Len()
	expected = 1
	if actual != expected {
		t.Errorf("List length expected to be %d, got %d", expected, actual)
	}

	node = nil
	actual = node.Len()
	expected = 0
	if actual != expected {
		t.Errorf("List length expected to be %d, got %d", expected, actual)
	}
}

func TestPadLeft(t *testing.T) {
	node := New(1)
	node.AppendToTail(2)
	node.AppendToTail(3)

	node.PadLeft(0, 3)
	expected := []int{0, 0, 0, 1, 2, 3}

	for i := 0; node != nil; i++ {
		if node.data != expected[i] {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, expected[i], node.data)
		}
		node = node.next
	}
}
