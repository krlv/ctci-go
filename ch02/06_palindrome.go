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

// IsPalindromeList checks if the linked list is a palindrome
// Implementation is based on Node.Len() method
// and uses a list as a storage for reverse values
func IsPalindromeList(node *Node) bool {
	nlen := node.Len()
	revl := make([]int, nlen)

	tail := node
	for i := 0; i < nlen; i++ {
		revl[nlen-1-i] = tail.data
		tail = tail.next
	}

	tail = node
	for i := 0; i < nlen; i++ {
		if revl[i] != tail.data {
			return false
		}

		tail = tail.next
	}

	return true
}
