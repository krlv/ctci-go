package ch04

import (
	"cmp"
	"testing"
)

func TestIsValidBSTInOrder(t *testing.T) {
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
			name: "2 nodes invalid BST",
			root: &BinaryTreeNode[int]{
				Data:  20,
				Right: &BinaryTreeNode[int]{Data: 20},
			},
			want: false,
		},
		{
			name: "3 nodes valid",
			root: &BinaryTreeNode[int]{
				Data:  2,
				Left:  &BinaryTreeNode[int]{Data: 1},
				Right: &BinaryTreeNode[int]{Data: 3},
			},
			want: true,
		},
		{
			name: "5 nodes invalid",
			root: &BinaryTreeNode[int]{
				Data: 5,
				Left: &BinaryTreeNode[int]{Data: 1},
				Right: &BinaryTreeNode[int]{
					Data:  4,
					Left:  &BinaryTreeNode[int]{Data: 3},
					Right: &BinaryTreeNode[int]{Data: 6},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBSTInOrder(tt.root); got != tt.want {
				t.Errorf("IsValidBSTInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidBSTMinMax(t *testing.T) {
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
			name: "2 nodes invalid BST",
			root: &BinaryTreeNode[int]{
				Data:  20,
				Right: &BinaryTreeNode[int]{Data: 20},
			},
			want: false,
		},
		{
			name: "3 nodes valid",
			root: &BinaryTreeNode[int]{
				Data:  2,
				Left:  &BinaryTreeNode[int]{Data: 1},
				Right: &BinaryTreeNode[int]{Data: 3},
			},
			want: true,
		},
		{
			name: "5 nodes invalid",
			root: &BinaryTreeNode[int]{
				Data: 5,
				Left: &BinaryTreeNode[int]{Data: 1},
				Right: &BinaryTreeNode[int]{
					Data:  4,
					Left:  &BinaryTreeNode[int]{Data: 3},
					Right: &BinaryTreeNode[int]{Data: 6},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBSTMinMax(tt.root); got != tt.want {
				t.Errorf("IsValidBSTMinMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidBSTTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *BinaryTreeNode[int]
		want bool
	}{
		{
			name: "empty tree",
			root: nil,
			want: true,
		},
		{
			name: "2 nodes invalid BST",
			root: &BinaryTreeNode[int]{
				Data:  20,
				Right: &BinaryTreeNode[int]{Data: 20},
			},
			want: false,
		},
		{
			name: "3 nodes valid",
			root: &BinaryTreeNode[int]{
				Data:  2,
				Left:  &BinaryTreeNode[int]{Data: 1},
				Right: &BinaryTreeNode[int]{Data: 3},
			},
			want: true,
		},
		{
			name: "5 nodes invalid",
			root: &BinaryTreeNode[int]{
				Data: 5,
				Left: &BinaryTreeNode[int]{Data: 1},
				Right: &BinaryTreeNode[int]{
					Data:  4,
					Left:  &BinaryTreeNode[int]{Data: 3},
					Right: &BinaryTreeNode[int]{Data: 6},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBSTTraversal(tt.root); got != tt.want {
				t.Errorf("IsValidBSTTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
