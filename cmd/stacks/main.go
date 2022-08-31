package main

import "fmt"

// Stack : struct that implements stacks data structure.
type Stack struct {
	Items []int
}

// Push : pushes data into the stack
func (s *Stack) Push(item int) {
	s.Items = append(s.Items, item)
}

// Pop : removes the last data that was pushed into the stack
func (s *Stack) Pop() int {
	if len(s.Items) > 0 {
		l := len(s.Items) - 1
		ToRemove := s.Items[l]
		s.Items = s.Items[:l]
		return ToRemove
	}
	return -1
}

func main() {
	rr := Stack{}
	rr.Push(100)
	rr.Push(200)
	rr.Push(300)
	rr.Push(400)
	fmt.Println("Stack :", rr)
	fmt.Println("Popped value :", rr.Pop())
	fmt.Println("Stack after pop:", rr)
	fmt.Println("Popped value :", rr.Pop())
	fmt.Println("Popped value :", rr.Pop())
	fmt.Println("Popped value :", rr.Pop())
	fmt.Println("Popped value :", rr.Pop())
	fmt.Println("Popped value :", rr.Pop())
}
