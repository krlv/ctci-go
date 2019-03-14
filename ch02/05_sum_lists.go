package ch02

// SumLists returns sum of n1 and n2 numbers as a linked list
func SumLists(n1 *Node, n2 *Node) *Node {
	head := New(0) // result sum list
	tail := head

	for {
		// tail.data is a carrier for (1 || 0) of the next digit
		value := tail.data

		if n1 != nil {
			value += n1.data
			n1 = n1.next
		}

		if n2 != nil {
			value += n2.data
			n2 = n2.next
		}

		tail.data = value % 10
		tail.next = New(value / 10)

		if n1 == nil && n2 == nil {
			// remove leading zero
			if tail.next.data == 0 {
				tail.next = nil
			}
			// no more digits to sum
			break
		}

		tail = tail.next
	}

	return head
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
