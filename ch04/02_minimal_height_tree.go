package ch04

import "cmp"

// MinimalHeightBST constructs a binary search tree of minimal height from a sorted slice.
// It selects the middle element as the root and recursively builds left and right subtrees.
// The input slice must be sorted in ascending order.
//
// Time complexity: O(n) — each element is processed exactly once to create a node.
// Space complexity: O(n) for the allocated tree nodes (output).
// Auxiliary space (call stack): O(log n) on average for balanced input, O(n) worst-case for highly unbalanced input.
func MinimalHeightBST[T cmp.Ordered](values []T) *BinaryTreeNode[T] {
	if len(values) == 0 {
		return nil
	}

	mid := len(values) / 2

	n := &BinaryTreeNode[T]{
		Data: values[mid],
	}

	n.Left = MinimalHeightBST(values[:mid])
	n.Right = MinimalHeightBST(values[mid+1:])
	return n
}
