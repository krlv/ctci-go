package ch04

// DepthList returns a slice of linked lists where each linked list contains
// the nodes at a particular depth (level) of the binary tree.
//
// Example: the list at index 0 contains the root, index 1 contains all nodes at depth 1, etc.
//
// The function performs a DFS traversal and builds/extends a list for each depth.
// Time complexity: O(n) — each node is visited exactly once.
// Space complexity: O(n) — the returned slice and contained lists store every node.
// Auxiliary space: O(h) — recursion stack proportional to tree height h.
func DepthList(root *BinaryTreeNode[int]) []*Node[*BinaryTreeNode[int]] {
	lists := make([]*Node[*BinaryTreeNode[int]], 0)
	depthList(root, &lists, 0)
	return lists
}

// depthList is a recursive helper that inserts the given root node into the
// linked list of the current depth. If a list for this depth does not yet exist,
// it appends a new list head to *lists.
// Traversal is pre-order (process node, then left, then right).
func depthList(root *BinaryTreeNode[int], lists *[]*Node[*BinaryTreeNode[int]], depth int) {
	if root == nil {
		return
	}

	if len(*lists) == depth {
		list := &Node[*BinaryTreeNode[int]]{data: root}
		*lists = append(*lists, list)
	} else {
		list := (*lists)[depth]
		list.Append(root)
	}

	depthList(root.Left, lists, depth+1)
	depthList(root.Right, lists, depth+1)
}
