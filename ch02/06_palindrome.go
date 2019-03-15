package ch02

import "errors"

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

// IsPalindromeIterative checks if the linked list is a palindrome
// Implementation is based on Node.Len() method
// and uses a list as a storage for reverse values
func IsPalindromeIterative(node *Node) bool {
	nlen := node.Len()
	xlen := nlen / 2
	revl := make([]int, xlen)

	tail := node
	for i := 0; i < xlen; i++ {
		revl[xlen-1-i] = tail.data
		tail = tail.next
	}

	if nlen%2 == 1 {
		tail = tail.next
	}

	for i := 0; i < xlen; i++ {
		if revl[i] != tail.data {
			return false
		}

		tail = tail.next
	}

	return true
}

// IsPalindromeRecursive checks if the linked list is a palindrome
// Recursive-based implementation
func IsPalindromeRecursive(node *Node) bool {
	nlen := node.Len()

	_, err := getPalindromePairNode(node, nlen)
	if err != nil {
		return false
	}

	return true
}

// getPalindromePairNode returns the (n-x) pair for x node
// or error if the list is not a palindrome
func getPalindromePairNode(node *Node, nlen int) (*Node, error) {
	// this is the center of even-length palindrome
	if nlen == 0 {
		return node, nil
	}

	// this is the center of even-length palindrome
	if nlen == 1 {
		return node.next, nil
	}

	pair, err := getPalindromePairNode(node.next, nlen-2)
	if err != nil {
		return nil, err
	}

	if pair.data != node.data {
		return nil, errors.New("Not a palindrome")
	}

	return pair.next, nil
}
