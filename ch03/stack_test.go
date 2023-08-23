package ch03

import "testing"

func TestNew(t *testing.T) {
	data := 1
	item := New(data)

	if item.data != data {
		t.Errorf("Item data expected to be %d, got %d", data, item.data)
	}
}

func TestStackPop(t *testing.T) {
	stack := new(Stack[int])
	if _, e := stack.Pop(); e == nil {
		t.Error("Empty stack should Pop an error")
	}

	top := New(3)
	stack.top = top
	top.next = New(2)
	top = top.next
	top.next = New(1)

	values := []int{3, 2, 1}
	i := 0
	for stack.top != nil {
		n, err := stack.Pop()
		if err != nil {
			t.Errorf("Stack expected to have at least one node")
		}

		if n != values[i] {
			t.Errorf("Stack %d item expected to be %d, got %d", (i + 1), values[i], n)
		}

		i++
	}
}

func TestStackPush(t *testing.T) {
	stack := new(Stack[int])
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	expected := []int{3, 2, 1}
	i := 0
	for stack.top != nil {
		if stack.top.data != expected[i] {
			t.Errorf("Stack %d item expected to be %d, got %d", (i + 1), expected[i], stack.top.data)
		}

		stack.top = stack.top.next
		i++
	}
}

func TestStackPeek(t *testing.T) {
	stack := new(Stack[int])
	if _, e := stack.Peek(); e == nil {
		t.Error("Empty stack should Peek an error")
	}

	top := New(3)
	stack.top = top
	top.next = New(2)
	top = top.next
	top.next = New(1)

	values := []int{3, 2, 1}
	i := 0
	for stack.top != nil {
		n, err := stack.Peek()
		if err != nil {
			t.Errorf("Stack expected to have at least one node")
		}

		if n != values[i] {
			t.Errorf("Stack %d item expected to be %d, got %d", (i + 1), values[i], n)
		}

		stack.top = stack.top.next
		i++
	}
}

func TestStackIsEmpty(t *testing.T) {
	stack := new(Stack[int])
	if !stack.IsEmpty() {
		t.Error("Empty stack should return false")
	}

	top := New(1)
	stack.top = top

	if stack.IsEmpty() {
		t.Error("Non-empty stack should return true")
	}
}
