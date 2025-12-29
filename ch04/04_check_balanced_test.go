package ch04

import (
	"cmp"
	"testing"
)

func TestHeight(t *testing.T) {
	balanced := &BinaryTreeNode[int]{
		Data: 11,
		Left: &BinaryTreeNode[int]{
			Data:  21,
			Left:  &BinaryTreeNode[int]{Data: 31},
			Right: &BinaryTreeNode[int]{Data: 32},
		},
		Right: &BinaryTreeNode[int]{
			Data:  22,
			Left:  &BinaryTreeNode[int]{Data: 33},
			Right: &BinaryTreeNode[int]{Data: 34},
		},
	}
	stillBalanced := &BinaryTreeNode[int]{
		Data: 11,
		Left: &BinaryTreeNode[int]{
			Data:  21,
			Left:  &BinaryTreeNode[int]{Data: 31},
			Right: &BinaryTreeNode[int]{Data: 32},
		},
		Right: &BinaryTreeNode[int]{
			Data: 22,
		},
	}
	unbalanced := &BinaryTreeNode[int]{
		Data: 11,
		Left: &BinaryTreeNode[int]{
			Data: 21,
			Left: &BinaryTreeNode[int]{
				Data: 31,
				Left: &BinaryTreeNode[int]{
					Data: 41,
					Left: &BinaryTreeNode[int]{
						Data: 51,
					},
				},
			},
			Right: &BinaryTreeNode[int]{Data: 32},
		},
		Right: &BinaryTreeNode[int]{Data: 22},
	}

	type testCase[T cmp.Ordered] struct {
		name string
		root *BinaryTreeNode[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "empty tree",
			root: nil,
			want: -1,
		},
		{
			name: "single node",
			root: &BinaryTreeNode[int]{Data: 1},
			want: 0,
		},
		{
			name: "3 level balanced tree",
			root: balanced,
			want: 2,
		},
		{
			name: "3 level balanced tree with 1 level diff",
			root: stillBalanced,
			want: 2,
		},
		{
			name: "5 level unbalanced tree",
			root: unbalanced,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Height(tt.root); got != tt.want {
				t.Errorf("Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBalanced(t *testing.T) {
	balanced := &BinaryTreeNode[int]{
		Data: 11,
		Left: &BinaryTreeNode[int]{
			Data:  21,
			Left:  &BinaryTreeNode[int]{Data: 31},
			Right: &BinaryTreeNode[int]{Data: 32},
		},
		Right: &BinaryTreeNode[int]{
			Data:  22,
			Left:  &BinaryTreeNode[int]{Data: 33},
			Right: &BinaryTreeNode[int]{Data: 34},
		},
	}
	stillBalanced := &BinaryTreeNode[int]{
		Data: 11,
		Left: &BinaryTreeNode[int]{
			Data:  21,
			Left:  &BinaryTreeNode[int]{Data: 31},
			Right: &BinaryTreeNode[int]{Data: 32},
		},
		Right: &BinaryTreeNode[int]{
			Data: 22,
		},
	}
	unbalanced := &BinaryTreeNode[int]{
		Data: 11,
		Left: &BinaryTreeNode[int]{
			Data: 21,
			Left: &BinaryTreeNode[int]{
				Data: 31,
				Left: &BinaryTreeNode[int]{
					Data: 41,
					Left: &BinaryTreeNode[int]{
						Data: 51,
					},
				},
			},
			Right: &BinaryTreeNode[int]{Data: 32},
		},
		Right: &BinaryTreeNode[int]{Data: 22},
	}

	type testCase[T cmp.Ordered] struct {
		name string
		root *BinaryTreeNode[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "empty tree",
			root: nil,
			want: true,
		},
		{
			name: "single node",
			root: &BinaryTreeNode[int]{Data: 1},
			want: true,
		},
		{
			name: "3 level balanced tree",
			root: balanced,
			want: true,
		},
		{
			name: "3 level balanced tree with 1 level diff",
			root: stillBalanced,
			want: true,
		},
		{
			name: "5 level unbalanced tree",
			root: unbalanced,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBalanced(tt.root); got != tt.want {
				t.Errorf("IsBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
