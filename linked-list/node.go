package linkedlist

type Node struct {
    Key   int
    Left  *Node
    Right *Node
}

func NewNode(key int) *Node {
    return &Node{
        Key: key,
        Left: nil,
        Right: nil,
    }
}
