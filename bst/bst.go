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
		if value < b.root.value {
			b.insertNodeOnLeft(b.root, value)
		} else if value > b.root.value {
			b.insertNodeOnRight(b.root, value)
		}
	}
}


func (b *BST) insertNodeOnLeft (node *Node, value int) {
	if node.left == nil {
		node.left = NewNode(value)
	} else {
		if value < node.left.value {
			b.insertNodeOnLeft(node.left, value)
		} else if value > node.left.value {
			b.insertNodeOnRight(node.left, value)
		}
	}
}

func (b *BST) insertNodeOnRight (node *Node, value int) {
	if node.right == nil {
		node.right = NewNode(value)
	} else {
		if value > node.right.value {
			b.insertNodeOnRight(node.right, value)
		} else if value < node.right.value {
			b.insertNodeOnLeft(node.right, value)
		}
	}
}

func (b *BST) Count () int {
	return b.countNodes(b.root)
}

func (b *BST) countNodes (node *Node) int {
	if node == nil {
		return 0
	} else {
		return 1 + b.countNodes(node.left) + b.countNodes(node.right)
	}
}

func (b *BST) Search(value int) bool {
	return b.searchValue(b.root, value)
}

func (b *BST) searchValue(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if value == node.value {
		return true
	} else if value < node.value {
		return b.searchValue(node.left, value)
	} else {
		return b.searchValue(node.right, value)
	}
}

func (b *BST) InOrder() {
	b.inOrderTraverseNode(b.root)
}

func (b *BST) inOrderTraverseNode(node *Node) {
	if node != nil {
		b.inOrderTraverseNode(node.left)
		fmt.Println(node.value)
		b.inOrderTraverseNode(node.right)
	}
}


func (b *BST) Remove(value int) *Node {
	return b.removeValue(b.root, value)
}


func (b *BST) removeValue(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	
	if value < node.value {
		node.left = b.removeValue(node.left, value)
	} else if value > node.value {
		node.right = b.removeValue(node.right, value)
	} else {  // value == node.Value, remover este nó
		if node.left == nil && node.right == nil {  // nó folha
			return nil
		} else if node.left != nil && node.right != nil {  // nó com dois filhos
			minNode := b.findMinValueNode(node.right)
			node.value = minNode.value
			node.right = b.removeValue(node.right, minNode.value)
		} else {  // nó com um filho
			if node.left != nil {
				return node.left
			} else {
				return node.right
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
    for current.left != nil {
        current = current.left
    }
    return current
}

func (b *BST) Max() *Node {
    return b.findMaxValueNode(b.root)
}

func (b *BST) findMaxValueNode(node *Node) *Node {
    current := node
    for current.right != nil {
        current = current.right
    }
    return current
}

func (b *BST) PrintNode(n *Node) string {
    return fmt.Sprintf(
        "Node{Value: %d, Left: %v, Right: %v}",
        n.value, n.left, n.right,
    )
}