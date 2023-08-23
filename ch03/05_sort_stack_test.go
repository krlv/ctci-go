package ch03

import "testing"

func TestSort(t *testing.T) {
	type args struct {
		stack  *Stack[int]
		values []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sorts random stack",
			args: args{stack: new(Stack[int]), values: []int{1, 2, 3, 4, 10}},
			want: []int{1, 2, 3, 4, 10},
		},
		{
			name: "sorts a stack that already sorted",
			args: args{stack: new(Stack[int]), values: []int{5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := tt.args.stack
			for _, v := range tt.args.values {
				stack.Push(v)
			}
			Sort(stack)
			for _, want := range tt.want {
				got, _ := stack.Pop()
				if got != want {
					t.Errorf("Sorted stack Pop() got = %v, want %v", got, want)
				}
			}
		})
	}
}
