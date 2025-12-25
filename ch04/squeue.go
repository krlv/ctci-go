package ch04

import "fmt"

// SQueue is a queue implementation based on a slice (simple- or slice-queue)
type SQueue[T comparable] []T

// Enqueue adds an item to the beginning of the queue
func (q *SQueue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

// Dequeue removes and returns the item at the front of the queue.
// If the queue is empty it returns the zero value for T and an error.
func (q *SQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return *new(T), fmt.Errorf("queue is empty")
	}

	value := (*q)[0]

	// set the dequeued element to its zero value to help the GC reclaim the referenced object sooner
	var zero T
	(*q)[0] = zero

	*q = (*q)[1:]
	return value, nil
}

// Peek returns the item at the front of the queue without removing it.
// If the queue is empty it returns the zero value for T and an error.
func (q *SQueue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		return *new(T), fmt.Errorf("queue is empty")
	}

	return (*q)[0], nil
}

// IsEmpty reports whether the queue contains no elements.
func (q *SQueue[T]) IsEmpty() bool {
	return len(*q) == 0
}

// Len returns the number of elements currently stored in the queue.
func (q *SQueue[T]) Len() int {
	return len(*q)
}
