package dbmsgo

import "fmt"

type DLLNode struct {
	Data string
	Prev *DLLNode
	Next *DLLNode
}

type DoublyLinkedList struct {
	name string
	head *DLLNode
	tail *DLLNode
}

func NewDoublyLinkedList(name string) *DoublyLinkedList {
	return &DoublyLinkedList{
		name: name,
		head: nil,
		tail: nil,
	}
}

func (d *DoublyLinkedList) PushFront(value string) {
	newNode := &DLLNode{Data: value, Prev: nil, Next: d.head}
	if d.head != nil {
		d.head.Prev = newNode
	}
	d.head = newNode
	if d.tail == nil {
		d.tail = newNode
	}
}

func (d *DoublyLinkedList) PushBack(value string) {
	newNode := &DLLNode{Data: value, Prev: d.tail, Next: nil}
	if d.tail != nil {
		d.tail.Next = newNode
	}
	d.tail = newNode
	if d.head == nil {
		d.head = newNode
	}
}

func (d *DoublyLinkedList) InsertBefore(target, value string) {
	current := d.head
	for current != nil && current.Data != target {
		current = current.Next
	}
	
	if current != nil {
		newNode := &DLLNode{Data: value, Prev: current.Prev, Next: current}
		if current.Prev != nil {
			current.Prev.Next = newNode
		} else {
			d.head = newNode
		}
		current.Prev = newNode
	}
}

func (d *DoublyLinkedList) InsertAfter(target, value string) {
	current := d.head
	for current != nil && current.Data != target {
		current = current.Next
	}
	
	if current != nil {
		newNode := &DLLNode{Data: value, Prev: current, Next: current.Next}
		if current.Next != nil {
			current.Next.Prev = newNode
		} else {
			d.tail = newNode
		}
		current.Next = newNode
	}
}

func (d *DoublyLinkedList) DeleteFront() {
	if d.head == nil {
		return
	}
	
	d.head = d.head.Next
	if d.head != nil {
		d.head.Prev = nil
	} else {
		d.tail = nil
	}
}

func (d *DoublyLinkedList) DeleteBack() {
	if d.tail == nil {
		return
	}
	
	d.tail = d.tail.Prev
	if d.tail != nil {
		d.tail.Next = nil
	} else {
		d.head = nil
	}
}

func (d *DoublyLinkedList) DeleteByValue(value string) {
	current := d.head
	for current != nil {
		if current.Data == value {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				d.head = current.Next
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				d.tail = current.Prev
			}
		}
		current = current.Next
	}
}

func (d *DoublyLinkedList) FindByValue(value string) *DLLNode {
	current := d.head
	for current != nil {
		if current.Data == value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (d *DoublyLinkedList) PrintForward() {
	fmt.Printf("Двусвязный список '%s' (прямой): ", d.name)
	current := d.head
	for current != nil {
		fmt.Printf("%s", current.Data)
		if current.Next != nil {
			fmt.Printf(" <-> ")
		}
		current = current.Next
	}
	fmt.Printf(" -> NULL\n")
}

func (d *DoublyLinkedList) PrintBackward() {
	fmt.Printf("Двусвязный список '%s' (обратный): ", d.name)
	current := d.tail
	for current != nil {
		fmt.Printf("%s", current.Data)
		if current.Prev != nil {
			fmt.Printf(" <-> ")
		}
		current = current.Prev
	}
	fmt.Printf(" -> NULL\n")
}

func (d *DoublyLinkedList) IsEmpty() bool {
	return d.head == nil
}

func (d *DoublyLinkedList) GetName() string {
	return d.name
}

func (d *DoublyLinkedList) GetHead() *DLLNode {
	return d.head
}

func (d *DoublyLinkedList) GetTail() *DLLNode {
	return d.tail
}

func (d *DoublyLinkedList) Cleanup() {
	d.head = nil
	d.tail = nil
}
