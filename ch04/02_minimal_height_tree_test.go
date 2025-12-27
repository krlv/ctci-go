package ch04

import (
	"cmp"
	"reflect"
	"testing"
)

func TestMinimalHeightBST(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name   string
		values []T
		want   *BinaryTreeNode[T]
	}
	tests := []testCase[int]{
		{
			name:   "empty values slice",
			values: []int{},
			want:   nil,
		},
		{
			name:   "single node tree",
			values: []int{1},
			want:   &BinaryTreeNode[int]{Data: 1},
		},
		{
			name:   "5 nodes tree",
			values: []int{1, 2, 3, 4, 5},
			want: &BinaryTreeNode[int]{
				Data: 3,
				Left: &BinaryTreeNode[int]{
					Data: 2,
					Left: &BinaryTreeNode[int]{Data: 1},
					// Right: nil
				},
				Right: &BinaryTreeNode[int]{
					Data: 5,
					Left: &BinaryTreeNode[int]{Data: 4},
					// Right: nil
				},
			},
		},
		{
			name:   "7 nodes tree",
			values: []int{1, 2, 3, 4, 5, 6, 7},
			want: &BinaryTreeNode[int]{
				Data: 4,
				Left: &BinaryTreeNode[int]{
					Data:  2,
					Left:  &BinaryTreeNode[int]{Data: 1},
					Right: &BinaryTreeNode[int]{Data: 3},
				},
				Right: &BinaryTreeNode[int]{
					Data:  6,
					Left:  &BinaryTreeNode[int]{Data: 5},
					Right: &BinaryTreeNode[int]{Data: 7},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimalHeightBST(tt.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinimalHeightBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
