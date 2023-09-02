package bst

import "fmt"

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

func (b *BST) InOrder() {
	b.inOrderTraverseNode(b.root)
}

func (b *BST) inOrderTraverseNode(node *Node) {
	if node != nil {
		b.inOrderTraverseNode(node.Left)
		fmt.Println(node.Value)
		b.inOrderTraverseNode(node.Right)
	}
}


func (b *BST) Remove(value int) *Node {
	return b.removeValue(b.root, value)
}


func (b *BST) removeValue(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	
	if value < node.Value {
		node.Left = b.removeValue(node.Left, value)
	} else if value > node.Value {
		node.Right = b.removeValue(node.Right, value)
	} else {  // value == node.Value, remover este nó
		if node.Left == nil && node.Right == nil {  // nó folha
			return nil
		} else if node.Left != nil && node.Right != nil {  // nó com dois filhos
			minNode := b.findMinValueNode(node.Right)
			node.Value = minNode.Value
			node.Right = b.removeValue(node.Right, minNode.Value)
		} else {  // nó com um filho
			if node.Left != nil {
				return node.Left
			} else {
				return node.Right
			}
		}
	}
	return node
}

func (b *BST) Min() *Node {
    return b.findMinValueNode(b.root)
}

func (b *BST) findMinValueNode(node *Node) *Node {
    current := node
    for current.Left != nil {
        current = current.Left
    }
    return current
}

func (b *BST) Max() *Node {
    return b.findMaxValueNode(b.root)
}

func (b *BST) findMaxValueNode(node *Node) *Node {
    current := node
    for current.Right != nil {
        current = current.Right
    }
    return current
}
