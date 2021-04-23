package ch03

import "errors"

// FixedMultiStack structure implements multiple stacks
// on top of the flat slice
type FixedMultiStack struct {
	n       int   // amount of stacks in MultiStack
	cap     int   // capacity of each n-th stack
	len     []int // length of each n-th stack
	storage []int // the main storage of stacks, flat slice
}

// NewFixedMultiStack returns empty initialized MultiStack
func NewFixedMultiStack(n int, cap int) (*FixedMultiStack, error) {
	if n < 1 {
		return nil, errors.New("number of stacks is less than 1")
	}

	if cap < 1 {
		return nil, errors.New("stack capacity is less than 1")
	}

	stack := new(FixedMultiStack)

	stack.n = n
	stack.cap = cap

	stack.len = make([]int, n)
	stack.storage = make([]int, n*cap)

	return stack, nil
}

// Pop returns (and removes) the top item from the n-th stack
func (stack *FixedMultiStack) Pop(n int) (int, error) {
	if stack.IsEmpty(n) {
		return 0, errors.New("empty stack error")
	}

	top := stack.top(n)
	val := stack.storage[top]
	stack.len[n]--

	return val, nil
}

// Push add an item to the n-th stack
func (stack *FixedMultiStack) Push(n int, val int) error {
	if stack.IsFull(n) {
		return errors.New("full stack error")
	}

	top := stack.top(n)
	stack.storage[top] = val
	stack.len[n]++

	return nil
}

// Peek returns the top item from the n-th stack (without removing it)
func (stack *FixedMultiStack) Peek(n int) (int, error) {
	if stack.IsEmpty(n) {
		return 0, errors.New("empty stack error")
	}

	top := stack.top(n)
	return stack.storage[top], nil
}

// IsEmpty returns true when the n-th stack is empty, false otherwise
func (stack *FixedMultiStack) IsEmpty(n int) bool {
	return stack.len[n] == 0
}

// IsFull returns true when the n-th stack is full, false otherwise
func (stack *FixedMultiStack) IsFull(n int) bool {
	return stack.len[n] == stack.cap
}

// Top returns top element index for n-th stack
// Top element's index calculated as:
// n-th stack offset + top index in n-th stack
func (stack *FixedMultiStack) top(n int) int {
	return n*stack.cap + stack.len[n] - 1
}
