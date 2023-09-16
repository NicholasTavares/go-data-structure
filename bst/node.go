package bst

type Node struct {
    value int
    parent  *Node
    left  *Node
	right  *Node
}

func NewNode(value int) *Node {
    return &Node{
        value: value,
        parent:  nil,
        left:  nil,
		right:  nil,
    }
}
