package ch04

import "cmp"

// BinaryTreeNode represents a node in a binary tree
type BinaryTreeNode[T cmp.Ordered] struct {
	Data  T
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}

// WalkFunc is a function that can be used to traverse a tree
type WalkFunc[T cmp.Ordered] func(T) error

// InOrderTraversal traverses a binary tree in order
// When performed on a binary search tree, the nodes are visited in ascending order, i.e. "in-order".
func InOrderTraversal[T cmp.Ordered](n *BinaryTreeNode[T], visitFunc WalkFunc[T]) error {
	if n == nil {
		return nil
	}

	if err := InOrderTraversal(n.Left, visitFunc); err != nil {
		return err
	}

	if err := visitFunc(n.Data); err != nil {
		return err
	}

	if err := InOrderTraversal(n.Right, visitFunc); err != nil {
		return err
	}

	return nil
}

// PreOrderTraversal traverses a binary tree in pre-order
// In a pre-order traversal, the root is the first node visited.
func PreOrderTraversal[T cmp.Ordered](n *BinaryTreeNode[T], visitFunc WalkFunc[T]) error {
	if n == nil {
		return nil
	}

	if err := visitFunc(n.Data); err != nil {
		return err
	}

	if err := PreOrderTraversal(n.Left, visitFunc); err != nil {
		return err
	}

	if err := PreOrderTraversal(n.Right, visitFunc); err != nil {
		return err
	}

	return nil
}

// PostOrderTraversal traverses a binary tree in post-order
// In a post-order traversal, the root is the last node visited.
func PostOrderTraversal[T cmp.Ordered](n *BinaryTreeNode[T], visitFunc WalkFunc[T]) error {
	if n == nil {
		return nil
	}

	if err := PostOrderTraversal(n.Left, visitFunc); err != nil {
		return err
	}

	if err := PostOrderTraversal(n.Right, visitFunc); err != nil {
		return err
	}

	return visitFunc(n.Data)
}
