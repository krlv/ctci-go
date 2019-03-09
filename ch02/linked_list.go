package ch02

// Node struct for linked list
type Node struct {
	value int
	next  *Node
}

// New returns a new node
func New(v int) *Node {
	n := new(Node)
	n.value = v
	return n
}

// AppendToTail appends a new node with value v at the tail of list l
func (node *Node) AppendToTail(v int) {
	end := New(v)
	for node.next != nil {
		node = node.next
	}
	node.next = end
}
