package ch02

// SumLists returns sum of n1 and n2 numbers as a linked list
func SumLists(n1 *Node, n2 *Node) *Node {
	head := New(0) // sum list; the first head node will remain empty
	tail := head

	value := 0
	for n1 != nil && n2 != nil {
		if n1 != nil {
			value += n1.data
			n1 = n1.next
		}

		if n2 != nil {
			value += n2.data
			n2 = n2.next
		}

		// add a new node to the list
		tail.next = New(value % 10)
		tail = tail.next

		// carry the (1 || 0) * 10 value for the next iteration
		value = value / 10
	}

	// the head node is empty, skip it
	return head.next
}

// SumListsReverse returns sum of n1 and n2 numbers as a linked list
// Numbers n1, n2 and resulted sum are stored in forward order
func SumListsReverse(n1 *Node, n2 *Node) *Node {
	var head, tail, prev *Node

	head = New(0) // sum list
	tail = head

	for {
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
			if prev == nil {
				tail.next = New(0)
				prev = tail
				tail = tail.next
			}

			prev.data++
			tail.data = value % 10
		} else {
			tail.data = value
		}

		tail.next = New(0)

		if n1 == nil && n2 == nil {
			tail.next = nil
			break
		}

		prev = tail
		tail = tail.next
	}

	return head
}
