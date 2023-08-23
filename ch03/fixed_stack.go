package ch03

import "errors"

// FixedStack is a capped stack (stack length is limited to cap)
type FixedStack[T any] struct {
	top *Item[T]
	len int
	cap int
}

// NewFixedStack desc
func NewFixedStack[T any](cap int) *FixedStack[T] {
	stack := new(FixedStack[T])
	stack.len = 0
	stack.cap = cap

	return stack
}

// Pop returns (and removes) the top item from the stack
func (stack *FixedStack[T]) Pop() (T, error) {
	if stack.IsEmpty() {
		return *new(T), errors.New("empty stack error")
	}

	top := stack.top
	stack.top = top.next
	stack.len--

	return top.data, nil
}

// Push add an item to the stack
func (stack *FixedStack[T]) Push(data T) error {
	if stack.IsFull() {
		return errors.New("full stack error")
	}

	top := New(data)
	top.next = stack.top

	stack.top = top
	stack.len++

	return nil
}

// Peek returns the top item from the satck (without removing it)
func (stack *FixedStack[T]) Peek() (T, error) {
	if stack.top == nil {
		return *new(T), errors.New("empty Stack error")
	}

	return stack.top.data, nil
}

// IsEmpty returns true when stack is empty, false otherwise
func (stack *FixedStack[T]) IsEmpty() bool {
	return stack.top == nil
}

// IsFull returns true when stack is its max capacity, false otherwise
func (stack *FixedStack[T]) IsFull() bool {
	return stack.cap == stack.len
}
