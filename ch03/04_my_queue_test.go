package ch03

import "testing"

func TestMyQueue_Enqueue(t *testing.T) {
	newQueue := NewMyQueue()

	emptyQueue := NewMyQueue()
	emptyQueue.Enqueue(1)
	emptyQueue.Dequeue()

	nonEmptyQueue := NewMyQueue()
	nonEmptyQueue.Enqueue(1)
	nonEmptyQueue.Enqueue(2)
	nonEmptyQueue.Enqueue(3)

	peekedQueue := NewMyQueue()
	peekedQueue.Enqueue(1)
	peekedQueue.Enqueue(2)
	peekedQueue.Enqueue(3)
	peekedQueue.Peek()

	type args struct {
		data int
	}
	tests := []struct {
		name  string
		queue *MyQueue
		args  args
	}{
		{
			name:  "adds an item to the new queue",
			queue: newQueue,
			args:  args{data: 1},
		},
		{
			name:  "adds an item to the empty (dequeued) queue",
			queue: emptyQueue,
			args:  args{data: 1},
		},
		{
			name:  "adds an item to the non-empty queue",
			queue: nonEmptyQueue,
			args:  args{data: 4},
		},
		{
			name:  "adds an item to the peeked queue",
			queue: peekedQueue,
			args:  args{data: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.queue.Enqueue(tt.args.data)
		})
	}
}

func TestMyQueue_Dequeue(t *testing.T) {
	newQueue := NewMyQueue()

	emptyQueue := NewMyQueue()
	emptyQueue.Enqueue(1)
	emptyQueue.Dequeue()

	nonEmptyQueue := NewMyQueue()
	nonEmptyQueue.Enqueue(1)
	nonEmptyQueue.Enqueue(2)
	nonEmptyQueue.Enqueue(3)

	peekedQueue := NewMyQueue()
	peekedQueue.Enqueue(1)
	peekedQueue.Enqueue(2)
	peekedQueue.Enqueue(3)
	peekedQueue.Peek()

	tests := []struct {
		name    string
		queue   *MyQueue
		want    int
		wantErr bool
	}{
		{
			name:    "returns error for new queue",
			queue:   newQueue,
			wantErr: true,
		},
		{
			name:    "returns error when queue (dequeued) is empty",
			queue:   emptyQueue,
			wantErr: true,
		},
		{
			name:  "returns the top value from the queue",
			queue: nonEmptyQueue,
			want:  1,
		},
		{
			name:  "returns the top value from the peeked queue",
			queue: peekedQueue,
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.queue.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Dequeue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_IsEmpty(t *testing.T) {
	newQueue := NewMyQueue()

	emptyQueue := NewMyQueue()
	emptyQueue.Enqueue(1)
	emptyQueue.Dequeue()

	nonEmptyQueue := NewMyQueue()
	nonEmptyQueue.Enqueue(1)

	tests := []struct {
		name  string
		queue *MyQueue
		want  bool
	}{
		{
			name:  "returns true for new queue",
			queue: newQueue,
			want:  true,
		},
		{
			name:  "returns true when queue is empty",
			queue: emptyQueue,
			want:  true,
		},
		{
			name:  "returns false when queue is not empty",
			queue: nonEmptyQueue,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.queue.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_Peek(t *testing.T) {
	newQueue := NewMyQueue()

	emptyQueue := NewMyQueue()
	emptyQueue.Enqueue(1)
	emptyQueue.Dequeue()

	nonEmptyQueue := NewMyQueue()
	nonEmptyQueue.Enqueue(1)
	nonEmptyQueue.Enqueue(2)
	nonEmptyQueue.Enqueue(3)

	tests := []struct {
		name    string
		queue   *MyQueue
		want    int
		wantErr bool
	}{
		{
			name:    "returns error for new queue",
			queue:   newQueue,
			wantErr: true,
		},
		{
			name:    "returns error when queue is empty",
			queue:   emptyQueue,
			wantErr: true,
		},
		{
			name:  "returns the top value from the queue",
			queue: nonEmptyQueue,
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.queue.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Peek() got = %v, want %v", got, tt.want)
			}
		})
	}
}
