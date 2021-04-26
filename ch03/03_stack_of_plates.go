package ch03

import (
	"errors"
)

// StackSet represents set of fixed (capped) stacks
type StackSet struct {
	stacks []*FixedStack
}

// Pop returns (and removes) the top item from the top stack
// If the top stack is empty, it will be removed from the set
func (set *StackSet) Pop() (int, error) {
	stack, err := set.topStack()
	if err != nil {
		return 0, err
	}

	value, err := stack.Pop()
	if err != nil {
		return 0, err
	}

	if stack.IsEmpty() {
		set.removeStack()
	}

	return value, nil
}

// Push add an item to the top stack
// If the top stack is full, new stack will be created
// and appended to the set
func (set *StackSet) Push(data int) error {
	stack, err := set.topStack()
	if err != nil {
		stack = set.appendStack()
	}

	if stack.IsFull() {
		stack = set.appendStack()
	}

	return stack.Push(data)
}

// TopStack returns the top stack from the set
func (set *StackSet) topStack() (*FixedStack, error) {
	n := len(set.stacks) - 1
	if n < 0 {
		return nil, errors.New("set is empty")
	}

	return set.stacks[n], nil
}

// AppendStack adds a new stack to the set
func (set *StackSet) appendStack() *FixedStack {
	stack := NewFixedStack(3)
	set.stacks = append(set.stacks, stack)

	return stack
}

// RemoveStack removes top stack from the set
func (set *StackSet) removeStack() {
	n := len(set.stacks) - 1
	if n >= 0 {
		set.stacks = set.stacks[:n]
	}
}

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
