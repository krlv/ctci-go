package ch03

import (
	"testing"
)

func TestStackMin_Pop(t *testing.T) {
	type fields struct {
		top *MinItem
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name:    "returns error on empty stack",
			fields:  fields{top: nil},
			wantErr: true,
		},
		{
			name:   "pops top value from stack",
			fields: fields{top: &MinItem{data: 1, min: 1, next: nil}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &StackMin{
				top: tt.fields.top,
			}
			got, err := stack.Pop()
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

func TestStackMin_Push(t *testing.T) {
	type fields struct {
		top *MinItem
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "updates top.min on new minimal value push",
			fields: fields{top: &MinItem{data: 1, min: 1, next: nil}},
			args:   args{data: -1},
			want:   -1,
		},
		{
			name:   "doesn't update top.min on push with non-minimal value",
			fields: fields{top: &MinItem{data: 1, min: 1, next: nil}},
			args:   args{data: 2},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &StackMin{
				top: tt.fields.top,
			}

			stack.Push(tt.args.data)

			got := stack.top.min
			if got != tt.want {
				t.Errorf("stack.top.min got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackMin_Peek(t *testing.T) {
	type fields struct {
		top *MinItem
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name:    "returns error on empty stack",
			fields:  fields{top: nil},
			wantErr: true,
		},
		{
			name:   "returns top value from stack",
			fields: fields{top: &MinItem{data: 1, min: 1, next: nil}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &StackMin{
				top: tt.fields.top,
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

func TestStackMin_IsEmpty(t *testing.T) {
	type fields struct {
		top *MinItem
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "returns true on empty stack",
			fields: fields{top: nil},
			want:   true,
		},
		{
			name:   "returns false on non-empty stack",
			fields: fields{top: &MinItem{data: 1, min: 1, next: nil}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &StackMin{
				top: tt.fields.top,
			}
			if got := stack.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackMin_Min(t *testing.T) {
	type fields struct {
		top *MinItem
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name:    "returns error on empty stack",
			fields:  fields{top: nil},
			wantErr: true,
		},
		{
			name:   "returns stack minimal value",
			fields: fields{top: &MinItem{data: 1, min: 1, next: nil}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &StackMin{
				top: tt.fields.top,
			}
			got, err := stack.Min()
			if (err != nil) != tt.wantErr {
				t.Errorf("Min() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Min() got = %v, want %v", got, tt.want)
			}
		})
	}
}
