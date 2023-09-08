package bst

type Node struct {
    value int
    left  *Node
	right  *Node
}

func NewNode(value int) *Node {
    return &Node{
        value: value,
        left:  nil,
		right:  nil,
    }
}
