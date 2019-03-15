package ch02

import "testing"

func TestSumLists(t *testing.T) {
	n1 := New(7)
	n1.AppendToTail(1)
	n1.AppendToTail(6)

	n2 := New(5)
	n2.AppendToTail(9)
	n2.AppendToTail(2)

	expected := []int{2, 1, 9}
	actual := SumLists(n1, n2)

	for i := 0; actual != nil; i++ {
		if actual.data != expected[i] {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, expected[i], actual.data)
		}
		actual = actual.next
	}

	n1 = New(1)
	n1.AppendToTail(1)

	n2 = New(9)
	n2.AppendToTail(9)

	expected = []int{0, 1, 1}
	actual = SumLists(n1, n2)

	for i := 0; actual != nil; i++ {
		if actual.data != expected[i] {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, expected[i], actual.data)
		}
		actual = actual.next
	}

	n1 = New(1)
	n1.AppendToTail(1)

	n2 = New(2)
	n2.AppendToTail(2)

	expected = []int{3, 3}
	actual = SumLists(n1, n2)

	for i := 0; actual != nil; i++ {
		if actual.data != expected[i] {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, expected[i], actual.data)
		}
		actual = actual.next
	}

	n1 = New(1)
	n1.AppendToTail(1)
	n1.AppendToTail(1)

	n2 = New(2)
	n2.AppendToTail(2)

	expected = []int{3, 3, 1}
	actual = SumLists(n1, n2)

	for i := 0; actual != nil; i++ {
		if actual.data != expected[i] {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, expected[i], actual.data)
		}
		actual = actual.next
	}
}

func TestSumListsReverse(t *testing.T) {
	n1 := New(6)
	n1.AppendToTail(1)
	n1.AppendToTail(7)

	n2 := New(2)
	n2.AppendToTail(9)
	n2.AppendToTail(5)

	expected := []int{9, 1, 2}
	actual := SumListsReverse(n1, n2)

	for i := 0; actual != nil; i++ {
		if actual.data != expected[i] {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, expected[i], actual.data)
		}
		actual = actual.next
	}

	n1 = New(1)
	n1.AppendToTail(1)

	n2 = New(9)
	n2.AppendToTail(9)

	expected = []int{1, 1, 0}
	actual = SumListsReverse(n1, n2)

	for i := 0; actual != nil; i++ {
		if actual.data != expected[i] {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, expected[i], actual.data)
		}
		actual = actual.next
	}

	n1 = New(1)
	n1.AppendToTail(1)

	n2 = New(2)
	n2.AppendToTail(2)

	expected = []int{3, 3}
	actual = SumListsReverse(n1, n2)

	for i := 0; actual != nil; i++ {
		if actual.data != expected[i] {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, expected[i], actual.data)
		}
		actual = actual.next
	}

	n1 = New(1)
	n1.AppendToTail(1)
	n1.AppendToTail(1)

	n2 = New(2)
	n2.AppendToTail(2)

	expected = []int{1, 3, 3}
	actual = SumListsReverse(n1, n2)

	for i := 0; actual != nil; i++ {
		if actual.data != expected[i] {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, expected[i], actual.data)
		}
		actual = actual.next
	}

	n1 = New(1)
	n1.AppendToTail(1)
	n1.AppendToTail(1)

	n2 = New(2)
	n2.AppendToTail(2)
	n2.AppendToTail(2)
	n2.AppendToTail(2)

	expected = []int{2, 3, 3, 3}
	actual = SumListsReverse(n1, n2)

	for i := 0; actual != nil; i++ {
		if actual.data != expected[i] {
			t.Errorf("%dth node expected to have data %v, got %v", i+1, expected[i], actual.data)
		}
		actual = actual.next
	}
}
