package ch02

import "testing"

func TestIntersection(t *testing.T) {
	n1 := New(1)
	n1.AppendToTail(2)
	n1.AppendToTail(3)

	n2 := New(1)
	n2.AppendToTail(2)
	n2.AppendToTail(3)

	m := Intersection(n1, n2)
	if m != nil {
		t.Errorf("Two lists don't have intersections\n")
	}

	c := New(3)
	c.AppendToTail(4)
	c.AppendToTail(5)

	n1 = New(1)
	n1.AppendToTail(2)
	n1.AppendToTail(3)

	tail := n1
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = c

	n2 = New(1)
	n2.AppendToTail(2)
	n2.AppendToTail(3)

	tail = n2
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = c

	m = Intersection(n1, n2)
	if m != c {
		t.Errorf("Intersection node expected to be %v, got %v\n", c, m)
	}

	c = New(3)
	c.AppendToTail(4)
	c.AppendToTail(5)

	n1 = New(0)
	n1.AppendToTail(1)
	n1.AppendToTail(2)
	n1.AppendToTail(3)

	tail = New(0)
	tail = n1
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = c

	n2 = New(1)
	n2.AppendToTail(2)
	n2.AppendToTail(3)

	tail = n2
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = c

	m = Intersection(n1, n2)
	if m != c {
		t.Errorf("Intersection node expected to be %v, got %v\n", c, m)
	}

	c = New(3)
	c.AppendToTail(4)
	c.AppendToTail(5)

	n1 = New(1)
	n1.AppendToTail(2)
	n1.AppendToTail(3)

	tail = n1
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = c

	n2 = New(0)
	n2.AppendToTail(1)
	n2.AppendToTail(2)
	n2.AppendToTail(3)

	tail = n2
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = c

	m = Intersection(n1, n2)
	if m != c {
		t.Errorf("Intersection node expected to be %v, got %v\n", c, m)
	}
}
