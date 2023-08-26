package main

import (
	"fmt"
	"go-data-structure/stack"
)

func main() {
	s := stack.NewStack()
	s.Push(10)
	s.Push(15)
	s.Push(19)
	s.Push(16)
	s.Push(41)
	s.Push(36)
	s.Push(50)
	s.Push(39)
	s.Display()
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())
	s.Display()
	s.Clear()
	s.Display()
}
