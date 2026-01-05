package ch04

import (
	"cmp"
	"reflect"
	"testing"
)

func TestInOrderSuccessor(t *testing.T) {
	single := &BinaryTreeParent[int]{Data: 1}

	node06 := &BinaryTreeParent[int]{Data: 6}
	node05 := &BinaryTreeParent[int]{Data: 5}
	node09 := &BinaryTreeParent[int]{Data: 9}
	node03 := &BinaryTreeParent[int]{Data: 3}
	node08 := &BinaryTreeParent[int]{Data: 8}
	node10 := &BinaryTreeParent[int]{Data: 10}
	node02 := &BinaryTreeParent[int]{Data: 2}
	node04 := &BinaryTreeParent[int]{Data: 4}
	node12 := &BinaryTreeParent[int]{Data: 12}
	node02.Parent, node04.Parent = node03, node03
	node12.Parent = node10
	node03.Left, node03.Right = node02, node04
	node10.Right = node12
	node03.Parent = node05
	node08.Parent, node10.Parent = node09, node09
	node05.Left = node03
	node09.Left, node09.Right = node08, node10
	node05.Parent, node09.Parent = node06, node06
	node06.Left, node06.Right = node05, node09

	type testCase[T cmp.Ordered] struct {
		name string
		node *BinaryTreeParent[T]
		want *BinaryTreeParent[T]
	}
	tests := []testCase[int]{
		{
			name: "empty tree",
			node: nil,
			want: nil,
		},
		{
			name: "single node",
			node: single,
			want: nil,
		},
		{
			name: "successor for root node",
			node: node06,
			want: node08,
		},
		{
			name: "successor for node in right subtree",
			node: node10,
			want: node12,
		},
		{
			name: "successor for node in left subtree",
			node: node03,
			want: node04,
		},
		{
			name: "successor for most right node",
			node: node12,
			want: nil,
		},
		{
			name: "successor for most left node",
			node: node02,
			want: node03,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InOrderSuccessor(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InOrderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
