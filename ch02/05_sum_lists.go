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
// Solves the problem with recursion approach
func SumListsReverse(n1 *Node, n2 *Node) *Node {
	n1Len := n1.Len()
	n2Len := n2.Len()

	if n1Len > n2Len {
		n2.PadLeft(0, n1Len-n2Len)
	}

	if n1Len < n2Len {
		n1.PadLeft(0, n2Len-n1Len)
	}

	head := sumListsReverseRecursion(n1, n2)

	// remove leading zero
	if head.data == 0 {
		head = head.next
	}

	return head
}

// sumListsReverseRecursion returns sum of n1 and n2 numbers as a linked list
// Numbers n1, n2 and resulted sum are stored in forward order
// Numbers n1, n2 should be the same length
func sumListsReverseRecursion(n1 *Node, n2 *Node) *Node {
	value := n1.data + n2.data

	if n1.next == nil && n2.next == nil {
		head := New(value / 10)
		head.next = New(value % 10)
		return head
	}

	tail := sumListsReverseRecursion(n1.next, n2.next)
	tail.data += value % 10

	head := New(value / 10)
	head.next = tail

	return head
}
