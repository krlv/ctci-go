package ch02

// GetLoopHead returns head of the loop inside the linked list
// If the linked list doesn't have a loop, nil will be returned
func GetLoopHead(node *Node) *Node {
	slow := node
	fast := node

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			break
		}
	}

	if slow != fast {
		return nil
	}

	slow = node
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}

	return slow
}
