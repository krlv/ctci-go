package ch03

import "errors"

// Item structure implements an item of the Stack/Queue
// Current implementation is based on Linked List
type Item[T any] struct {
	data T
	next *Item[T]
}

// New returns new stack node value data as a top value
func New[T any](data T) *Item[T] {
	item := new(Item[T])
	item.data = data

	return item
}

// Stack structure implementation based on Linked List
type Stack[T any] struct {
	top *Item[T]
}

// Pop returns (and removes) the top item from the stack
func (stack *Stack[T]) Pop() (T, error) {
	if stack.top == nil {
		return *new(T), errors.New("empty Stack error")
	}

	top := stack.top
	stack.top = top.next

	return top.data, nil
}

// Push add an item to the stack
func (stack *Stack[T]) Push(data T) {
	top := New(data)
	top.next = stack.top

	stack.top = top
}

// Peek returns the top item from the satck (without removing it)
func (stack *Stack[T]) Peek() (T, error) {
	if stack.top == nil {
		return *new(T), errors.New("empty Stack error")
	}

	return stack.top.data, nil
}

// IsEmpty returns true when stack is empty, false otherwise
func (stack *Stack[T]) IsEmpty() bool {
	return stack.top == nil
}
