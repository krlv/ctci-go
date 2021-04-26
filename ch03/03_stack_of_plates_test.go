package ch03

import (
	"reflect"
	"testing"
)

func TestStackSet_Pop(t *testing.T) {
	stack1 := NewFixedStack(2)
	stack1.Push(1)
	stack1.Push(2)

	stack2 := NewFixedStack(2)
	stack2.Push(3)

	type fields struct {
		stacks []*FixedStack
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
			fields:  fields{stacks: []*FixedStack{}},
			wantErr: true,
		},
		{
			name:    "returns error if the top stack is empty",
			fields:  fields{stacks: append([]*FixedStack{}, NewFixedStack(2))},
			wantErr: true,
		},
		{
			name:   "returns the top value and removes empty stack from the set",
			fields: fields{stacks: append([]*FixedStack{}, stack1, stack2)},
			want:   want{pop: 3, len: 1},
		},
		{
			name:   "returns the top value from the top stack in the set",
			fields: fields{stacks: append([]*FixedStack{}, stack1)},
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
	stack1 := NewFixedStack(2)
	stack1.Push(1)

	stack2 := NewFixedStack(2)
	stack2.Push(1)
	stack2.Push(2)

	type fields struct {
		stacks []*FixedStack
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
			fields: fields{stacks: []*FixedStack{}},
			args:   args{data: 1},
			want:   want{len: 1},
		},
		{
			name:   "pushes a value to non-full top stack",
			fields: fields{stacks: append([]*FixedStack{}, stack1)},
			args:   args{data: 1},
			want:   want{len: 1},
		},
		{
			name:   "appends new stack to the set on Push() when top stack is full",
			fields: fields{stacks: append([]*FixedStack{}, stack2)},
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
	stack := NewFixedStack(2)
	stack.Push(1)
	stack.Push(2)

	type fields struct {
		stacks []*FixedStack
	}
	tests := []struct {
		name    string
		fields  fields
		want    *FixedStack
		wantErr bool
	}{
		{
			name: "returns error on empty set",
			fields: fields{},
			wantErr: true,
		},
		{
			name: "returns the top stack from the set",
			fields: fields{stacks: append([]*FixedStack{}, stack)},
			want: stack,
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
		stacks []*FixedStack
	}
	tests := []struct {
		name   string
		fields fields
		want   *FixedStack
	}{
		{
			name: "appends new stack to the set",
			fields: fields{},
			want: NewFixedStack(3),
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
		stacks []*FixedStack
	}
	tests := []struct {
		name   string
		fields fields
		want   []*FixedStack
	}{
		{
			name: "do nothing if no stacks in the set",
			fields: fields{stacks: []*FixedStack{}},
			want: []*FixedStack{},
		},
		{
			name: "removes top stack from the set",
			fields: fields{stacks: append([]*FixedStack{}, NewFixedStack(1), NewFixedStack(1))},
			want: append([]*FixedStack{}, NewFixedStack(1)),
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

func TestNewFixedStack(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want *FixedStack
	}{
		{
			name: "creates an empty stack with predefined capacity",
			args: args{cap: 10},
			want: &FixedStack{cap: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFixedStack(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFixedStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixedStack_Pop(t *testing.T) {
	type fields struct {
		top *Item
		len int
		cap int
	}
	type want struct {
		top   int
		stack *FixedStack
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
			fields: fields{top: &Item{data: 10}, len: 1, cap: 1},
			want:   want{top: 10, stack: &FixedStack{len: 0, cap: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedStack{
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
		top *Item
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
		want    *FixedStack
		wantErr bool
	}{
		{
			name:   "pushes new value to an empty stack",
			fields: fields{top: nil, len: 0, cap: 2},
			args:   args{data: 1},
			want:   &FixedStack{top: &Item{data: 1}, len: 1, cap: 2},
		},
		{
			name:   "pushes new value to a non empty stack",
			fields: fields{top: &Item{data: 1}, len: 1, cap: 2},
			args:   args{data: 2},
			want:   &FixedStack{top: &Item{data: 2, next: &Item{data: 1}}, len: 2, cap: 2},
		},
		{
			name:    "returns error when push a value to a full stack",
			fields:  fields{top: &Item{data: 1}, len: 1, cap: 1},
			args:    args{data: 2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedStack{
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
		top *Item
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
			fields: fields{top: &Item{data: 10}, len: 1, cap: 2},
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedStack{
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
		top *Item
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
			fields: fields{top: &Item{data: 1}, len: 1, cap: 1},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedStack{
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
		top *Item
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
			fields: fields{top: &Item{data: 1}, len: 1, cap: 2},
			want:   false,
		},
		{
			name:   "returns true if stack is full",
			fields: fields{top: &Item{data: 1}, len: 1, cap: 1},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &FixedStack{
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
