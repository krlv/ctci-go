package ch02

import (
	"testing"
)

func TestDeleteMiddleNode(t *testing.T) {
	node := New(1)
	node.AppendToTail(2)
	node.AppendToTail(3)
	node.AppendToTail(4)
	node.AppendToTail(5)

	m := node
	for i := 0; i < 2; i++ {
		m = m.next
	}

	DeleteMiddleNode(m)
	n := node
	for i, d := range []int{1, 2, 4, 5} {
		if n.data != d {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, d, n.data)
		}
		n = n.next
	}
}
