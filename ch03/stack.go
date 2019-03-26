package ch03

import "errors"

// Item structure implements an item of the Stack/Queue
// Current implementation is based on Linked List
type Item struct {
	data int
	next *Item
}

// New returns new stack node value data as a top value
func New(data int) *Item {
	item := new(Item)
	item.data = data

	return item
}

// Stack structure implementation based on Linked List
type Stack struct {
	top *Item
}

// Pop returns (and removes) the top item from the stack
func (stack *Stack) Pop() (int, error) {
	if stack.top == nil {
		return 0, errors.New("Empty Stack error")
	}

	top := stack.top
	stack.top = top.next

	return top.data, nil
}

// Push add an item to the stack
func (stack *Stack) Push(data int) {
	top := New(data)
	top.next = stack.top

	stack.top = top
}

// Peek returns the top item from the satck (without removing it)
func (stack *Stack) Peek() (int, error) {
	if stack.top == nil {
		return 0, errors.New("Empty Stack error")
	}

	return stack.top.data, nil
}

// IsEmpty returns true when stack is empty, false otherwise
func (stack *Stack) IsEmpty() bool {
	return stack.top == nil
}
