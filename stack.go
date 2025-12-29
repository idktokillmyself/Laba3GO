package dbmsgo

import (
	"fmt"
)

type StackNode struct {
	Data string
	Next *StackNode
}

type Stack struct {
	name string
	top  *StackNode
	size int
}

func NewStack(name string) *Stack {
	return &Stack{
		name: name,
		top:  nil,
		size: 0,
	}
}

func (s *Stack) Push(value string) {
	newNode := &StackNode{Data: value, Next: s.top}
	s.top = newNode
	s.size++
}

func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("stack is empty")
	}
	
	value := s.top.Data
	s.top = s.top.Next
	s.size--
	return value, nil
}

func (s *Stack) Peek() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("stack is empty")
	}
	return s.top.Data, nil
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Print() {
	fmt.Printf("Стек '%s': [", s.name)
	current := s.top
	first := true
	for current != nil {
		if !first {
			fmt.Printf(", ")
		}
		fmt.Printf("%s", current.Data)
		current = current.Next
		first = false
	}
	fmt.Printf("]\n")
}

func (s *Stack) GetSize() int {
	return s.size
}

func (s *Stack) GetName() string {
	return s.name
}

func (s *Stack) GetTop() *StackNode {
	return s.top
}

func (s *Stack) Cleanup() {
	for !s.IsEmpty() {
		s.Pop()
	}
}
