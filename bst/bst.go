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
		} else if (value > b.root.Value){
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
		} else if (value > node.Left.Value) {
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
		} else if(value < node.Right.Value) {
			b.insertNodeOnLeft(node.Right, value)
		}
	}
}

func (b *BST) Count () int {
	return b.countNodes(b.root)
}

func (b *BST) countNodes (node *Node) int {
	if(node == nil) {
		return 0
	} else {
		return 1 + b.countNodes(node.Left) + b.countNodes(node.Right)
	}
}

func (b *BST) Search(value int) bool {
	return b.searchValue(b.root, value)
}

func (b *BST) searchValue(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	} else if value < node.Value {
		return b.searchValue(node.Left, value)
	} else {
		return b.searchValue(node.Right, value)
	}
}

