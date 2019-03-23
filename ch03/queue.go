package ch03

// Queue data structure implementation based on Linked List
type Queue struct {
	first *Node
	last  *Node
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
func (q *Queue) Dequeue() int {
	data := q.first.data
	q.first = q.first.next

	return data
}

// Peek returns the top item from the queue (without removing it)
func (q *Queue) Peek() int {
	return q.first.data
}

// IsEmpty returns true when the queue is empty, false otherwise
func (q *Queue) IsEmpty() bool {
	return q.first == nil
}
