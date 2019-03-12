package ch02

// DeleteMiddleNode deletes middle node m from the list
func DeleteMiddleNode(m *Node) {
	for ; m.next != nil; m = m.next {
		m.data = m.next.data
		m.next = m.next.next
	}
}
