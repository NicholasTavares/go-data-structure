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

func (b *BST) Successor(node *Node) *Node {
	var aux *Node
	if node.right != nil {
		return b.findMinValueNode(node.right)
	} else {
		aux = node.parent
		for aux != nil && node == aux.right {
			node = aux
			aux = aux.parent
		}
		return aux
	}
}

func (b *BST) transplant(u, v *Node) {
	if u.parent == nil {
		b.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}

func (b *BST) Delete(value int) {
	node := b.Search(value)
	var aux *Node
	if node != nil {
		if node.left == nil {
			b.transplant(node, node.right)
		} else if node.right == nil {
			b.transplant(node, node.left)
		} else {
			aux = b.findMinValueNode(node.right)
			if (aux != node.right) {
				b.transplant(aux, aux.right)
				aux.right = node.right
				aux.right.parent = aux
			}
			b.transplant(node, aux)
			aux.left = node.left
			aux.left.parent = aux
		}
	}
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