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

// DepthListBFS builds level lists using a queue (BFS) and per-level tail pointers.
// Time complexity: O(n) — each node processed once.
// Space complexity: O(n) — returned slice and lists store every node.
// Auxiliary space: O(w) — queue holds up to the maximum width w of the tree.
func DepthListBFS(root *BinaryTreeNode[int]) []*Node[*BinaryTreeNode[int]] {
	if root == nil {
		return nil
	}

	type treeNodeQueue struct {
		node  *BinaryTreeNode[int]
		depth int
	}

	queue := []treeNodeQueue{{node: root, depth: 0}}
	lists := make([]*Node[*BinaryTreeNode[int]], 0)

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item.depth == len(lists) {
			head := &Node[*BinaryTreeNode[int]]{data: item.node}
			lists = append(lists, head)
		} else {
			lists[item.depth].Append(item.node)
		}

		if item.node.Left != nil {
			queue = append(queue, treeNodeQueue{node: item.node.Left, depth: item.depth + 1})
		}
		if item.node.Right != nil {
			queue = append(queue, treeNodeQueue{node: item.node.Right, depth: item.depth + 1})
		}
	}

	return lists
}
