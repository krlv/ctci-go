package ch04

import (
	"reflect"
	"testing"
)

func TestNode_Append(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		n    *Node[T]
		data T
		want *Node[T]
	}
	tests := []testCase[int]{
		{
			name: "empty list",
			n:    &Node[int]{},
			data: 1,
			want: &Node[int]{next: &Node[int]{data: 1}},
		},
		{
			name: "non-empty list",
			n:    &Node[int]{data: 1, next: &Node[int]{data: 2, next: &Node[int]{data: 3}}},
			data: 4,
			want: &Node[int]{data: 1, next: &Node[int]{data: 2, next: &Node[int]{data: 3, next: &Node[int]{data: 4}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.Append(tt.data)
			if !reflect.DeepEqual(tt.n, tt.want) {
				t.Errorf("Append() got = %v, want %v", tt.n, tt.want)
			}
		})
	}
}
