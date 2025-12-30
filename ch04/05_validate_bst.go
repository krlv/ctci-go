package ch04

import "cmp"

// IsValidBSTInOrder checks whether the binary tree is a valid binary search tree.
// It performs an in-order traversal and verifies that the
// visited node values are strictly increasing (each node's value must be
// greater than the previously visited value).
//
// Time complexity: O(n) — each node is visited once.
// Space complexity: O(1) excluding recursion stack.
// Auxiliary space: O(h) — recursion stack proportional to tree height h.
func IsValidBSTInOrder[T cmp.Ordered](root *BinaryTreeNode[T]) bool {
	var prev *BinaryTreeNode[T]

	var inorder func(root *BinaryTreeNode[T]) bool
	inorder = func(root *BinaryTreeNode[T]) bool {
		if root == nil {
			return true
		}

		if !inorder(root.Left) {
			return false
		}

		if prev != nil && root.Data <= prev.Data {
			return false
		}
		prev = root

		return inorder(root.Right)
	}

	return inorder(root)
}
