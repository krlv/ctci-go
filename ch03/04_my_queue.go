package ch03

// MyQueue data structure implementation based on two stacks
type MyQueue struct {
	// front of the queue (oldest elements)
	front *Stack

	// back of the queue (newest elements)
	back *Stack
}

// NewMyQueue returns new MyQueue
func NewMyQueue() *MyQueue {
	return &MyQueue{
		front: new(Stack),
		back:  new(Stack),
	}
}

// Enqueue adds an item to the queue
// Time complexity: O(1)
func (q *MyQueue) Enqueue(data int) {
	q.back.Push(data)
}

// Dequeue returns (and removes) an item from the queue
// Time complexity: O(n)
func (q *MyQueue) Dequeue() (int, error) {
	if err := q.shift(); err != nil {
		return 0, err
	}

	return q.front.Pop()
}

// Peek returns the top item from the queue (without removing it)
// Time complexity: O(n)
func (q *MyQueue) Peek() (int, error) {
	if err := q.shift(); err != nil {
		return 0, err
	}

	return q.front.Peek()
}

// IsEmpty returns true when the queue is empty, false otherwise
func (q *MyQueue) IsEmpty() bool {
	return q.front.IsEmpty() && q.back.IsEmpty()
}

// shift moves all elements from back to front
func (q *MyQueue) shift() error {
	// already have elements in the front, no need to shift
	if !q.front.IsEmpty() {
		return nil
	}

	// shift all elements from back to front
	for !q.back.IsEmpty() {
		data, err := q.back.Pop()
		if err != nil {
			return nil
		}

		q.front.Push(data)
	}

	return nil
}
