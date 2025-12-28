package ch04

import (
	"reflect"
	"testing"
)

func TestDepthList(t *testing.T) {
	node31 := &BinaryTreeNode[int]{Data: 31}
	node32 := &BinaryTreeNode[int]{Data: 32}
	node33 := &BinaryTreeNode[int]{Data: 33}
	node34 := &BinaryTreeNode[int]{Data: 34}
	node21 := &BinaryTreeNode[int]{
		Data:  21,
		Left:  node31,
		Right: node32,
	}
	node22 := &BinaryTreeNode[int]{
		Data:  22,
		Left:  node33,
		Right: node34,
	}
	node11 := &BinaryTreeNode[int]{
		Data:  11,
		Left:  node21,
		Right: node22,
	}

	tests := []struct {
		name string
		root *BinaryTreeNode[int]
		want []*Node[*BinaryTreeNode[int]]
	}{
		{
			name: "empty tree",
			root: nil,
			want: []*Node[*BinaryTreeNode[int]]{},
		},
		{
			name: "3 level tree",
			root: node11,
			want: []*Node[*BinaryTreeNode[int]]{
				{data: node11},
				{data: node21, next: &Node[*BinaryTreeNode[int]]{data: node22}},
				{data: node31, next: &Node[*BinaryTreeNode[int]]{data: node32, next: &Node[*BinaryTreeNode[int]]{data: node33, next: &Node[*BinaryTreeNode[int]]{data: node34}}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DepthList(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DepthList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDepthListBFS(t *testing.T) {
	node31 := &BinaryTreeNode[int]{Data: 31}
	node32 := &BinaryTreeNode[int]{Data: 32}
	node33 := &BinaryTreeNode[int]{Data: 33}
	node34 := &BinaryTreeNode[int]{Data: 34}
	node21 := &BinaryTreeNode[int]{
		Data:  21,
		Left:  node31,
		Right: node32,
	}
	node22 := &BinaryTreeNode[int]{
		Data:  22,
		Left:  node33,
		Right: node34,
	}
	node11 := &BinaryTreeNode[int]{
		Data:  11,
		Left:  node21,
		Right: node22,
	}

	tests := []struct {
		name string
		root *BinaryTreeNode[int]
		want []*Node[*BinaryTreeNode[int]]
	}{
		{
			name: "empty tree",
			root: nil,
			want: nil,
		},
		{
			name: "3 level tree",
			root: node11,
			want: []*Node[*BinaryTreeNode[int]]{
				{data: node11},
				{data: node21, next: &Node[*BinaryTreeNode[int]]{data: node22}},
				{data: node31, next: &Node[*BinaryTreeNode[int]]{data: node32, next: &Node[*BinaryTreeNode[int]]{data: node33, next: &Node[*BinaryTreeNode[int]]{data: node34}}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DepthListBFS(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DepthListBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
