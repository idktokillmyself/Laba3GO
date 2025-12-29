package dbmsgo

import (
	"fmt"
	"strings"
)

type Array struct {
	name string
	data []string
}

func NewArray(name string) *Array {
	return &Array{
		name: name,
		data: make([]string, 0),
	}
}

func (a *Array) PushBack(value string) {
	a.data = append(a.data, value)
}

func (a *Array) Insert(index int, value string) error {
	if index < 0 || index > len(a.data) {
		return fmt.Errorf("index out of range")
	}
	
	a.data = append(a.data, "")
	copy(a.data[index+1:], a.data[index:])
	a.data[index] = value
	return nil
}

func (a *Array) Get(index int) (string, error) {
	if index < 0 || index >= len(a.data) {
		return "", fmt.Errorf("index out of range")
	}
	return a.data[index], nil
}

func (a *Array) Remove(index int) error {
	if index < 0 || index >= len(a.data) {
		return fmt.Errorf("index out of range")
	}
	
	a.data = append(a.data[:index], a.data[index+1:]...)
	return nil
}

func (a *Array) Replace(index int, value string) error {
	if index < 0 || index >= len(a.data) {
		return fmt.Errorf("index out of range")
	}
	a.data[index] = value
	return nil
}

func (a *Array) Length() int {
	return len(a.data)
}

func (a *Array) IsEmpty() bool {
	return len(a.data) == 0
}

func (a *Array) Print() {
	fmt.Printf("Массив '%s': [%s]\n", a.name, strings.Join(a.data, ", "))
}

func (a *Array) Cleanup() {
	a.data = make([]string, 0)
}

func (a *Array) GetName() string {
	return a.name
}

func (a *Array) GetData() []string {
	return a.data
}
