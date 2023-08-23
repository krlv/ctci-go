package ch03

import (
	"errors"
)

// StackSet represents set of fixed (capped) stacks
type StackSet struct {
	stacks []*FixedStack[int]
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
func (set *StackSet) topStack() (*FixedStack[int], error) {
	n := len(set.stacks) - 1
	if n < 0 {
		return nil, errors.New("set is empty")
	}

	return set.stacks[n], nil
}

// AppendStack adds a new stack to the set
func (set *StackSet) appendStack() *FixedStack[int] {
	stack := NewFixedStack[int](3)
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
