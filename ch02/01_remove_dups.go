package ch02

// RemoveDups removes duplicates from linked list
// Uses additional map to track duplicates
func RemoveDups(n *Node) {
	var prev *Node
	dups := make(map[int]bool)

	for n != nil {
		d := n.data

		if _, ok := dups[d]; ok {
			// n.Remove
			prev.next = n.next
		} else {
			dups[d] = true
			prev = n
		}

		n = n.next
	}
}

// RemoveDupsRunner removes duplicates from linked list
// Uses runner technique, no need in additional data structures
func RemoveDupsRunner(n *Node) {
	for n != nil {
		r := n
		for r.next != nil {
			if r.next.data == n.data {
				r.next = r.next.next
			} else {
				r = r.next
			}
		}

		n = n.next
	}
}
