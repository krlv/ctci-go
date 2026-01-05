package ch04

import "cmp"

type BinaryTreeParent[T cmp.Ordered] struct {
	Data   T
	Left   *BinaryTreeParent[T]
	Right  *BinaryTreeParent[T]
	Parent *BinaryTreeParent[T]
}

func (n *BinaryTreeParent[T]) LeftMostChild() *BinaryTreeParent[T] {
	if n == nil {
		return nil
	}

	for n.Left != nil {
		n = n.Left
	}

	return n
}

func InOrderSuccessor[T cmp.Ordered](node *BinaryTreeParent[T]) *BinaryTreeParent[T] {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		return node.Right.LeftMostChild()
	}

	curr, prev := node, node.Parent
	for prev != nil && prev.Left != curr {
		curr, prev = prev, prev.Parent
	}

	return prev
}
