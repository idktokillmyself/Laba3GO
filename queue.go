package dbmsgo

import (
	"fmt"
)

type QueueNode struct {
	Data string
	Next *QueueNode
}

type Queue struct {
	name  string
	front *QueueNode
	rear  *QueueNode
	size  int
}

func NewQueue(name string) *Queue {
	return &Queue{
		name:  name,
		front: nil,
		rear:  nil,
		size:  0,
	}
}

func (q *Queue) Push(value string) {
	newNode := &QueueNode{Data: value, Next: nil}
	if q.IsEmpty() {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.Next = newNode
		q.rear = newNode
	}
	q.size++
}

func (q *Queue) Pop() (string, error) {
	if q.IsEmpty() {
		return "", fmt.Errorf("queue is empty")
	}
	
	value := q.front.Data
	q.front = q.front.Next
	if q.front == nil {
		q.rear = nil
	}
	q.size--
	return value, nil
}

func (q *Queue) Peek() (string, error) {
	if q.IsEmpty() {
		return "", fmt.Errorf("queue is empty")
	}
	return q.front.Data, nil
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

func (q *Queue) Print() {
	fmt.Printf("Очередь '%s': [", q.name)
	current := q.front
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

func (q *Queue) GetSize() int {
	return q.size
}

func (q *Queue) GetName() string {
	return q.name
}

func (q *Queue) GetFront() *QueueNode {
	return q.front
}

func (q *Queue) GetRear() *QueueNode {
	return q.rear
}

func (q *Queue) Cleanup() {
	for !q.IsEmpty() {
		q.Pop()
	}
}
