package ch02

import "testing"

func TestKthToLast(t *testing.T) {
	node := New(1)
	node.AppendToTail(2)
	node.AppendToTail(3)
	node.AppendToTail(4)
	node.AppendToTail(5)

	k := 0
	actual := KthToLast(node, k).data
	expected := 5
	if actual != expected {
		t.Errorf("%dth element expected to be %v, got %v", k, expected, actual)
	}

	k = 2
	actual = KthToLast(node, k).data
	expected = 3
	if actual != expected {
		t.Errorf("%dth element expected to be %v, got %v", k, expected, actual)
	}

	k = 4
	actual = KthToLast(node, k).data
	expected = 1
	if actual != expected {
		t.Errorf("%dth element expected to be %v, got %v", k, expected, actual)
	}
}
