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
