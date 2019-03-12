package ch02

import "testing"

func TestPartition(t *testing.T) {
	node := New(3)
	node.AppendToTail(5)
	node.AppendToTail(8)
	node.AppendToTail(5)
	node.AppendToTail(2)
	node.AppendToTail(10)
	node.AppendToTail(1)

	node = Partition(node, 5)
	n := node
	for i, d := range []int{1, 2, 3, 5, 8, 5, 10} {
		if n.data != d {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, d, n.data)
		}
		n = n.next
	}
}
