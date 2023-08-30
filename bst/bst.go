package bst

type BST struct {
    root *Node
}

func NewBST() *BST {
    return &BST{
        root: nil,
    }
}

func (b *BST) Insert(value int) {
    if b.root == nil {
		b.root = NewNode(value)
	} else {
		if (value < b.root.Value){
			b.insertNodeOnLeft(b.root, value)
		} else {
			b.insertNodeOnRight(b.root, value)
		}
	}
}


func (b *BST) insertNodeOnLeft (node *Node, value int) {
	if(node.Left == nil) {
		node.Left = NewNode(value)
	} else {
		if(value < node.Left.Value) {
			b.insertNodeOnLeft(node.Left, value)
		} else {
			b.insertNodeOnRight(node.Left, value)
		}
	}
}

func (b *BST) insertNodeOnRight (node *Node, value int) {
	if(node.Right == nil) {
		node.Right = NewNode(value)
	} else {
		if(value > node.Right.Value) {
			b.insertNodeOnRight(node.Right, value)
		} else {
			b.insertNodeOnLeft(node.Right, value)
		}
	}
}