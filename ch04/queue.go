package ch04

import (
	"errors"
	"fmt"
	"strings"
)

// Node is a linked list node
type Node[T comparable] struct {
	data T
	next *Node[T]
}

// String returns a string representation of the node
func (n *Node[T]) String() string {
	return fmt.Sprintf("%v", n.data)
}

// Queue data structure implementation based on Linked List
type Queue[T comparable] struct {
	first *Node[T]
	last  *Node[T]
}

// Enqueue adds an item to the beginning of the queue
func (q *Queue[T]) Enqueue(data T) {
	node := &Node[T]{data: data}

	if q.last != nil {
		q.last.next = node
	}

	// Same as q.last = q.last.next
	q.last = node

	if q.first == nil {
		q.first = q.last
	}
}

// Dequeue returns (and removes) an item from the queue
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return *new(T), errors.New("queue is empty")
	}

	data := q.first.data
	q.first = q.first.next

	return data, nil
}

// Peek returns the top item from the queue (without removing it)
func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		return *new(T), errors.New("queue is empty")
	}

	return q.first.data, nil
}

// IsEmpty returns true when the queue is empty, false otherwise
func (q *Queue[T]) IsEmpty() bool {
	return q.first == nil
}

// String returns a string representation of the queue
func (q *Queue[T]) String() string {
	if q.IsEmpty() {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")
	for node := q.first; node != nil; node = node.next {
		sb.WriteString(" ")
		sb.WriteString(node.String())
	}
	sb.WriteString(" ]")

	return sb.String()
}
