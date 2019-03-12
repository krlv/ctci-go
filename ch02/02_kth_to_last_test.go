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

func TestKthToLastRecursive(t *testing.T) {
	node := New(1)
	node.AppendToTail(2)
	node.AppendToTail(3)
	node.AppendToTail(4)
	node.AppendToTail(5)

	k := 0
	actual, _ := KthToLastRecursive(node, k)
	expected := 5
	if actual.data != expected {
		t.Errorf("%dth element expected to be %v, got %v", k, expected, actual.data)
	}

	k = 2
	actual, _ = KthToLastRecursive(node, k)
	expected = 3
	if actual.data != expected {
		t.Errorf("%dth element expected to be %v, got %v", k, expected, actual.data)
	}

	k = 4
	actual, _ = KthToLastRecursive(node, k)
	expected = 1
	if actual.data != expected {
		t.Errorf("%dth element expected to be %v, got %v", k, expected, actual.data)
	}
}
