package ch03

import "errors"

// FixedStack is a capped stack (stack length is limited to cap)
type FixedStack struct {
	top *Item
	len int
	cap int
}

// NewFixedStack desc
func NewFixedStack(cap int) *FixedStack {
	stack := new(FixedStack)
	stack.len = 0
	stack.cap = cap

	return stack
}

// Pop returns (and removes) the top item from the stack
func (stack *FixedStack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("empty stack error")
	}

	top := stack.top
	stack.top = top.next
	stack.len--

	return top.data, nil
}

// Push add an item to the stack
func (stack *FixedStack) Push(data int) error {
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
func (stack *FixedStack) Peek() (int, error) {
	if stack.top == nil {
		return 0, errors.New("empty Stack error")
	}

	return stack.top.data, nil
}

// IsEmpty returns true when stack is empty, false otherwise
func (stack *FixedStack) IsEmpty() bool {
	return stack.top == nil
}

// IsFull returns true when stack is its max capacity, false otherwise
func (stack *FixedStack) IsFull() bool {
	return stack.cap == stack.len
}
