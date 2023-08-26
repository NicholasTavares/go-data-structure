package linkedList

import "fmt"

type LinkedList struct {
    head *Node
    count int
}

func NewLinkedList() *LinkedList {
    return &LinkedList{
        head: nil,
    }
}

func (ll *LinkedList) Push(value int) {
    newNode := NewNode(value)
    if ll.head == nil {
        ll.head = newNode
    } else {
        current := ll.head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
    ll.count++
}

func (ll *LinkedList) Insert(index, value int) bool {
    if (index >= 0 && index <= ll.count) {
        newNode := NewNode(value)

        if(index == 0) {
            current := ll.head
            newNode.Next = current
            ll.head = newNode
        } else {
            prev, err := ll.getElementAt(index - 1)
            if (err != nil) {
                return false
            }

            current := prev.Next
            newNode.Next = current;
            prev.Next = newNode
        }

        ll.count++
        return true
    }
    return true
}

func (ll *LinkedList) Remove(value int) (int, error) {
    index := ll.indexOf(value)
    value, err := ll.removeAt(index)
    if(err != nil) {
        return -1, fmt.Errorf("Invalid position")
    }

    return value, nil
}

func (ll *LinkedList) indexOf(value int) int {
    current := ll.head
    for i := 0; i < ll.count && current != nil; i++ {
        if(value == current.Value) {
            return i
        }
        current = current.Next
    }

    return -1
}

func (ll *LinkedList) removeAt(index int) (int, error) {
    if index < 0 || index >= ll.count {
        return -1, fmt.Errorf("Invalid position")
    }

    current := ll.head
    if(index == 0) {
        ll.head = current.Next
        ll.count--
        return ll.head.Value, nil
    } else {
        prev, err := ll.getElementAt(index - 1)
        if (err != nil) {
            return -1, fmt.Errorf("Invalid position") 
        }
        current = prev.Next
        prev.Next = current.Next

        ll.count--
        return current.Next.Value, nil
    }
}

func (ll *LinkedList) getElementAt(index int) (*Node, error) {
    if index < 0 || index >= ll.count {
        return nil, fmt.Errorf("Invalid position")
    }

    node := ll.head
    for i := 0; i < index && node != nil; i++ {
        node = node.Next
    }

    return node, nil
}

func (ll *LinkedList) Size() int {
    return ll.count;
}

func (ll *LinkedList) IsEmpty() bool {
    return ll.count == 0;
}

func (ll *LinkedList) Display() {
    current := ll.head
    for current != nil {
        fmt.Println(current.Value)
        current = current.Next
    }
}
