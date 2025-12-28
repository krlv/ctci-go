package ch04

// Append adds a new node with data at the tail of the list
func (n *Node[T]) Append(data T) {
	end := &Node[T]{data: data}

	for n.next != nil {
		n = n.next
	}

	n.next = end
}
