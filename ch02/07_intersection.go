package ch02

// Intersection returns intersection node of two lists
// If lists don't have intersection, nil will be returned
func Intersection(n1 *Node, n2 *Node) *Node {
	n1len := n1.Len()
	n2len := n2.Len()
	nlen := n1len

	if n1len > n2len {
		for i := 0; i < n1len-n2len; i++ {
			n1 = n1.next
		}

		nlen = n2len
	}

	if n2len > n1len {
		for i := 0; i < n2len-n1len; i++ {
			n2 = n2.next
		}

		nlen = n1len
	}

	for i := 0; i < nlen; i++ {
		if n1 == n2 {
			return n1
		}

		n1 = n1.next
		n2 = n2.next
	}

	return nil
}
