package main

import (
	"fmt"
	"go-data-structure/linkedList"
)

func main() {
    ll := linkedList.NewLinkedList()
	ll.Push(50)
	ll.Push(51)
	ll.Push(8)
	ll.Push(53)
	ll.Insert(4 ,61)
	ll.Insert(4 ,62)
	ll.Push(50)
	ll.Remove(50)
	ll.Remove(8)
	ll.Push(11)
	ll.Display()
	fmt.Printf("Tamanho da lista: %d\n", ll.Size())
	fmt.Printf("A list está vazia? %t\n", ll.IsEmpty())
}
