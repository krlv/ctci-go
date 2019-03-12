package ch02

// Partition partitions linked list around value k
func Partition(node *Node, k int) *Node {
	head := node
	tail := head

	for n := node.next; n != nil; n = n.next {
		temp := New(n.data)

		if temp.data < k {
			temp.next = head
			head = temp
		} else {
			tail.next = temp
			tail = tail.next
		}
	}

	return head
}
