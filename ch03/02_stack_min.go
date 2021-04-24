package ch03

import "errors"

// StackMin is stack that returns minimum stored value
// in O(1) time complexity
type StackMin struct {
	top *MinItem
}

// MinItem is linked list node that stores int data
// pointer to the next node and current stack minimum value
type MinItem struct {
	data int      // node value
	min  int      // minimal stack value
	next *MinItem // next node pointer
}

// Pop returns (and removes) the top item from the stack
func (stack *StackMin) Pop() (int, error) {
	if stack.top == nil {
		return 0, errors.New("empty Stack error")
	}

	top := stack.top
	stack.top = top.next

	return top.data, nil
}

// Push adds an item to the stack
// If new value is less than previous stack minimum,
// we have to use it as min for the new stack item.
// If new value is bigger than previous stack minimum,
// we will reuse previous minimum value.
func (stack *StackMin) Push(data int) {
	top := new(MinItem)
	top.data = data
	top.next = stack.top

	top.min = data
	if stack.top != nil && stack.top.min < data {
		top.min = stack.top.min
	}

	stack.top = top
}

// Peek returns the top item from the stack (without removing it)
func (stack *StackMin) Peek() (int, error) {
	if stack.top == nil {
		return 0, errors.New("empty Stack error")
	}

	return stack.top.data, nil
}

// IsEmpty returns true when stack is empty, false otherwise
func (stack *StackMin) IsEmpty() bool {
	return stack.top == nil
}

// Min returns minimum value in the stack in O(1) time
func (stack *StackMin) Min() (int, error) {
	if stack.top == nil {
		return 0, errors.New("empty Stack error")
	}

	return stack.top.min, nil
}
