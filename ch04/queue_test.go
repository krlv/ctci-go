package ch04_test

import (
	"reflect"
	"testing"

	"github.com/krlv/ctci-go/ch04"
)

func TestQueue_Dequeue(t *testing.T) {
	emptyQueue := &ch04.Queue[int]{}

	oneItemQueue := &ch04.Queue[int]{}
	oneItemQueue.Enqueue(1)

	multiItemQueue := &ch04.Queue[int]{}
	multiItemQueue.Enqueue(1)
	multiItemQueue.Enqueue(2)
	multiItemQueue.Enqueue(3)

	type testCase[T comparable] struct {
		name    string
		q       *ch04.Queue[T]
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "empty queue",
			q:       emptyQueue,
			want:    0,
			wantErr: true,
		},
		{
			name:    "one item queue",
			q:       oneItemQueue,
			want:    1,
			wantErr: false,
		},
		{
			name:    "multi item queue",
			q:       multiItemQueue,
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dequeue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	emptyQueue := &ch04.Queue[int]{}

	oneItemQueue := &ch04.Queue[int]{}
	oneItemQueue.Enqueue(1)

	multiItemQueue := &ch04.Queue[int]{}
	multiItemQueue.Enqueue(1)
	multiItemQueue.Enqueue(2)
	multiItemQueue.Enqueue(3)

	type testCase[T comparable] struct {
		name    string
		q       *ch04.Queue[T]
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "empty queue",
			q:       emptyQueue,
			want:    0,
			wantErr: true,
		},
		{
			name:    "one item queue",
			q:       oneItemQueue,
			want:    1,
			wantErr: false,
		},
		{
			name:    "multi item queue",
			q:       multiItemQueue,
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	emptyQueue := &ch04.Queue[int]{}

	peekedQueue := &ch04.Queue[int]{}
	peekedQueue.Enqueue(1)
	peekedQueue.Peek()

	dequeuedQueue := &ch04.Queue[int]{}
	dequeuedQueue.Enqueue(1)
	dequeuedQueue.Dequeue()

	type testCase[T comparable] struct {
		name string
		q    *ch04.Queue[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "empty queue",
			q:    emptyQueue,
			want: true,
		},
		{
			name: "one item queue after peek()",
			q:    peekedQueue,
			want: false,
		},
		{
			name: "one item queue after dequeue()",
			q:    dequeuedQueue,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_String(t *testing.T) {
	emptyQueue := &ch04.Queue[int]{}

	oneItemQueue := &ch04.Queue[int]{}
	oneItemQueue.Enqueue(1)

	peekedQueue := &ch04.Queue[int]{}
	peekedQueue.Enqueue(1)
	peekedQueue.Enqueue(2)
	peekedQueue.Peek()

	dequeuedQueue := &ch04.Queue[int]{}
	dequeuedQueue.Enqueue(1)
	dequeuedQueue.Enqueue(2)
	dequeuedQueue.Enqueue(3)
	dequeuedQueue.Dequeue()

	type testCase[T comparable] struct {
		name string
		q    *ch04.Queue[T]
		want string
	}
	tests := []testCase[int]{
		{
			name: "empty queue",
			q:    emptyQueue,
			want: "[]",
		},
		{
			name: "one item queue",
			q:    oneItemQueue,
			want: "[ 1 ]",
		},
		{
			name: "two item queue after peek()",
			q:    peekedQueue,
			want: "[ 1 2 ]",
		},
		{
			name: "three item queue after dequeue()",
			q:    dequeuedQueue,
			want: "[ 2 3 ]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
