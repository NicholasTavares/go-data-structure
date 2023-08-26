package stack

import "fmt"

type Stack struct {
    items []int
    count int
}

func NewStack() *Stack {
    return &Stack{
        items: nil,
		count: 0,
    }
}

func (s *Stack) Push(element int) {
	s.items = append(s.items, element)
	s.count++
}

func (s *Stack) IsEmpty() bool {
    return s.count == 0;
}

func (s *Stack) Pop() (int, error) {
	if(s.IsEmpty()) {
		return 0, fmt.Errorf("Stack is empty")
	}
	s.count--
	result := s.items[s.count]
	s.items = s.items[:s.count]
	return result, nil

}

func (s *Stack) Peek() (int, error) {
	if (s.IsEmpty()) {
		return 0, fmt.Errorf("Stack is empty")
	}

	return s.items[s.count - 1], nil
}

func (s *Stack) Clear()  {
    s.count = 0
	s.items = nil
}

func (s *Stack) Size() int {
	return s.count
}

func (s *Stack) Display() {
	fmt.Print("[")
    for i := 0; i < s.count; i++ {
        fmt.Printf(" %d ", s.items[i])
    }
	fmt.Println("]")
}
