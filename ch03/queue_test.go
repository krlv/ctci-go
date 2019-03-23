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
		data := queue.Dequeue()

		if data != values[i] {
			t.Errorf("Queue %d item expected to be %d, got %d", (i + 1), values[i], data)
		}

		i++
	}
}

func TestQueuePeek(t *testing.T) {
	queue := new(Queue)

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
		data := queue.Peek()

		if data != values[i] {
			t.Errorf("Queue %d item expected to be %d, got %d", (i + 1), values[i], data)
		}

		queue.first = queue.first.next
		i++
	}
}

func TestQueueIsEmpty(t *testing.T) {
	queue := new(Queue)
	queue.first = New(1)

	if queue.IsEmpty() {
		t.Error("Non-empty queue should return true")
	}
}
