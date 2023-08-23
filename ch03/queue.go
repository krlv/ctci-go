package ch03

import "errors"

// Queue data structure implementation based on Linked List
type Queue[T any] struct {
	first *Item[T]
	last  *Item[T]
}

// Enqueue adds an item to the beggining of the queue
func (q *Queue[T]) Enqueue(data T) {
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
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return *new(T), errors.New("empty Queue error")
	}

	data := q.first.data
	q.first = q.first.next

	return data, nil
}

// Peek returns the top item from the queue (without removing it)
func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		return *new(T), errors.New("empty Queue error")
	}

	return q.first.data, nil
}

// IsEmpty returns true when the queue is empty, false otherwise
func (q *Queue[T]) IsEmpty() bool {
	return q.first == nil
}
