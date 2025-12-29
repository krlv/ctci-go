package ch04

import (
	"cmp"
	"errors"
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

// BalancedHeight returns the height of the binary tree
// or an error if the tree is not height-balanced.
//
// The function performs a post-order traversal: it obtains the heights of the
// left and right subtrees, returns an error up the call stack if either is
// unbalanced, otherwise computes the current node's height. An empty
// tree has height -1.
//
// Time complexity: O(n) — each node is visited once.
// Space complexity: O(1) excluding recursion stack.
// Auxiliary space: O(h) — recursion stack proportional to tree height h.
func BalancedHeight[T cmp.Ordered](root *BinaryTreeNode[T]) (int, error) {
	if root == nil {
		return -1, nil
	}

	leftHeight, err := BalancedHeight(root.Left)
	if err != nil {
		return 0, err
	}

	rightHeigh, err := BalancedHeight(root.Right)
	if err != nil {
		return 0, err
	}

	diff := int(math.Abs(float64(leftHeight - rightHeigh)))
	if diff > 1 {
		return 0, errors.New("unbalanced height")
	}

	return max(leftHeight, rightHeigh) + 1, nil
}

// IsBalancedSingleTraversal returns true if the binary tree is height-balanced.
// This implementation performs a single traversal: the BalancedHeight returns
// both the subtree height and whether the subtree is balanced.
func IsBalancedSingleTraversal[T cmp.Ordered](root *BinaryTreeNode[T]) bool {
	_, err := BalancedHeight(root)
	return err == nil
}
