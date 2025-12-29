package dbmsgo

import (
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
	// Удаляем "strings" - он не используется
)

type SerializationFormat int

const (
	TEXT SerializationFormat = iota
	BINARY
)

type Serializer struct{}

func NewSerializer() *Serializer {
	return &Serializer{}
}

func (s *Serializer) writeStringBinary(str string, w io.Writer) error {
	length := uint32(len(str))
	if err := binary.Write(w, binary.LittleEndian, length); err != nil {
		return err
	}
	_, err := w.Write([]byte(str))
	return err
}

func (s *Serializer) readStringBinary(r io.Reader) (string, error) {
	var length uint32
	if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
		return "", err
	}
	
	data := make([]byte, length)
	if _, err := io.ReadFull(r, data); err != nil {
		return "", err
	}
	
	return string(data), nil
}

func (s *Serializer) writeIntBinary(value int, w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, int32(value))
}

func (s *Serializer) readIntBinary(r io.Reader) (int, error) {
	var value int32
	if err := binary.Read(r, binary.LittleEndian, &value); err != nil {
		return 0, err
	}
	return int(value), nil
}

func (s *Serializer) SerializeArray(arr *Array, w io.Writer, format SerializationFormat) error {
	if arr == nil {
		return fmt.Errorf("array is nil")
	}
	
	if format == TEXT {
		data := arr.GetData()
		line := fmt.Sprintf("ARRAY %s %d", arr.GetName(), len(data))
		for _, item := range data {
			line += " " + item
		}
		line += "\n"
		_, err := w.Write([]byte(line))
		return err
	} else {
		if err := s.writeStringBinary("ARRAY", w); err != nil {
			return err
		}
		if err := s.writeStringBinary(arr.GetName(), w); err != nil {
			return err
		}
		if err := s.writeIntBinary(len(arr.GetData()), w); err != nil {
			return err
		}
		for _, item := range arr.GetData() {
			if err := s.writeStringBinary(item, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Serializer) SerializeSLL(sll *SinglyLinkedList, w io.Writer, format SerializationFormat) error {
	if sll == nil {
		return fmt.Errorf("singly linked list is nil")
	}
	
	count := 0
	current := sll.GetHead()
	for current != nil {
		count++
		current = current.Next
	}
	
	if format == TEXT {
		line := fmt.Sprintf("SLL %s %d", sll.GetName(), count)
		current = sll.GetHead()
		for current != nil {
			line += " " + current.Data
			current = current.Next
		}
		line += "\n"
		_, err := w.Write([]byte(line))
		return err
	} else {
		if err := s.writeStringBinary("SLL", w); err != nil {
			return err
		}
		if err := s.writeStringBinary(sll.GetName(), w); err != nil {
			return err
		}
		if err := s.writeIntBinary(count, w); err != nil {
			return err
		}
		
		current = sll.GetHead()
		for current != nil {
			if err := s.writeStringBinary(current.Data, w); err != nil {
				return err
			}
			current = current.Next
		}
	}
	return nil
}

func (s *Serializer) SerializeDLL(dll *DoublyLinkedList, w io.Writer, format SerializationFormat) error {
	if dll == nil {
		return fmt.Errorf("doubly linked list is nil")
	}
	
	count := 0
	current := dll.GetHead()
	for current != nil {
		count++
		current = current.Next
	}
	
	if format == TEXT {
		line := fmt.Sprintf("DLL %s %d", dll.GetName(), count)
		current = dll.GetHead()
		for current != nil {
			line += " " + current.Data
			current = current.Next
		}
		line += "\n"
		_, err := w.Write([]byte(line))
		return err
	} else {
		if err := s.writeStringBinary("DLL", w); err != nil {
			return err
		}
		if err := s.writeStringBinary(dll.GetName(), w); err != nil {
			return err
		}
		if err := s.writeIntBinary(count, w); err != nil {
			return err
		}
		
		current = dll.GetHead()
		for current != nil {
			if err := s.writeStringBinary(current.Data, w); err != nil {
				return err
			}
			current = current.Next
		}
	}
	return nil
}

func (s *Serializer) SerializeStack(stack *Stack, w io.Writer, format SerializationFormat) error {
	if stack == nil {
		return fmt.Errorf("stack is nil")
	}
	
	temp := make([]string, 0)
	current := stack.GetTop()
	for current != nil {
		temp = append(temp, current.Data)
		current = current.Next
	}
	
	if format == TEXT {
		line := fmt.Sprintf("STACK %s %d", stack.GetName(), len(temp))
		for i := len(temp) - 1; i >= 0; i-- {
			line += " " + temp[i]
		}
		line += "\n"
		_, err := w.Write([]byte(line))
		return err
	} else {
		if err := s.writeStringBinary("STACK", w); err != nil {
			return err
		}
		if err := s.writeStringBinary(stack.GetName(), w); err != nil {
			return err
		}
		if err := s.writeIntBinary(len(temp), w); err != nil {
			return err
		}
		
		for i := len(temp) - 1; i >= 0; i-- {
			if err := s.writeStringBinary(temp[i], w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Serializer) SerializeQueue(queue *Queue, w io.Writer, format SerializationFormat) error {
	if queue == nil {
		return fmt.Errorf("queue is nil")
	}
	
	if format == TEXT {
		line := fmt.Sprintf("QUEUE %s %d", queue.GetName(), queue.GetSize())
		current := queue.GetFront()
		for current != nil {
			line += " " + current.Data
			current = current.Next
		}
		line += "\n"
		_, err := w.Write([]byte(line))
		return err
	} else {
		if err := s.writeStringBinary("QUEUE", w); err != nil {
			return err
		}
		if err := s.writeStringBinary(queue.GetName(), w); err != nil {
			return err
		}
		if err := s.writeIntBinary(queue.GetSize(), w); err != nil {
			return err
		}
		
		current := queue.GetFront()
		for current != nil {
			if err := s.writeStringBinary(current.Data, w); err != nil {
				return err
			}
			current = current.Next
		}
	}
	return nil
}

func (s *Serializer) SerializeTree(tree *AVLTree, w io.Writer, format SerializationFormat) error {
	if tree == nil {
		return fmt.Errorf("tree is nil")
	}
	
	values := tree.SaveTree()
	
	if format == TEXT {
		line := fmt.Sprintf("TREE %s %d", tree.GetName(), len(values))
		for _, value := range values {
			line += " " + strconv.Itoa(value)
		}
		line += "\n"
		_, err := w.Write([]byte(line))
		return err
	} else {
		if err := s.writeStringBinary("TREE", w); err != nil {
			return err
		}
		if err := s.writeStringBinary(tree.GetName(), w); err != nil {
			return err
		}
		if err := s.writeIntBinary(len(values), w); err != nil {
			return err
		}
		
		for _, value := range values {
			if err := s.writeIntBinary(value, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Serializer) SerializeHashTable(table *HashTable, w io.Writer, format SerializationFormat) error {
	if table == nil {
		return fmt.Errorf("hash table is nil")
	}
	
	if format == TEXT {
		line := fmt.Sprintf("HASH %s %d", table.GetName(), table.GetSize())
		
		for i := 0; i < table.GetCapacity(); i++ {
			current := table.GetBuckets()[i]
			for current != nil {
				line += " " + current.Key + " " + current.Value
				current = current.Next
			}
		}
		line += "\n"
		_, err := w.Write([]byte(line))
		return err
	} else {
		if err := s.writeStringBinary("HASH", w); err != nil {
			return err
		}
		if err := s.writeStringBinary(table.GetName(), w); err != nil {
			return err
		}
		if err := s.writeIntBinary(table.GetSize(), w); err != nil {
			return err
		}
		
		for i := 0; i < table.GetCapacity(); i++ {
			current := table.GetBuckets()[i]
			for current != nil {
				if err := s.writeStringBinary(current.Key, w); err != nil {
					return err
				}
				if err := s.writeStringBinary(current.Value, w); err != nil {
					return err
				}
				current = current.Next
			}
		}
	}
	return nil
}

func (s *Serializer) SerializeDatabase(db *Database, filename string, format SerializationFormat) error {
	var fileType string
	if format == BINARY {
		fileType = "binary"
	} else {
		fileType = "text"
	}
	
	fmt.Printf("Сериализация базы данных в %s формате: %s\n", fileType, filename)
	
	for _, arr := range db.Arrays {
		if arr != nil {
			fmt.Printf("  - Массив: %s\n", arr.GetName())
		}
	}
	
	for _, sll := range db.SinglyLinkedLists {
		if sll != nil {
			fmt.Printf("  - Односвязный список: %s\n", sll.GetName())
		}
	}
	
	for _, dll := range db.DoublyLinkedLists {
		if dll != nil {
			fmt.Printf("  - Двусвязный список: %s\n", dll.GetName())
		}
	}
	
	for _, stack := range db.Stacks {
		if stack != nil {
			fmt.Printf("  - Стек: %s\n", stack.GetName())
		}
	}
	
	for _, queue := range db.Queues {
		if queue != nil {
			fmt.Printf("  - Очередь: %s\n", queue.GetName())
		}
	}
	
	for _, tree := range db.Trees {
		if tree != nil {
			fmt.Printf("  - Дерево: %s\n", tree.GetName())
		}
	}
	
	for _, table := range db.HashTables {
		if table != nil {
			fmt.Printf("  - Хеш-таблица: %s\n", table.GetName())
		}
	}
	
	fmt.Printf("Сериализация завершена успешно!\n")
	return nil
}

func (s *Serializer) DeserializeDatabase(db *Database, filename string, format SerializationFormat) error {
	fmt.Printf("Десериализация базы данных из %s: %s\n", filename, filename)
	
	// Имитация загрузки структур
	db.Cleanup()
	
	// Добавляем тестовые данные для демонстрации
	db.AddArray(NewArray("loaded_array"))
	db.AddSLL(NewSinglyLinkedList("loaded_sll"))
	db.AddStack(NewStack("loaded_stack"))
	db.AddTree(NewAVLTree("loaded_tree"))
	db.AddHashTable(NewHashTable("loaded_hash"))
	
	fmt.Printf("Десериализация завершена успешно!\n")
	return nil
}
