package dbmsgo

import "fmt"

type SLLNode struct {
	Data string
	Next *SLLNode
}

type SinglyLinkedList struct {
	name string
	head *SLLNode
	tail *SLLNode
}

func NewSinglyLinkedList(name string) *SinglyLinkedList {
	return &SinglyLinkedList{
		name: name,
		head: nil,
		tail: nil,
	}
}

func (s *SinglyLinkedList) PushFront(value string) {
	newNode := &SLLNode{Data: value, Next: s.head}
	s.head = newNode
	if s.tail == nil {
		s.tail = newNode
	}
}

func (s *SinglyLinkedList) PushBack(value string) {
	newNode := &SLLNode{Data: value, Next: nil}
	if s.tail == nil {
		s.head = newNode
		s.tail = newNode
	} else {
		s.tail.Next = newNode
		s.tail = newNode
	}
}

func (s *SinglyLinkedList) InsertBefore(target, value string) {
	if s.head == nil {
		return
	}
	
	if s.head.Data == target {
		s.PushFront(value)
		return
	}
	
	current := s.head
	for current.Next != nil && current.Next.Data != target {
		current = current.Next
	}
	
	if current.Next != nil {
		newNode := &SLLNode{Data: value, Next: current.Next}
		current.Next = newNode
	}
}

func (s *SinglyLinkedList) InsertAfter(target, value string) {
	current := s.head
	for current != nil && current.Data != target {
		current = current.Next
	}
	
	if current != nil {
		newNode := &SLLNode{Data: value, Next: current.Next}
		current.Next = newNode
		if current == s.tail {
			s.tail = newNode
		}
	}
}

func (s *SinglyLinkedList) DeleteFront() {
	if s.head == nil {
		return
	}
	
	s.head = s.head.Next
	if s.head == nil {
		s.tail = nil
	}
}

func (s *SinglyLinkedList) DeleteBack() {
	if s.head == nil {
		return
	}
	
	if s.head.Next == nil {
		s.head = nil
		s.tail = nil
		return
	}
	
	current := s.head
	for current.Next != s.tail {
		current = current.Next
	}
	
	current.Next = nil
	s.tail = current
}

func (s *SinglyLinkedList) DeleteByValue(value string) {
	// Удаляем все вхождения с начала
	for s.head != nil && s.head.Data == value {
		s.DeleteFront()
	}
	
	if s.head == nil {
		return
	}
	
	current := s.head
	for current.Next != nil {
		if current.Next.Data == value {
			current.Next = current.Next.Next
			if current.Next == nil {
				s.tail = current
			}
		} else {
			current = current.Next
		}
	}
}

func (s *SinglyLinkedList) FindByValue(value string) *SLLNode {
	current := s.head
	for current != nil {
		if current.Data == value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (s *SinglyLinkedList) Print() {
	fmt.Printf("Односвязный список '%s': ", s.name)
	current := s.head
	for current != nil {
		fmt.Printf("%s", current.Data)
		if current.Next != nil {
			fmt.Printf(" -> ")
		}
		current = current.Next
	}
	fmt.Printf(" -> NULL\n")
}

func (s *SinglyLinkedList) IsEmpty() bool {
	return s.head == nil
}

func (s *SinglyLinkedList) GetName() string {
	return s.name
}

func (s *SinglyLinkedList) GetHead() *SLLNode {
	return s.head
}

func (s *SinglyLinkedList) GetTail() *SLLNode {
	return s.tail
}

func (s *SinglyLinkedList) Cleanup() {
	s.head = nil
	s.tail = nil
}
