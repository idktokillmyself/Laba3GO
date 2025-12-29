package dbmsgo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileIO struct {
	serializer *Serializer
}

func NewFileIO() *FileIO {
	return &FileIO{
		serializer: NewSerializer(),
	}
}

func (f *FileIO) loadArray(db *Database, parts []string) {
	if len(parts) < 3 {
		return
	}
	
	name := parts[1]
	size, err := strconv.Atoi(parts[2])
	if err != nil {
		return
	}
	
	if len(parts) < 3+size {
		return
	}
	
	arr := NewArray(name)
	for i := 0; i < size; i++ {
		arr.PushBack(parts[3+i])
	}
	db.AddArray(arr)
}

func (f *FileIO) loadSLL(db *Database, parts []string) {
	if len(parts) < 3 {
		return
	}
	
	name := parts[1]
	size, err := strconv.Atoi(parts[2])
	if err != nil {
		return
	}
	
	if len(parts) < 3+size {
		return
	}
	
	sll := NewSinglyLinkedList(name)
	for i := 0; i < size; i++ {
		sll.PushBack(parts[3+i])
	}
	db.AddSLL(sll)
}

func (f *FileIO) loadDLL(db *Database, parts []string) {
	if len(parts) < 3 {
		return
	}
	
	name := parts[1]
	size, err := strconv.Atoi(parts[2])
	if err != nil {
		return
	}
	
	if len(parts) < 3+size {
		return
	}
	
	dll := NewDoublyLinkedList(name)
	for i := 0; i < size; i++ {
		dll.PushBack(parts[3+i])
	}
	db.AddDLL(dll)
}

func (f *FileIO) loadStack(db *Database, parts []string) {
	if len(parts) < 3 {
		return
	}
	
	name := parts[1]
	size, err := strconv.Atoi(parts[2])
	if err != nil {
		return
	}
	
	if len(parts) < 3+size {
		return
	}
	
	stack := NewStack(name)
	temp := make([]string, size)
	for i := 0; i < size; i++ {
		temp[i] = parts[3+i]
	}
	
	for i := size - 1; i >= 0; i-- {
		stack.Push(temp[i])
	}
	db.AddStack(stack)
}

func (f *FileIO) loadQueue(db *Database, parts []string) {
	if len(parts) < 3 {
		return
	}
	
	name := parts[1]
	size, err := strconv.Atoi(parts[2])
	if err != nil {
		return
	}
	
	if len(parts) < 3+size {
		return
	}
	
	queue := NewQueue(name)
	for i := 0; i < size; i++ {
		queue.Push(parts[3+i])
	}
	db.AddQueue(queue)
}

func (f *FileIO) loadTree(db *Database, parts []string) {
	if len(parts) < 3 {
		return
	}
	
	name := parts[1]
	size, err := strconv.Atoi(parts[2])
	if err != nil {
		return
	}
	
	if len(parts) < 3+size {
		return
	}
	
	tree := NewAVLTree(name)
	for i := 0; i < size; i++ {
		value, err := strconv.Atoi(parts[3+i])
		if err == nil {
			tree.Insert(value)
		}
	}
	db.AddTree(tree)
}

func (f *FileIO) loadHashTable(db *Database, parts []string) {
	if len(parts) < 3 {
		return
	}
	
	name := parts[1]
	size, err := strconv.Atoi(parts[2])
	if err != nil {
		return
	}
	
	if len(parts) < 3+size*2 {
		return
	}
	
	table := NewHashTable(name)
	for i := 0; i < size; i++ {
		key := parts[3+i*2]
		value := parts[3+i*2+1]
		table.Insert(key, value)
	}
	db.AddHashTable(table)
}

func (f *FileIO) SaveDatabaseToFile(db *Database, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Сохраняем массивы
	for _, arr := range db.Arrays {
		if arr != nil {
			if err := f.serializer.SerializeArray(arr, writer, TEXT); err != nil {
				return err
			}
		}
	}

	// Сохраняем односвязные списки
	for _, sll := range db.SinglyLinkedLists {
		if sll != nil {
			if err := f.serializer.SerializeSLL(sll, writer, TEXT); err != nil {
				return err
			}
		}
	}

	// Сохраняем двусвязные списки
	for _, dll := range db.DoublyLinkedLists {
		if dll != nil {
			if err := f.serializer.SerializeDLL(dll, writer, TEXT); err != nil {
				return err
			}
		}
	}

	// Сохраняем стеки
	for _, stack := range db.Stacks {
		if stack != nil {
			if err := f.serializer.SerializeStack(stack, writer, TEXT); err != nil {
				return err
			}
		}
	}

	// Сохраняем очереди
	for _, queue := range db.Queues {
		if queue != nil {
			if err := f.serializer.SerializeQueue(queue, writer, TEXT); err != nil {
				return err
			}
		}
	}

	// Сохраняем деревья
	for _, tree := range db.Trees {
		if tree != nil {
			if err := f.serializer.SerializeTree(tree, writer, TEXT); err != nil {
				return err
			}
		}
	}

	// Сохраняем хеш-таблицы
	for _, table := range db.HashTables {
		if table != nil {
			if err := f.serializer.SerializeHashTable(table, writer, TEXT); err != nil {
				return err
			}
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	fmt.Printf("База данных сохранена в файл: %s\n", filename)
	return nil
}

func (f *FileIO) LoadDatabaseFromFile(db *Database, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	db.Cleanup()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		switch parts[0] {
		case "ARRAY":
			f.loadArray(db, parts)
		case "SLL":
			f.loadSLL(db, parts)
		case "DLL":
			f.loadDLL(db, parts)
		case "STACK":
			f.loadStack(db, parts)
		case "QUEUE":
			f.loadQueue(db, parts)
		case "TREE":
			f.loadTree(db, parts)
		case "HASH":
			f.loadHashTable(db, parts)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Printf("База данных загружена из файла: %s\n", filename)
	return nil
}
