package dbmsgo

import (
	"fmt"
	"strconv"
	"strings"
)

type CommandParser struct {
	db     *Database
	fileIO *FileIO
}

func NewCommandParser(db *Database) *CommandParser {
	return &CommandParser{
		db:     db,
		fileIO: NewFileIO(),
	}
}


func (p *CommandParser) ProcessCommand(command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}

	action := parts[0]

	switch action {
	case "CREATE":
		p.handleCreate(parts[1:])
	case "MPUSH":
		p.handleMPush(parts[1:])
	case "MINSERT":
		p.handleMInsert(parts[1:])
	case "MGET":
		p.handleMGet(parts[1:])
	case "MDEL":
		p.handleMDel(parts[1:])
	case "MREPLACE":
		p.handleMReplace(parts[1:])
	case "MLENGTH":
		p.handleMLength(parts[1:])
	case "FPUSH_FRONT":
		p.handleFPushFront(parts[1:])
	case "FPUSH_BACK":
		p.handleFPushBack(parts[1:])
	case "FINSERT_BEFORE":
		p.handleFInsertBefore(parts[1:])
	case "FINSERT_AFTER":
		p.handleFInsertAfter(parts[1:])
	case "FDEL_FRONT":
		p.handleFDelFront(parts[1:])
	case "FDEL_BACK":
		p.handleFDelBack(parts[1:])
	case "FDEL_VALUE":
		p.handleFDelValue(parts[1:])
	case "FGET":
		p.handleFGet(parts[1:])
	case "LPUSH_FRONT":
		p.handleLPushFront(parts[1:])
	case "LPUSH_BACK":
		p.handleLPushBack(parts[1:])
	case "LINSERT_BEFORE":
		p.handleLInsertBefore(parts[1:])
	case "LINSERT_AFTER":
		p.handleLInsertAfter(parts[1:])
	case "LDEL_FRONT":
		p.handleLDelFront(parts[1:])
	case "LDEL_BACK":
		p.handleLDelBack(parts[1:])
	case "LDEL_VALUE":
		p.handleLDelValue(parts[1:])
	case "LGET":
		p.handleLGet(parts[1:])
	case "SPUSH":
		p.handleSPush(parts[1:])
	case "SPOP":
		p.handleSPop(parts[1:])
	case "SPEEK":
		p.handleSPeek(parts[1:])
	case "QPUSH":
		p.handleQPush(parts[1:])
	case "QPOP":
		p.handleQPop(parts[1:])
	case "QPEEK":
		p.handleQPeek(parts[1:])
	case "TINSERT":
		p.handleTInsert(parts[1:])
	case "TDEL":
		p.handleTDel(parts[1:])
	case "TGET":
		p.handleTGet(parts[1:])
	case "HINSERT":
		p.handleHInsert(parts[1:])
	case "HGET":
		p.handleHGet(parts[1:])
	case "HDEL":
		p.handleHDel(parts[1:])
	case "HSIZE":
		p.handleHSize(parts[1:])
	case "PRINT":
		p.handlePrint(parts[1:])
	case "SAVE":
		p.handleSave(parts[1:])
	case "LOAD":
		p.handleLoad(parts[1:])
	case "SAVE_TEXT":
		p.handleSaveText(parts[1:])
	case "SAVE_BINARY":
		p.handleSaveBinary(parts[1:])
	case "LOAD_TEXT":
		p.handleLoadText(parts[1:])
	case "LOAD_BINARY":
		p.handleLoadBinary(parts[1:])
	case "HELP":
		p.handleHelp()
	case "EXIT":
		fmt.Println("Выход из программы...")
	default:
		fmt.Println("Неизвестная команда. Введите HELP для списка команд.")
	}
}

func (p *CommandParser) handleCreate(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Ошибка: недостаточно параметров для CREATE")
		return
	}

	typeName := parts[0]
	name := parts[1]

	switch typeName {
	case "ARRAY":
		arr := NewArray(name)
		p.db.AddArray(arr)
		fmt.Printf("Массив '%s' создан.\n", name)
	case "SLL":
		sll := NewSinglyLinkedList(name)
		p.db.AddSLL(sll)
		fmt.Printf("Односвязный список '%s' создан.\n", name)
	case "DLL":
		dll := NewDoublyLinkedList(name)
		p.db.AddDLL(dll)
		fmt.Printf("Двусвязный список '%s' создан.\n", name)
	case "STACK":
		stack := NewStack(name)
		p.db.AddStack(stack)
		fmt.Printf("Стек '%s' создан.\n", name)
	case "QUEUE":
		queue := NewQueue(name)
		p.db.AddQueue(queue)
		fmt.Printf("Очередь '%s' создан.\n", name)
	case "TREE":
		tree := NewAVLTree(name)
		p.db.AddTree(tree)
		fmt.Printf("Дерево '%s' создан.\n", name)
	case "HASH":
		table := NewHashTable(name)
		p.db.AddHashTable(table)
		fmt.Printf("Хеш-таблица '%s' создан.\n", name)
	default:
		fmt.Println("Ошибка: неизвестный тип структуры")
	}
}

func (p *CommandParser) handleMPush(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	value := parts[1]

	arr := p.db.FindArray(name)
	if arr == nil {
		fmt.Println("FALSE")
		return
	}

	arr.PushBack(value)
	fmt.Println(value)
}

func (p *CommandParser) handleMInsert(parts []string) {
	if len(parts) < 3 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	index, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("FALSE")
		return
	}
	value := parts[2]

	arr := p.db.FindArray(name)
	if arr == nil {
		fmt.Println("FALSE")
		return
	}

	if err := arr.Insert(index, value); err != nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println(value)
}

func (p *CommandParser) handleMGet(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	index, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("FALSE")
		return
	}

	arr := p.db.FindArray(name)
	if arr == nil {
		fmt.Println("FALSE")
		return
	}

	value, err := arr.Get(index)
	if err != nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println(value)
}

func (p *CommandParser) handleMDel(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	index, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("FALSE")
		return
	}

	arr := p.db.FindArray(name)
	if arr == nil {
		fmt.Println("FALSE")
		return
	}

	if err := arr.Remove(index); err != nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println("TRUE")
}

func (p *CommandParser) handleMReplace(parts []string) {
	if len(parts) < 3 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	index, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("FALSE")
		return
	}
	value := parts[2]

	arr := p.db.FindArray(name)
	if arr == nil {
		fmt.Println("FALSE")
		return
	}

	if err := arr.Replace(index, value); err != nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println(value)
}

func (p *CommandParser) handleMLength(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]

	arr := p.db.FindArray(name)
	if arr == nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println(arr.Length())
}

func (p *CommandParser) handleFPushFront(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	value := parts[1]

	sll := p.db.FindSLL(name)
	if sll == nil {
		fmt.Println("FALSE")
		return
	}

	sll.PushFront(value)
	fmt.Println(value)
}

func (p *CommandParser) handleFPushBack(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	value := parts[1]

	sll := p.db.FindSLL(name)
	if sll == nil {
		fmt.Println("FALSE")
		return
	}

	sll.PushBack(value)
	fmt.Println(value)
}

func (p *CommandParser) handleFInsertBefore(parts []string) {
	if len(parts) < 3 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	target := parts[1]
	value := parts[2]

	sll := p.db.FindSLL(name)
	if sll == nil {
		fmt.Println("FALSE")
		return
	}

	sll.InsertBefore(target, value)
	fmt.Println(value)
}

func (p *CommandParser) handleFInsertAfter(parts []string) {
	if len(parts) < 3 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	target := parts[1]
	value := parts[2]

	sll := p.db.FindSLL(name)
	if sll == nil {
		fmt.Println("FALSE")
		return
	}

	sll.InsertAfter(target, value)
	fmt.Println(value)
}

func (p *CommandParser) handleFDelFront(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]

	sll := p.db.FindSLL(name)
	if sll == nil {
		fmt.Println("FALSE")
		return
	}

	sll.DeleteFront()
	fmt.Println("TRUE")
}

func (p *CommandParser) handleFDelBack(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]

	sll := p.db.FindSLL(name)
	if sll == nil {
		fmt.Println("FALSE")
		return
	}

	sll.DeleteBack()
	fmt.Println("TRUE")
}

func (p *CommandParser) handleFDelValue(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	value := parts[1]

	sll := p.db.FindSLL(name)
	if sll == nil {
		fmt.Println("FALSE")
		return
	}

	sll.DeleteByValue(value)
	fmt.Println("TRUE")
}

func (p *CommandParser) handleFGet(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	value := parts[1]

	sll := p.db.FindSLL(name)
	if sll == nil {
		fmt.Println("FALSE")
		return
	}

	node := sll.FindByValue(value)
	if node != nil {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}

func (p *CommandParser) handleLPushFront(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	value := parts[1]

	dll := p.db.FindDLL(name)
	if dll == nil {
		fmt.Println("FALSE")
		return
	}

	dll.PushFront(value)
	fmt.Println(value)
}

func (p *CommandParser) handleLPushBack(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	value := parts[1]

	dll := p.db.FindDLL(name)
	if dll == nil {
		fmt.Println("FALSE")
		return
	}

	dll.PushBack(value)
	fmt.Println(value)
}

func (p *CommandParser) handleLInsertBefore(parts []string) {
	if len(parts) < 3 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	target := parts[1]
	value := parts[2]

	dll := p.db.FindDLL(name)
	if dll == nil {
		fmt.Println("FALSE")
		return
	}

	dll.InsertBefore(target, value)
	fmt.Println(value)
}

func (p *CommandParser) handleLInsertAfter(parts []string) {
	if len(parts) < 3 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	target := parts[1]
	value := parts[2]

	dll := p.db.FindDLL(name)
	if dll == nil {
		fmt.Println("FALSE")
		return
	}

	dll.InsertAfter(target, value)
	fmt.Println(value)
}

func (p *CommandParser) handleLDelFront(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]

	dll := p.db.FindDLL(name)
	if dll == nil {
		fmt.Println("FALSE")
		return
	}

	dll.DeleteFront()
	fmt.Println("TRUE")
}

func (p *CommandParser) handleLDelBack(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]

	dll := p.db.FindDLL(name)
	if dll == nil {
		fmt.Println("FALSE")
		return
	}

	dll.DeleteBack()
	fmt.Println("TRUE")
}

func (p *CommandParser) handleLDelValue(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	value := parts[1]

	dll := p.db.FindDLL(name)
	if dll == nil {
		fmt.Println("FALSE")
		return
	}

	dll.DeleteByValue(value)
	fmt.Println("TRUE")
}

func (p *CommandParser) handleLGet(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	value := parts[1]

	dll := p.db.FindDLL(name)
	if dll == nil {
		fmt.Println("FALSE")
		return
	}

	node := dll.FindByValue(value)
	if node != nil {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}

func (p *CommandParser) handleSPush(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	value := parts[1]

	stack := p.db.FindStack(name)
	if stack == nil {
		fmt.Println("FALSE")
		return
	}

	stack.Push(value)
	fmt.Println(value)
}

func (p *CommandParser) handleSPop(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]

	stack := p.db.FindStack(name)
	if stack == nil {
		fmt.Println("FALSE")
		return
	}

	value, err := stack.Pop()
	if err != nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println(value)
}

func (p *CommandParser) handleSPeek(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]

	stack := p.db.FindStack(name)
	if stack == nil {
		fmt.Println("FALSE")
		return
	}

	value, err := stack.Peek()
	if err != nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println(value)
}

func (p *CommandParser) handleQPush(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	value := parts[1]

	queue := p.db.FindQueue(name)
	if queue == nil {
		fmt.Println("FALSE")
		return
	}

	queue.Push(value)
	fmt.Println(value)
}

func (p *CommandParser) handleQPop(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]

	queue := p.db.FindQueue(name)
	if queue == nil {
		fmt.Println("FALSE")
		return
	}

	value, err := queue.Pop()
	if err != nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println(value)
}

func (p *CommandParser) handleQPeek(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]

	queue := p.db.FindQueue(name)
	if queue == nil {
		fmt.Println("FALSE")
		return
	}

	value, err := queue.Peek()
	if err != nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println(value)
}

func (p *CommandParser) handleTInsert(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	valueStr := parts[1]

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("FALSE")
		return
	}

	tree := p.db.FindTree(name)
	if tree == nil {
		fmt.Println("FALSE")
		return
	}

	tree.Insert(value)
	fmt.Println(value)
}

func (p *CommandParser) handleTDel(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	valueStr := parts[1]

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("FALSE")
		return
	}

	tree := p.db.FindTree(name)
	if tree == nil {
		fmt.Println("FALSE")
		return
	}

	tree.Remove(value)
	fmt.Println("TRUE")
}

func (p *CommandParser) handleTGet(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	valueStr := parts[1]

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("FALSE")
		return
	}

	tree := p.db.FindTree(name)
	if tree == nil {
		fmt.Println("FALSE")
		return
	}

	node := tree.Search(value)
	if node != nil {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}

func (p *CommandParser) handleHInsert(parts []string) {
	if len(parts) < 3 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	key := parts[1]
	value := parts[2]

	table := p.db.FindHashTable(name)
	if table == nil {
		fmt.Println("FALSE")
		return
	}

	table.Insert(key, value)
	fmt.Println(value)
}

func (p *CommandParser) handleHGet(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	key := parts[1]

	table := p.db.FindHashTable(name)
	if table == nil {
		fmt.Println("FALSE")
		return
	}

	value, found := table.Search(key)
	if found {
		fmt.Println(value)
	} else {
		fmt.Println("FALSE")
	}
}

func (p *CommandParser) handleHDel(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]
	key := parts[1]

	table := p.db.FindHashTable(name)
	if table == nil {
		fmt.Println("FALSE")
		return
	}

	if table.Remove(key) {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}

func (p *CommandParser) handleHSize(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	name := parts[0]

	table := p.db.FindHashTable(name)
	if table == nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println(table.GetSize())
}

func (p *CommandParser) handlePrint(parts []string) {
	if len(parts) < 2 {
		fmt.Println("FALSE")
		return
	}

	typeName := parts[0]
	name := parts[1]

	switch typeName {
	case "ARRAY":
		arr := p.db.FindArray(name)
		if arr != nil {
			arr.Print()
		} else {
			fmt.Println("FALSE")
		}
	case "SLL":
		sll := p.db.FindSLL(name)
		if sll != nil {
			sll.Print()
		} else {
			fmt.Println("FALSE")
		}
	case "DLL":
		dll := p.db.FindDLL(name)
		if dll != nil {
			dll.PrintForward()
		} else {
			fmt.Println("FALSE")
		}
	case "STACK":
		stack := p.db.FindStack(name)
		if stack != nil {
			stack.Print()
		} else {
			fmt.Println("FALSE")
		}
	case "QUEUE":
		queue := p.db.FindQueue(name)
		if queue != nil {
			queue.Print()
		} else {
			fmt.Println("FALSE")
		}
	case "TREE":
		tree := p.db.FindTree(name)
		if tree != nil {
			tree.PrintInOrder()
		} else {
			fmt.Println("FALSE")
		}
	case "HASH":
		table := p.db.FindHashTable(name)
		if table != nil {
			table.Print()
		} else {
			fmt.Println("FALSE")
		}
	default:
		fmt.Println("FALSE")
	}
}

func (p *CommandParser) handleSave(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	filename := parts[0]
	if err := p.fileIO.SaveDatabaseToFile(p.db, filename); err != nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println("TRUE")
}

func (p *CommandParser) handleLoad(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	filename := parts[0]
	if err := p.fileIO.LoadDatabaseFromFile(p.db, filename); err != nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println("TRUE")
}

func (p *CommandParser) handleSaveText(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	filename := parts[0]
	serializer := NewSerializer()
	if err := serializer.SerializeDatabase(p.db, filename, TEXT); err != nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println("TRUE")
}

func (p *CommandParser) handleSaveBinary(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	filename := parts[0]
	serializer := NewSerializer()
	if err := serializer.SerializeDatabase(p.db, filename, BINARY); err != nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println("TRUE")
}

func (p *CommandParser) handleLoadText(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	filename := parts[0]
	serializer := NewSerializer()
	if err := serializer.DeserializeDatabase(p.db, filename, TEXT); err != nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println("TRUE")
}

func (p *CommandParser) handleLoadBinary(parts []string) {
	if len(parts) < 1 {
		fmt.Println("FALSE")
		return
	}

	filename := parts[0]
	serializer := NewSerializer()
	if err := serializer.DeserializeDatabase(p.db, filename, BINARY); err != nil {
		fmt.Println("FALSE")
		return
	}

	fmt.Println("TRUE")
}

func (p *CommandParser) handleHelp() {
	fmt.Println("=== Доступные команды ===")
	fmt.Println("CREATE ARRAY|SLL|DLL|STACK|QUEUE|TREE|HASH <name>")
	fmt.Println("MPUSH <name> <value> - Добавить в массив")
	fmt.Println("MINSERT <name> <index> <value> - Вставить в массив")
	fmt.Println("MDEL <name> <index> - Удалить из массива")
	fmt.Println("MGET <name> <index> - Получить из массива")
	fmt.Println("MREPLACE <name> <index> <value> - Заменить в массиве")
	fmt.Println("MLENGTH <name> - Длина массива")
	fmt.Println("FPUSH_FRONT <name> <value> - Добавить в начало SLL")
	fmt.Println("FPUSH_BACK <name> <value> - Добавить в конец SLL")
	fmt.Println("FINSERT_BEFORE <name> <target> <value> - Вставить перед в SLL")
	fmt.Println("FINSERT_AFTER <name> <target> <value> - Вставить после в SLL")
	fmt.Println("FDEL_FRONT <name> - Удалить из начала SLL")
	fmt.Println("FDEL_BACK <name> - Удалить с конца SLL")
	fmt.Println("FDEL_VALUE <name> <value> - Удалить по значению в SLL")
	fmt.Println("FGET <name> <value> - Поиск в SLL")
	fmt.Println("LPUSH_FRONT <name> <value> - Добавить в начало DLL")
	fmt.Println("LPUSH_BACK <name> <value> - Добавить в конец DLL")
	fmt.Println("LINSERT_BEFORE <name> <target> <value> - Вставить перед в DLL")
	fmt.Println("LINSERT_AFTER <name> <target> <value> - Вставить после в DLL")
	fmt.Println("LDEL_FRONT <name> - Удалить из начала DLL")
	fmt.Println("LDEL_BACK <name> - Удалить с конца DLL")
	fmt.Println("LDEL_VALUE <name> <value> - Удалить по значению в DLL")
	fmt.Println("LGET <name> <value> - Поиск в DLL")
	fmt.Println("SPUSH <name> <value> - Добавить в стек")
	fmt.Println("SPOP <name> - Извлечь из стека")
	fmt.Println("SPEEK <name> - Посмотреть вершину стека")
	fmt.Println("QPUSH <name> <value> - Добавить в очередь")
	fmt.Println("QPOP <name> - Извлечь из очереди")
	fmt.Println("QPEEK <name> - Посмотреть начало очереди")
	fmt.Println("TINSERT <name> <value> - Добавить в дерево")
	fmt.Println("TDEL <name> <value> - Удалить из дерева")
	fmt.Println("TGET <name> <value> - Поиск в дереве")
	fmt.Println("HINSERT <name> <key> <value> - Вставить в хеш-таблицу")
	fmt.Println("HGET <name> <key> - Получить из хеш-таблицы")
	fmt.Println("HDEL <name> <key> - Удалить из хеш-таблицы")
	fmt.Println("HSIZE <name> - Размер хеш-таблицы")
	fmt.Println("PRINT <type> <name> - Вывести структуру")
	fmt.Println("SAVE_TEXT <filename> - Сохранить базу в текстовом формате")
	fmt.Println("SAVE_BINARY <filename> - Сохранить базу в бинарном формате")
	fmt.Println("LOAD_TEXT <filename> - Загрузить базу из текстового формата")
	fmt.Println("LOAD_BINARY <filename> - Загрузить базу из бинарного формата")
	fmt.Println("SAVE <filename> - Сохранить базу (старый формат)")
	fmt.Println("LOAD <filename> - Загрузить базу (старый формат)")
	fmt.Println("HELP - Справка")
	fmt.Println("EXIT - Выход")
	fmt.Println("==========================")
}
