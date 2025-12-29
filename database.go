package dbmsgo

type Database struct {
	Arrays           []*Array
	SinglyLinkedLists []*SinglyLinkedList
	DoublyLinkedLists []*DoublyLinkedList
	Stacks           []*Stack
	Queues           []*Queue
	Trees            []*AVLTree
	HashTables       []*HashTable
}

func NewDatabase() *Database {
	return &Database{
		Arrays:           make([]*Array, 0),
		SinglyLinkedLists: make([]*SinglyLinkedList, 0),
		DoublyLinkedLists: make([]*DoublyLinkedList, 0),
		Stacks:           make([]*Stack, 0),
		Queues:           make([]*Queue, 0),
		Trees:            make([]*AVLTree, 0),
		HashTables:       make([]*HashTable, 0),
	}
}

func (d *Database) FindArray(name string) *Array {
	for _, arr := range d.Arrays {
		if arr.GetName() == name {
			return arr
		}
	}
	return nil
}

func (d *Database) FindSLL(name string) *SinglyLinkedList {
	for _, sll := range d.SinglyLinkedLists {
		if sll.GetName() == name {
			return sll
		}
	}
	return nil
}

func (d *Database) FindDLL(name string) *DoublyLinkedList {
	for _, dll := range d.DoublyLinkedLists {
		if dll.GetName() == name {
			return dll
		}
	}
	return nil
}

func (d *Database) FindStack(name string) *Stack {
	for _, stack := range d.Stacks {
		if stack.GetName() == name {
			return stack
		}
	}
	return nil
}

func (d *Database) FindQueue(name string) *Queue {
	for _, queue := range d.Queues {
		if queue.GetName() == name {
			return queue
		}
	}
	return nil
}

func (d *Database) FindTree(name string) *AVLTree {
	for _, tree := range d.Trees {
		if tree.GetName() == name {
			return tree
		}
	}
	return nil
}

func (d *Database) FindHashTable(name string) *HashTable {
	for _, table := range d.HashTables {
		if table.GetName() == name {
			return table
		}
	}
	return nil
}

func (d *Database) AddArray(arr *Array) {
	d.Arrays = append(d.Arrays, arr)
}

func (d *Database) AddSLL(sll *SinglyLinkedList) {
	d.SinglyLinkedLists = append(d.SinglyLinkedLists, sll)
}

func (d *Database) AddDLL(dll *DoublyLinkedList) {
	d.DoublyLinkedLists = append(d.DoublyLinkedLists, dll)
}

func (d *Database) AddStack(stack *Stack) {
	d.Stacks = append(d.Stacks, stack)
}

func (d *Database) AddQueue(queue *Queue) {
	d.Queues = append(d.Queues, queue)
}

func (d *Database) AddTree(tree *AVLTree) {
	d.Trees = append(d.Trees, tree)
}

func (d *Database) AddHashTable(table *HashTable) {
	d.HashTables = append(d.HashTables, table)
}

func (d *Database) Cleanup() {
	d.Arrays = make([]*Array, 0)
	d.SinglyLinkedLists = make([]*SinglyLinkedList, 0)
	d.DoublyLinkedLists = make([]*DoublyLinkedList, 0)
	d.Stacks = make([]*Stack, 0)
	d.Queues = make([]*Queue, 0)
	d.Trees = make([]*AVLTree, 0)
	d.HashTables = make([]*HashTable, 0)
}
