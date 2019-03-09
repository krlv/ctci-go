package ch02

import (
	"testing"
)

func TestAppendToTail(t *testing.T) {
	n := New(1)
	n.AppendToTail(2)
	n.AppendToTail(3)

	for i, v := range []int{1, 2, 3} {
		if n.value != v {
			t.Errorf("%dth node expected to have value %v, got %v", i+1, v, n.value)
		}
		n = n.next
	}
}
