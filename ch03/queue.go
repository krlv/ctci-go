package ch03

import "errors"

// Queue data structure implementation based on Linked List
type Queue struct {
	first *Item
	last  *Item
}

// Enqueue adds an item to the beggining of the queue
func (q *Queue) Enqueue(data int) {
	item := New(data)

	if q.last != nil {
		q.last.next = item
	}

	// Same as q.last = q.last.next
	q.last = item

	if q.first == nil {
		q.first = q.last
	}
}

// Dequeue returns (and removes) an item from the queue
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty Queue error")
	}

	data := q.first.data
	q.first = q.first.next

	return data, nil
}

// Peek returns the top item from the queue (without removing it)
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty Queue error")
	}

	return q.first.data, nil
}

// IsEmpty returns true when the queue is empty, false otherwise
func (q *Queue) IsEmpty() bool {
	return q.first == nil
}
