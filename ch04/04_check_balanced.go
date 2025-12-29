package ch04

import (
	"cmp"
	"math"
)

// Height returns the height of the binary tree.
//
// Time complexity: O(n) — each node is visited once.
// Space complexity: O(1) excluding recursion stack.
// Auxiliary space: O(h) — recursion stack proportional to tree height h.
func Height[T cmp.Ordered](root *BinaryTreeNode[T]) int {
	if root == nil {
		return -1
	}

	return max(Height(root.Left), Height(root.Right)) + 1
}

// IsBalanced returns true if the binary tree is height-balanced.
// A tree is balanced if, for every node, the heights of the left and right
// subtrees differ by at most 1.
//
// Time complexity: O(n^2) in the worst case because Height is re-called for many nodes.
// Space complexity: O(1) excluding recursion stack.
// Auxiliary space: O(h) — recursion stack proportional to tree height h.
func IsBalanced[T cmp.Ordered](root *BinaryTreeNode[T]) bool {
	if root == nil {
		return true
	}

	diff := Height(root.Left) - Height(root.Right)
	if math.Abs(float64(diff)) > 1 {
		return false
	} else {
		return IsBalanced(root.Left) && IsBalanced(root.Right)
	}
}
