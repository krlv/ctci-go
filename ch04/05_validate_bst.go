package ch04

import (
	"cmp"
	"math"
)

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

// validateMinMax checks whether the binary subtree satisfies the binary search
// tree (BST) property given exclusive `min` and `max` bounds.
// For each node, the function ensures `min < node.Data < max` and recursively
// enforces appropriate bounds for left and right subtrees.
//
// Time complexity: O(n) — each node is visited once.
// Space complexity: O(1) excluding recursion stack.
// Auxiliary space: O(h) — recursion stack proportional to tree height h.
func validateMinMax[T cmp.Ordered](root *BinaryTreeNode[T], min, max *T) bool {
	if root == nil {
		return true
	}

	if (min != nil && root.Data <= *min) && (max != nil && root.Data >= *max) {
		return false
	}

	return validateMinMax(root.Left, min, &root.Data) && validateMinMax(root.Right, &root.Data, max)
}

// IsValidBSTMinMax checks whether the binary tree is a valid BST.
// It uses the min/max recursion strategy and starts without bounds.
func IsValidBSTMinMax[T cmp.Ordered](root *BinaryTreeNode[T]) bool {
	return validateMinMax(root, nil, nil)
}

// intTreeCheck represents a single stack frame used by the iterative BST validator.
// The frame carries the node and its lower/upper boundaries such that
// `lower < node.Data < upper`.
type intTreeCheck struct {
	node  *BinaryTreeNode[int]
	lower int
	upper int
}

// IsValidBSTTraversal checks if the binary tree of ints is a valid BST using an
// explicit stack (iterative DFS). Each pushed frame stores the allowed range for the node's value.
// The function returns false as soon as it finds a value outside its permitted bounds.
//
// Time complexity: O(n) — each node is visited and processed once.
// Memory complexity: O(n) worst-case (skewed tree); O(h) for balanced trees,
// where h is the tree height — due to the explicit stack holding frames.
// Space complexity: O(h) average, O(n) worst-case — size of the explicit stack.
func IsValidBSTTraversal(root *BinaryTreeNode[int]) bool {
	if root == nil {
		return true
	}

	stack := []intTreeCheck{
		{
			node:  root,
			lower: math.MinInt,
			upper: math.MaxInt,
		},
	}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.node.Data <= node.lower || node.node.Data >= node.upper {
			return false
		}

		if node.node.Left != nil {
			stack = append(stack, intTreeCheck{
				node:  node.node.Left,
				lower: node.lower,
				upper: node.node.Data,
			})
		}

		if node.node.Right != nil {
			stack = append(stack, intTreeCheck{
				node:  node.node.Right,
				lower: node.node.Data,
				upper: node.upper,
			})
		}
	}

	return true
}
