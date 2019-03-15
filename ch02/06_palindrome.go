package ch02

// IsPalindrome checks if the linked list is a palindrome
func IsPalindrome(node *Node) bool {
	// head of the reverse list
	head := New(0)

	for n := node; n != nil; n = n.next {
		head.data = n.data
		temp := New(0)
		temp.next = head

		head = temp
	}

	h := head.next
	for n := node; n != nil; n = n.next {
		if h.data != n.data {
			return false
		}

		h = h.next
	}

	return true
}
