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
	node := NewNode(value)
	var prev *Node = nil;
	x := b.root;

	for x != nil { // achando o nó que vai receber o filho
		prev = x
		if node.value < x.value {
			x = x.left
		} else {
			x = x.right
		}
	}
	
	node.parent = prev; 

	if prev == nil { // árvore vazia
		b.root = node;
	} else if node.value < prev.value { // condicionais para decidir se o nó filho fica na esquerda ou direita
		prev.left = node
	} else {
		prev.right = node
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

func (b *BST) Search(value int) *Node {
	x := b.root
	for x != nil && value != x.value {
		if value < x.value {
			x = x.left
		} else {
			x = x.right
		}
	}
	
	return x;
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
        "Node{Value: %d, Parent: %v, Left: %v, Right: %v}",
        n.value, n.parent, n.left, n.right,
    )
}