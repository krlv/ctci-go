package ch02

import "testing"

func TestRemoveDups(t *testing.T) {
	node := New(1)
	node.AppendToTail(1)
	node.AppendToTail(2)
	node.AppendToTail(2)
	node.AppendToTail(2)
	node.AppendToTail(3)

	RemoveDups(node)
	n := node
	for i, d := range []int{1, 2, 3} {
		if n.data != d {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, d, n.data)
		}
		n = n.next
	}
}

func TestRemoveDupsRunner(t *testing.T) {
	node := New(1)
	node.AppendToTail(1)
	node.AppendToTail(2)
	node.AppendToTail(2)
	node.AppendToTail(2)
	node.AppendToTail(3)

	RemoveDupsRunner(node)
	n := node
	for i, d := range []int{1, 2, 3} {
		if n.data != d {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, d, n.data)
		}
		n = n.next
	}
}
