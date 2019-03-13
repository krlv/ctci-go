package ch02

import "testing"

func TestSumLists(t *testing.T) {
	n1 := New(7)
	n1.AppendToTail(1)
	n1.AppendToTail(6)

	n2 := New(5)
	n2.AppendToTail(9)
	n2.AppendToTail(2)

	sum := SumLists(n1, n2)
	for i, d := range []int{2, 1, 9} {
		if sum.data != d {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, d, sum.data)
		}
		sum = sum.next
	}

	n1 = New(1)
	n1.AppendToTail(1)

	n2 = New(9)
	n2.AppendToTail(9)

	sum = SumLists(n1, n2)
	for i, d := range []int{0, 1, 1} {
		if sum.data != d {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, d, sum.data)
		}
		sum = sum.next
	}
}
