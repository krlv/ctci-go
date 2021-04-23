package ch03

import (
	"reflect"
	"testing"
)

func TestNewFixedMultiStack(t *testing.T) {
	multiStack := FixedMultiStack{}
	multiStack.n = 3
	multiStack.cap = 2
	multiStack.len = make([]int, 3)
	multiStack.storage = make([]int, 6)

	type args struct {
		n   int
		cap int
	}
	tests := []struct {
		name    string
		args    args
		want    *FixedMultiStack
		wantErr bool
	}{
		{
			name:    "creates a new fixed multi stack",
			args:    args{n: 3, cap: 2},
			want:    &multiStack,
			wantErr: false,
		},
		{
			name:    "fails to create new multi stack with zero stacks",
			args:    args{n: 0, cap: 2},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "fails to create new multi stack with negative capacity",
			args:    args{n: 3, cap: -1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFixedMultiStack(tt.args.n, tt.args.cap)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFixedMultiStack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFixedMultiStack() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixedMultiStack_Pop(t *testing.T) {
	type fields struct {
		n       int
		cap     int
		len     []int
		storage []int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name:   "pops element from the first stack",
			fields: fields{n: 1, cap: 3, len: []int{2}, storage: []int{1, 1, 0}},
			args:   args{n: 0},
			want:   1,
		},
		{
			name:   "pops last element from the first stack",
			fields: fields{n: 1, cap: 3, len: []int{1}, storage: []int{1, 0, 0}},
			args:   args{n: 0},
			want:   1,
		},
		{
			name:    "returns errors when stack is empty",
			fields:  fields{n: 1, cap: 3, len: []int{0}, storage: []int{}},
			args:    args{n: 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedMultiStack{
				n:       tt.fields.n,
				cap:     tt.fields.cap,
				len:     tt.fields.len,
				storage: tt.fields.storage,
			}
			got, err := stack.Pop(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixedMultiStack_Push(t *testing.T) {
	storage := []int{1, 2, 0, 1, 2, 0, 1, 2, 0}
	type fields struct {
		n       int
		cap     int
		len     []int
		storage []int
	}
	type args struct {
		n   int
		val int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "pushes element to the first stack",
			fields: fields{n: 3, cap: 3, len: []int{2, 2, 2}, storage: storage},
			args:   args{n: 0, val: 3},
		},
		{
			name:   "pushes element to the second stack",
			fields: fields{n: 3, cap: 3, len: []int{3, 2, 2}, storage: storage},
			args:   args{n: 1, val: 3},
		},
		{
			name:   "pushes element to the second stack",
			fields: fields{n: 3, cap: 3, len: []int{3, 3, 2}, storage: storage},
			args:   args{n: 2, val: 3},
		},
		{
			name:    "returns errors when stack is full",
			fields:  fields{n: 3, cap: 3, len: []int{3, 3, 3}, storage: storage},
			args:    args{n: 2, val: 4},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedMultiStack{
				n:       tt.fields.n,
				cap:     tt.fields.cap,
				len:     tt.fields.len,
				storage: tt.fields.storage,
			}
			if err := stack.Push(tt.args.n, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFixedMultiStack_Peek(t *testing.T) {
	type fields struct {
		n       int
		cap     int
		len     []int
		storage []int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name:   "returns peek for the first stack",
			fields: fields{n: 3, cap: 1, len: []int{1, 0, 0}, storage: []int{1, 0, 0}},
			args:   args{n: 0},
			want:   1,
		},
		{
			name:   "returns peek for the second stack",
			fields: fields{n: 3, cap: 2, len: []int{0, 2, 0}, storage: []int{0, 0, 1, 2, 0, 0}},
			args:   args{n: 1},
			want:   2,
		},
		{
			name:   "returns peek for the second stack",
			fields: fields{n: 3, cap: 3, len: []int{0, 0, 3}, storage: []int{0, 0, 0, 0, 0, 0, 1, 2, 3}},
			args:   args{n: 2},
			want:   3,
		},
		{
			name:    "returns errors when stack is empty",
			fields:  fields{n: 3, cap: 3, len: []int{0, 0, 0}, storage: []int{}},
			args:    args{n: 2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedMultiStack{
				n:       tt.fields.n,
				cap:     tt.fields.cap,
				len:     tt.fields.len,
				storage: tt.fields.storage,
			}
			got, err := stack.Peek(tt.args.n)
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

func TestFixedMultiStack_IsEmpty(t *testing.T) {
	type fields struct {
		n       int
		cap     int
		len     []int
		storage []int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "returns false when non-empty stack",
			fields: fields{n: 3, cap: 3, len: []int{2, 0, 0}, storage: []int{}},
			args:   args{n: 0},
			want:   false,
		},
		{
			name:   "returns true when stack is empty",
			fields: fields{n: 3, cap: 3, len: []int{0, 0, 0}, storage: []int{}},
			args:   args{n: 0},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedMultiStack{
				n:       tt.fields.n,
				cap:     tt.fields.cap,
				len:     tt.fields.len,
				storage: tt.fields.storage,
			}
			if got := stack.IsEmpty(tt.args.n); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixedMultiStack_IsFull(t *testing.T) {
	type fields struct {
		n       int
		cap     int
		len     []int
		storage []int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "returns false for non-full stack",
			fields: fields{n: 3, cap: 3, len: []int{2, 0, 0}, storage: []int{}},
			args:   args{n: 0},
			want:   false,
		},
		{
			name:   "returns true when stack is full",
			fields: fields{n: 3, cap: 3, len: []int{3, 0, 0}, storage: []int{}},
			args:   args{n: 0},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedMultiStack{
				n:       tt.fields.n,
				cap:     tt.fields.cap,
				len:     tt.fields.len,
				storage: tt.fields.storage,
			}
			if got := stack.IsFull(tt.args.n); got != tt.want {
				t.Errorf("IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixedMultiStack_top(t *testing.T) {
	type fields struct {
		n       int
		cap     int
		len     []int
		storage []int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "returns top for the first stack",
			fields: fields{n: 3, cap: 1, len: []int{1, 0, 0}, storage: []int{}},
			args:   args{n: 0},
			want:   0,
		},
		{
			name:   "returns top for the second stack",
			fields: fields{n: 3, cap: 2, len: []int{0, 2, 0}, storage: []int{}},
			args:   args{n: 1},
			want:   3,
		},
		{
			name:   "returns top for the second stack",
			fields: fields{n: 3, cap: 3, len: []int{0, 0, 3}, storage: []int{}},
			args:   args{n: 2},
			want:   8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedMultiStack{
				n:       tt.fields.n,
				cap:     tt.fields.cap,
				len:     tt.fields.len,
				storage: tt.fields.storage,
			}
			if got := stack.top(tt.args.n); got != tt.want {
				t.Errorf("top() = %v, want %v", got, tt.want)
			}
		})
	}
}
