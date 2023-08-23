package ch03

import (
	"reflect"
	"testing"
)

func TestStackSet_Pop(t *testing.T) {
	stack1 := NewFixedStack[int](2)
	stack1.Push(1)
	stack1.Push(2)

	stack2 := NewFixedStack[int](2)
	stack2.Push(3)

	type fields struct {
		stacks []*FixedStack[int]
	}
	type want struct {
		pop int
		len int
	}
	tests := []struct {
		name    string
		fields  fields
		want    want
		wantErr bool
	}{
		{
			name:    "returns error if the set is empty",
			fields:  fields{stacks: []*FixedStack[int]{}},
			wantErr: true,
		},
		{
			name:    "returns error if the top stack is empty",
			fields:  fields{stacks: append([]*FixedStack[int]{}, NewFixedStack[int](2))},
			wantErr: true,
		},
		{
			name:   "returns the top value and removes empty stack from the set",
			fields: fields{stacks: append([]*FixedStack[int]{}, stack1, stack2)},
			want:   want{pop: 3, len: 1},
		},
		{
			name:   "returns the top value from the top stack in the set",
			fields: fields{stacks: append([]*FixedStack[int]{}, stack1)},
			want:   want{pop: 2, len: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := &StackSet{
				stacks: tt.fields.stacks,
			}
			got, err := set.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if got != tt.want.pop {
				t.Errorf("Pop() got = %v, want %v", got, tt.want.pop)
			}
			got = len(set.stacks)
			if got != tt.want.len {
				t.Errorf("Pop() len(set.stacks) = %v, want %v", got, tt.want.len)
			}
		})
	}
}

func TestStackSet_Push(t *testing.T) {
	stack1 := NewFixedStack[int](2)
	stack1.Push(1)

	stack2 := NewFixedStack[int](2)
	stack2.Push(1)
	stack2.Push(2)

	type fields struct {
		stacks []*FixedStack[int]
	}
	type args struct {
		data int
	}
	type want struct {
		len int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		wantErr bool
	}{
		{
			name:   "pushes a value to an empty SetStack",
			fields: fields{stacks: []*FixedStack[int]{}},
			args:   args{data: 1},
			want:   want{len: 1},
		},
		{
			name:   "pushes a value to non-full top stack",
			fields: fields{stacks: append([]*FixedStack[int]{}, stack1)},
			args:   args{data: 1},
			want:   want{len: 1},
		},
		{
			name:   "appends new stack to the set on Push() when top stack is full",
			fields: fields{stacks: append([]*FixedStack[int]{}, stack2)},
			args:   args{data: 1},
			want:   want{len: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := &StackSet{
				stacks: tt.fields.stacks,
			}
			if err := set.Push(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
			got := len(set.stacks)
			if got != tt.want.len {
				t.Errorf("Pop() len(set.stacks) = %v, want %v", got, tt.want.len)
			}
		})
	}
}

func TestStackSet_topStack(t *testing.T) {
	stack := NewFixedStack[int](2)
	stack.Push(1)
	stack.Push(2)

	type fields struct {
		stacks []*FixedStack[int]
	}
	tests := []struct {
		name    string
		fields  fields
		want    *FixedStack[int]
		wantErr bool
	}{
		{
			name:    "returns error on empty set",
			fields:  fields{},
			wantErr: true,
		},
		{
			name:   "returns the top stack from the set",
			fields: fields{stacks: append([]*FixedStack[int]{}, stack)},
			want:   stack,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := &StackSet{
				stacks: tt.fields.stacks,
			}
			got, err := set.topStack()
			if (err != nil) != tt.wantErr {
				t.Errorf("topStack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topStack() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackSet_appendStack(t *testing.T) {
	type fields struct {
		stacks []*FixedStack[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   *FixedStack[int]
	}{
		{
			name:   "appends new stack to the set",
			fields: fields{},
			want:   NewFixedStack[int](3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := &StackSet{
				stacks: tt.fields.stacks,
			}
			if got := set.appendStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackSet_removeStack(t *testing.T) {
	type fields struct {
		stacks []*FixedStack[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   []*FixedStack[int]
	}{
		{
			name:   "do nothing if no stacks in the set",
			fields: fields{stacks: []*FixedStack[int]{}},
			want:   []*FixedStack[int]{},
		},
		{
			name:   "removes top stack from the set",
			fields: fields{stacks: append([]*FixedStack[int]{}, NewFixedStack[int](1), NewFixedStack[int](1))},
			want:   append([]*FixedStack[int]{}, NewFixedStack[int](1)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := &StackSet{
				stacks: tt.fields.stacks,
			}
			set.removeStack()
			if !reflect.DeepEqual(set.stacks, tt.want) {
				t.Errorf("appendStack() set.stacks = %v, want %v", set.stacks, tt.want)
			}
		})
	}
}
