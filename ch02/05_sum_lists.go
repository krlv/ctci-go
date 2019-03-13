package ch02

// SumLists returns sum of n1 and n2 numbers as a linked list
func SumLists(n1 *Node, n2 *Node) *Node {
	head := New(0) // sum list
	tail := head

	for n1 != nil || n2 != nil {
		value := tail.data

		if n1 != nil {
			value += n1.data
			n1 = n1.next
		}

		if n2 != nil {
			value += n2.data
			n2 = n2.next
		}

		if value > 9 {
			tail.data = value % 10
			tail.next = New(1)
		} else {
			tail.data = value
		}

		tail = tail.next
	}

	return head
}
