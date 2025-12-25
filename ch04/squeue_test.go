package ch04

import (
	"reflect"
	"testing"
)

func TestSQueue_Dequeue(t *testing.T) {
	type testCase[T comparable] struct {
		name    string
		q       SQueue[T]
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "empty queue",
			q:       SQueue[int]{},
			want:    0,
			wantErr: true,
		},
		{
			name:    "non empty queue",
			q:       SQueue[int]{10, 20, 30},
			want:    10,
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

func TestSQueue_Enqueue(t *testing.T) {
	type args[T comparable] struct {
		value T
	}
	type testCase[T comparable] struct {
		name string
		q    SQueue[T]
		args args[T]
		want SQueue[T]
	}
	tests := []testCase[int]{
		{
			name: "empty queue",
			q:    SQueue[int]{},
			args: args[int]{10},
			want: []int{10},
		},
		{
			name: "non empty queue",
			q:    SQueue[int]{10, 20, 30},
			args: args[int]{40},
			want: []int{10, 20, 30, 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.Enqueue(tt.args.value)
			if !reflect.DeepEqual(tt.q, tt.want) {
				t.Errorf("Enqueue() got = %v, want %v", tt.q, tt.want)
			}
		})
	}
}

func TestSQueue_IsEmpty(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		q    SQueue[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "empty queue",
			q:    SQueue[int]{},
			want: true,
		},
		{
			name: "non empty queue",
			q:    SQueue[int]{10, 20, 30},
			want: false,
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

func TestSQueue_Len(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		q    SQueue[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "empty queue",
			q:    SQueue[int]{},
			want: 0,
		},
		{
			name: "non empty queue",
			q:    SQueue[int]{10, 20, 30},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQueue_Peek(t *testing.T) {
	type testCase[T comparable] struct {
		name    string
		q       SQueue[T]
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "empty queue",
			q:       SQueue[int]{},
			want:    0,
			wantErr: true,
		},
		{
			name:    "non empty queue",
			q:       SQueue[int]{10, 20, 30},
			want:    10,
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
