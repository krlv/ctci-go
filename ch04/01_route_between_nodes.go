package ch04

// HasRouteBetweenNodes checks if there is a route between two nodes in a graph.
// It uses a breadth-first search (BFS) to find the route.
// Time complexity: O(V + E), where V is the number of vertices and E is the number of edges.
// Space complexity: O(V), where V is the number of vertices. There can be atmost V elements in the queue.
func HasRouteBetweenNodes[T comparable](g *Graph[T], start, end *GraphNode[T]) bool {
	// track visited nodes
	visited := make(map[*GraphNode[T]]bool)

	// BFS: use a queue to track nodes to visit
	queue := SQueue[*GraphNode[T]]{}
	queue.Enqueue(start)

	for !queue.IsEmpty() {
		node, _ := queue.Dequeue()
		if node == end {
			return true
		}

		for _, v := range node.adjacent {
			if !visited[v] {
				queue.Enqueue(v)
			}
		}

		visited[node] = true
	}

	return false
}
