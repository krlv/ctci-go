package ch03

import "testing"

func TestQueueEnqueue(t *testing.T) {
	queue := new(Queue)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	values := []int{1, 2, 3}
	i := 0

	for queue.first != nil {
		data := queue.first.data

		if data != values[i] {
			t.Errorf("Queue %d item expected to be %d, got %d", (i + 1), values[i], data)
		}

		queue.first = queue.first.next
		i++
	}
}

func TestQueueDequeue(t *testing.T) {
	queue := new(Queue)
	if _, e := queue.Dequeue(); e == nil {
		t.Error("Dequeue on empty queue should produce an error")
	}

	last := New(1)
	queue.last = last
	queue.first = queue.last

	last.next = New(2)
	last = last.next
	queue.last = last

	last.next = New(3)
	last = last.next
	queue.last = last

	values := []int{1, 2, 3}
	i := 0

	for queue.first != nil {
		data, err := queue.Dequeue()
		if err != nil {
			t.Errorf("Queue expected to have at least one item")
		}

		if data != values[i] {
			t.Errorf("Queue %d item expected to be %d, got %d", (i + 1), values[i], data)
		}

		i++
	}
}

func TestQueuePeek(t *testing.T) {
	queue := new(Queue)
	if _, e := queue.Peek(); e == nil {
		t.Error("Peek on empty queue should produce an error")
	}

	last := New(1)
	queue.last = last
	queue.first = queue.last

	last.next = New(2)
	last = last.next
	queue.last = last

	last.next = New(3)
	last = last.next
	queue.last = last

	values := []int{1, 2, 3}
	i := 0

	for queue.first != nil {
		data, err := queue.Peek()
		if err != nil {
			t.Errorf("Queue expected to have at least one item")
		}

		if data != values[i] {
			t.Errorf("Queue %d item expected to be %d, got %d", (i + 1), values[i], data)
		}

		queue.first = queue.first.next
		i++
	}
}

func TestQueueIsEmpty(t *testing.T) {
	queue := new(Queue)
	if !queue.IsEmpty() {
		t.Error("Empty queue should return true")
	}

	queue.first = New(1)

	if queue.IsEmpty() {
		t.Error("Non-empty queue should return true")
	}
}
