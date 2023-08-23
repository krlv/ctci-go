package ch03

import (
	"reflect"
	"testing"
)

func TestNewFixedStack(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want *FixedStack[int]
	}{
		{
			name: "creates an empty stack with predefined capacity",
			args: args{cap: 10},
			want: &FixedStack[int]{cap: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFixedStack[int](tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFixedStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixedStack_Pop(t *testing.T) {
	type fields struct {
		top *Item[int]
		len int
		cap int
	}
	type want struct {
		top   int
		stack *FixedStack[int]
	}
	tests := []struct {
		name    string
		fields  fields
		want    want
		wantErr bool
	}{
		{
			name:    "returns error on empty stack",
			fields:  fields{},
			wantErr: true,
		},
		{
			name:   "pops top value from the stack",
			fields: fields{top: &Item[int]{data: 10}, len: 1, cap: 1},
			want:   want{top: 10, stack: &FixedStack[int]{len: 0, cap: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedStack[int]{
				top: tt.fields.top,
				len: tt.fields.len,
				cap: tt.fields.cap,
			}
			got, err := stack.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if got != tt.want.top {
				t.Errorf("Pop() got = %v, want %v", got, tt.want.top)
			}
			if !reflect.DeepEqual(stack, tt.want.stack) {
				t.Errorf("Pop() stock = %v, want %v", stack, tt.want.stack)
			}
		})
	}
}

func TestFixedStack_Push(t *testing.T) {
	type fields struct {
		top *Item[int]
		len int
		cap int
	}
	type args struct {
		data int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FixedStack[int]
		wantErr bool
	}{
		{
			name:   "pushes new value to an empty stack",
			fields: fields{top: nil, len: 0, cap: 2},
			args:   args{data: 1},
			want:   &FixedStack[int]{top: &Item[int]{data: 1}, len: 1, cap: 2},
		},
		{
			name:   "pushes new value to a non empty stack",
			fields: fields{top: &Item[int]{data: 1}, len: 1, cap: 2},
			args:   args{data: 2},
			want:   &FixedStack[int]{top: &Item[int]{data: 2, next: &Item[int]{data: 1}}, len: 2, cap: 2},
		},
		{
			name:    "returns error when push a value to a full stack",
			fields:  fields{top: &Item[int]{data: 1}, len: 1, cap: 1},
			args:    args{data: 2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedStack[int]{
				top: tt.fields.top,
				len: tt.fields.len,
				cap: tt.fields.cap,
			}
			err := stack.Push(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(stack, tt.want) {
				t.Errorf("Push() stock = %v, want %v", stack, tt.want)
			}
		})
	}
}

func TestFixedStack_Peek(t *testing.T) {
	type fields struct {
		top *Item[int]
		len int
		cap int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name:    "returns error on empty stack",
			fields:  fields{},
			wantErr: true,
		},
		{
			name:   "returns top value from the stack",
			fields: fields{top: &Item[int]{data: 10}, len: 1, cap: 2},
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedStack[int]{
				top: tt.fields.top,
				len: tt.fields.len,
				cap: tt.fields.cap,
			}
			got, err := stack.Peek()
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

func TestFixedStack_IsEmpty(t *testing.T) {
	type fields struct {
		top *Item[int]
		len int
		cap int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "returns true if stack is empty",
			fields: fields{},
			want:   true,
		},
		{
			name:   "returns false if stack is not empty",
			fields: fields{top: &Item[int]{data: 1}, len: 1, cap: 1},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedStack[int]{
				top: tt.fields.top,
				len: tt.fields.len,
				cap: tt.fields.cap,
			}
			if got := stack.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixedStack_IsFull(t *testing.T) {
	type fields struct {
		top *Item[int]
		len int
		cap int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "returns false if stack is empty",
			fields: fields{len: 0, cap: 1},
			want:   false,
		},
		{
			name:   "returns false if stack is not full",
			fields: fields{top: &Item[int]{data: 1}, len: 1, cap: 2},
			want:   false,
		},
		{
			name:   "returns true if stack is full",
			fields: fields{top: &Item[int]{data: 1}, len: 1, cap: 1},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedStack[int]{
				top: tt.fields.top,
				len: tt.fields.len,
				cap: tt.fields.cap,
			}
			if got := stack.IsFull(); got != tt.want {
				t.Errorf("IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}
