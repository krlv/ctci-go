package ch02

// Node struct for linked list
type Node struct {
	data int
	next *Node
}

// New returns a new node
func New(d int) *Node {
	n := new(Node)
	n.data = d
	return n
}

// AppendToTail appends a new node with data d at the tail of list l
func (node *Node) AppendToTail(d int) {
	end := New(d)
	for node.next != nil {
		node = node.next
	}
	node.next = end
}

// DeleteNode deletes node with data d from the list
func (node *Node) DeleteNode(d int) {
	for n := node; n != nil; n = n.next {
		if n.data == d {
			if n.next != nil {
				n.data = n.next.data
				n.next = n.next.next
			} else {
				n.data = 0
				n.next = nil
			}

			return
		}
	}
}
