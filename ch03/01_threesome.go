package ch03

import "errors"

// NthStack structure implements a N stacks
// on top of the flat slice
type NthStack struct {
	n       int   // amount of stacks inside NthStack
	storage []int // the main storage of stacks, flat int slice
	start   []int // start positions for each of `n`th stacks inside the main storage
	len     []int // length for each of `n`th stacks
	cap     []int // capacity for each of `n`th stacks
}

// NewNthStack returns empty initialized NthStack stack
func NewNthStack(n int) *NthStack {
	nth := new(NthStack)

	nth.n = n
	nth.storage = make([]int, nth.n)
	nth.start = make([]int, nth.n)
	nth.len = make([]int, nth.n)
	nth.cap = make([]int, nth.n)

	for i := 0; i < nth.n; i++ {
		// empty storage
		nth.storage[i] = 0
		// every stack has o elements
		nth.len[i] = 0
		// every stack has 1 capacity
		nth.cap[i] = 1
	}

	for i := 0; i < nth.n; i++ {
		nth.start[i] = i
	}

	return nth
}

// Pop returns (and removes) the top item from the `n`th stack
func (nth *NthStack) Pop(n int) (int, error) {
	if nth.IsEmpty(n) {
		return 0, errors.New("Empty Stack Error")
	}

	idx := nth.start[n] + nth.len[n] - 1
	val := nth.storage[idx]
	nth.len[n]--

	return val, nil
}

// Push add an item to the `n`th stack
func (nth *NthStack) Push(d int, n int) {
	// TBD
}

// Peek returns the top item from the `n`th stack (without removing it)
func (nth *NthStack) Peek(n int) (int, error) {
	if nth.IsEmpty(n) {
		return 0, errors.New("empty Stack Error")
	}

	idx := nth.start[n] + nth.len[n] - 1
	return nth.storage[idx], nil
}

// IsEmpty returns true when the `n`th stack is empty, false otherwise
func (nth *NthStack) IsEmpty(n int) bool {
	return nth.len[n] == 0
}
