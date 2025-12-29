package dbmsgo

import (
	"bytes"
	"os"
	"testing"
        "github.com/stretchr/testify/assert"
)

func TestArrayComprehensive(t *testing.T) {
	arr := NewArray("test_array")

	// Test basic operations
	assert.True(t, arr.IsEmpty())
	assert.Equal(t, 0, arr.Length())

	// Test PushBack
	arr.PushBack("value1")
	arr.PushBack("value2")
	assert.Equal(t, 2, arr.Length())
	assert.False(t, arr.IsEmpty())

	// Test Get
	val, err := arr.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, "value1", val)

	val, err = arr.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, "value2", val)

	// Test Insert
	err = arr.Insert(1, "middle")
	assert.NoError(t, err)
	assert.Equal(t, 3, arr.Length())

	val, err = arr.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, "middle", val)

	// Test Replace
	err = arr.Replace(0, "new_value")
	assert.NoError(t, err)
	val, err = arr.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, "new_value", val)

	// Test Remove
	err = arr.Remove(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, arr.Length())

	// Test error cases
	_, err = arr.Get(10)
	assert.Error(t, err)

	err = arr.Insert(10, "value")
	assert.Error(t, err)

	err = arr.Remove(10)
	assert.Error(t, err)

	err = arr.Replace(10, "value")
	assert.Error(t, err)

	// Test Print (should not panic)
	arr.Print()

	// Test Cleanup
	arr.Cleanup()
	assert.True(t, arr.IsEmpty())
	assert.Equal(t, 0, arr.Length())
}

func TestSinglyLinkedListComprehensive(t *testing.T) {
	sll := NewSinglyLinkedList("test_sll")

	assert.True(t, sll.IsEmpty())
	assert.Nil(t, sll.GetHead())
	assert.Nil(t, sll.GetTail())

	// Test PushFront
	sll.PushFront("front")
	assert.Equal(t, "front", sll.GetHead().Data)
	assert.Equal(t, "front", sll.GetTail().Data)

	// Test PushBack
	sll.PushBack("back")
	assert.Equal(t, "front", sll.GetHead().Data)
	assert.Equal(t, "back", sll.GetTail().Data)

	// Test InsertBefore
	sll.InsertBefore("back", "middle")
	current := sll.GetHead()
	assert.Equal(t, "front", current.Data)
	assert.Equal(t, "middle", current.Next.Data)
	assert.Equal(t, "back", current.Next.Next.Data)

	// Test InsertAfter
	sll.InsertAfter("front", "after_front")
	current = sll.GetHead()
	assert.Equal(t, "front", current.Data)
	assert.Equal(t, "after_front", current.Next.Data)

	// Test FindByValue
	node := sll.FindByValue("middle")
	assert.NotNil(t, node)
	assert.Equal(t, "middle", node.Data)

	node = sll.FindByValue("nonexistent")
	assert.Nil(t, node)

	// Test DeleteFront
	sll.DeleteFront()
	assert.Equal(t, "after_front", sll.GetHead().Data)

	// Test DeleteBack
	sll.DeleteBack()
	assert.Equal(t, "middle", sll.GetTail().Data)

	// Test DeleteByValue
	sll.PushBack("to_delete")
	sll.PushBack("to_delete") // multiple occurrences
	sll.DeleteByValue("to_delete")
	assert.Nil(t, sll.FindByValue("to_delete"))

	// Test Print
	sll.Print()

	// Test Cleanup
	sll.Cleanup()
	assert.True(t, sll.IsEmpty())
}

func TestDoublyLinkedListComprehensive(t *testing.T) {
	dll := NewDoublyLinkedList("test_dll")

	assert.True(t, dll.IsEmpty())

	// Test PushFront and PushBack
	dll.PushFront("front")
	dll.PushBack("back")
	assert.Equal(t, "front", dll.GetHead().Data)
	assert.Equal(t, "back", dll.GetTail().Data)

	// Test InsertBefore and InsertAfter
	dll.InsertBefore("back", "before_back")
	dll.InsertAfter("front", "after_front")

	assert.Equal(t, "before_back", dll.GetTail().Prev.Data)
	assert.Equal(t, "after_front", dll.GetHead().Next.Data)

	// Test navigation
	head := dll.GetHead()
	assert.Equal(t, "front", head.Data)
	assert.Equal(t, "after_front", head.Next.Data)
	assert.Nil(t, head.Prev)

	tail := dll.GetTail()
	assert.Equal(t, "back", tail.Data)
	assert.Equal(t, "before_back", tail.Prev.Data)
	assert.Nil(t, tail.Next)

	// Test FindByValue
	node := dll.FindByValue("after_front")
	assert.NotNil(t, node)
	assert.Equal(t, "after_front", node.Data)

	// Test deletion operations
	dll.DeleteFront()
	assert.Equal(t, "after_front", dll.GetHead().Data)

	dll.DeleteBack()
	assert.Equal(t, "before_back", dll.GetTail().Data)

	dll.DeleteByValue("before_back")
	assert.Nil(t, dll.FindByValue("before_back"))

	// Test Print functions
	dll.PrintForward()
	dll.PrintBackward()

	// Test Cleanup
	dll.Cleanup()
	assert.True(t, dll.IsEmpty())
}

func TestStackComprehensive(t *testing.T) {
	stack := NewStack("test_stack")

	assert.True(t, stack.IsEmpty())
	assert.Equal(t, 0, stack.GetSize())

	// Test Push and Peek
	stack.Push("value1")
	assert.False(t, stack.IsEmpty())
	assert.Equal(t, 1, stack.GetSize())

	peekVal, err := stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "value1", peekVal)

	// Test multiple pushes
	stack.Push("value2")
	stack.Push("value3")
	assert.Equal(t, 3, stack.GetSize())

	peekVal, err = stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "value3", peekVal)

	// Test Pop
	popVal, err := stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "value3", popVal)
	assert.Equal(t, 2, stack.GetSize())

	// Test LIFO order
	popVal, err = stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "value2", popVal)

	popVal, err = stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "value1", popVal)

	assert.True(t, stack.IsEmpty())

	// Test error cases
	_, err = stack.Pop()
	assert.Error(t, err)

	_, err = stack.Peek()
	assert.Error(t, err)

	// Test Print
	stack.Push("test")
	stack.Print()

	// Test Cleanup
	stack.Cleanup()
	assert.True(t, stack.IsEmpty())
}

func TestQueueComprehensive(t *testing.T) {
	queue := NewQueue("test_queue")

	assert.True(t, queue.IsEmpty())
	assert.Equal(t, 0, queue.GetSize())

	// Test Push and Peek
	queue.Push("value1")
	assert.False(t, queue.IsEmpty())
	assert.Equal(t, 1, queue.GetSize())

	peekVal, err := queue.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "value1", peekVal)

	// Test multiple pushes
	queue.Push("value2")
	queue.Push("value3")
	assert.Equal(t, 3, queue.GetSize())

	peekVal, err = queue.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "value1", peekVal) // FIFO - first should remain

	// Test Pop
	popVal, err := queue.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "value1", popVal)
	assert.Equal(t, 2, queue.GetSize())

	// Test FIFO order
	popVal, err = queue.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "value2", popVal)

	popVal, err = queue.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "value3", popVal)

	assert.True(t, queue.IsEmpty())

	// Test error cases
	_, err = queue.Pop()
	assert.Error(t, err)

	_, err = queue.Peek()
	assert.Error(t, err)

	// Test Print
	queue.Push("test")
	queue.Print()

	// Test Cleanup
	queue.Cleanup()
	assert.True(t, queue.IsEmpty())
}

func TestAVLTreeComprehensive(t *testing.T) {
	tree := NewAVLTree("test_tree")

	assert.True(t, tree.IsEmpty())
	assert.Equal(t, 0, tree.CountElements())

	// Test Insert
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)

	assert.Equal(t, 7, tree.CountElements())

	// Test Search
	node := tree.Search(50)
	assert.NotNil(t, node)
	assert.Equal(t, 50, node.Data)

	node = tree.Search(100)
	assert.Nil(t, node)

	// Test Remove
	tree.Remove(20)
	assert.Equal(t, 6, tree.CountElements())
	assert.Nil(t, tree.Search(20))

	tree.Remove(30)
	assert.Equal(t, 5, tree.CountElements())

	tree.Remove(50) // root
	assert.Equal(t, 4, tree.CountElements())

	// Test SaveTree
	values := tree.SaveTree()
	assert.Greater(t, len(values), 0)

	// Test Print (should not panic)
	tree.PrintInOrder()

	// Test Cleanup
	tree.Cleanup()
	assert.True(t, tree.IsEmpty())
}

func TestHashTableComprehensive(t *testing.T) {
	table := NewHashTable("test_hash")

	assert.True(t, table.IsEmpty())
	assert.Equal(t, 0, table.GetSize())

	// Test Insert
	table.Insert("key1", "value1")
	table.Insert("key2", "value2")
	table.Insert("key3", "value3")

	assert.Equal(t, 3, table.GetSize())
	assert.False(t, table.IsEmpty())

	// Test Search
	val, found := table.Search("key1")
	assert.True(t, found)
	assert.Equal(t, "value1", val)

	val, found = table.Search("key2")
	assert.True(t, found)
	assert.Equal(t, "value2", val)

	// Test update existing key
	table.Insert("key1", "updated_value")
	val, found = table.Search("key1")
	assert.True(t, found)
	assert.Equal(t, "updated_value", val)

	// Test Remove
	removed := table.Remove("key1")
	assert.True(t, removed)
	assert.Equal(t, 2, table.GetSize())

	removed = table.Remove("nonexistent")
	assert.False(t, removed)

	// Test Print (should not panic)
	table.Print()

	// Test Cleanup
	table.Cleanup()
	assert.True(t, table.IsEmpty())
	assert.Equal(t, 0, table.GetSize())
}

func TestDatabaseComprehensive(t *testing.T) {
	db := NewDatabase()

	// Test adding and finding all structure types
	arr := NewArray("test_array")
	db.AddArray(arr)
	assert.Equal(t, arr, db.FindArray("test_array"))

	sll := NewSinglyLinkedList("test_sll")
	db.AddSLL(sll)
	assert.Equal(t, sll, db.FindSLL("test_sll"))

	dll := NewDoublyLinkedList("test_dll")
	db.AddDLL(dll)
	assert.Equal(t, dll, db.FindDLL("test_dll"))

	stack := NewStack("test_stack")
	db.AddStack(stack)
	assert.Equal(t, stack, db.FindStack("test_stack"))

	queue := NewQueue("test_queue")
	db.AddQueue(queue)
	assert.Equal(t, queue, db.FindQueue("test_queue"))

	tree := NewAVLTree("test_tree")
	db.AddTree(tree)
	assert.Equal(t, tree, db.FindTree("test_tree"))

	hash := NewHashTable("test_hash")
	db.AddHashTable(hash)
	assert.Equal(t, hash, db.FindHashTable("test_hash"))

	// Test finding non-existent structures
	assert.Nil(t, db.FindArray("nonexistent"))
	assert.Nil(t, db.FindSLL("nonexistent"))
	assert.Nil(t, db.FindStack("nonexistent"))

	// Test Cleanup
	db.Cleanup()
	assert.Nil(t, db.FindArray("test_array"))
	assert.Nil(t, db.FindSLL("test_sll"))
	assert.Nil(t, db.FindStack("test_stack"))
}

func TestFileIOComprehensive(t *testing.T) {
	db := NewDatabase()
	fileIO := NewFileIO()

	// Create test data
	arr := NewArray("test_array")
	arr.PushBack("arr_value1")
	arr.PushBack("arr_value2")
	db.AddArray(arr)

	sll := NewSinglyLinkedList("test_sll")
	sll.PushBack("sll_value1")
	sll.PushBack("sll_value2")
	db.AddSLL(sll)

	stack := NewStack("test_stack")
	stack.Push("stack_value1")
	stack.Push("stack_value2")
	db.AddStack(stack)

	queue := NewQueue("test_queue")
	queue.Push("queue_value1")
	queue.Push("queue_value2")
	db.AddQueue(queue)

	tree := NewAVLTree("test_tree")
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	db.AddTree(tree)

	hash := NewHashTable("test_hash")
	hash.Insert("key1", "value1")
	hash.Insert("key2", "value2")
	db.AddHashTable(hash)

	// Test SaveDatabaseToFile
	err := fileIO.SaveDatabaseToFile(db, "test_save.txt")
	assert.NoError(t, err)

	// Test LoadDatabaseFromFile
	newDB := NewDatabase()
	err = fileIO.LoadDatabaseFromFile(newDB, "test_save.txt")
	assert.NoError(t, err)

	// Verify loaded data
	loadedArr := newDB.FindArray("test_array")
	assert.NotNil(t, loadedArr)
	assert.Equal(t, 2, loadedArr.Length())

	loadedSLL := newDB.FindSLL("test_sll")
	assert.NotNil(t, loadedSLL)
	assert.False(t, loadedSLL.IsEmpty())

	// Cleanup
	os.Remove("test_save.txt")
}

func TestSerializerComprehensive(t *testing.T) {
	serializer := NewSerializer()

	// Test Array serialization
	arr := NewArray("test_array")
	arr.PushBack("value1")
	arr.PushBack("value2")

	var buf bytes.Buffer
	err := serializer.SerializeArray(arr, &buf, TEXT)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf.String())

	// Test SLL serialization
	sll := NewSinglyLinkedList("test_sll")
	sll.PushBack("sll_value1")
	sll.PushBack("sll_value2")

	buf.Reset()
	err = serializer.SerializeSLL(sll, &buf, TEXT)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf.String())

	// Test Stack serialization
	stack := NewStack("test_stack")
	stack.Push("stack_value1")
	stack.Push("stack_value2")

	buf.Reset()
	err = serializer.SerializeStack(stack, &buf, TEXT)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf.String())

	// Test Queue serialization
	queue := NewQueue("test_queue")
	queue.Push("queue_value1")
	queue.Push("queue_value2")

	buf.Reset()
	err = serializer.SerializeQueue(queue, &buf, TEXT)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf.String())

	// Test Tree serialization
	tree := NewAVLTree("test_tree")
	tree.Insert(10)
	tree.Insert(5)

	buf.Reset()
	err = serializer.SerializeTree(tree, &buf, TEXT)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf.String())

	// Test HashTable serialization
	hash := NewHashTable("test_hash")
	hash.Insert("key1", "value1")

	buf.Reset()
	err = serializer.SerializeHashTable(hash, &buf, TEXT)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf.String())

	// Test Database serialization (text format)
	db := NewDatabase()
	db.AddArray(arr)

	buf.Reset()
	err = serializer.SerializeDatabase(db, "test_serialize.txt", TEXT)
	assert.NoError(t, err)

	// Test Database deserialization
	err = serializer.DeserializeDatabase(db, "test_deserialize.txt", TEXT)
	assert.NoError(t, err)

	// Cleanup
	os.Remove("test_serialize.txt")
}

func TestIntegration(t *testing.T) {
	// Test complete workflow
	db := NewDatabase()
	parser := NewCommandParser(db)

	// Create multiple structures
	parser.ProcessCommand("CREATE ARRAY users")
	parser.ProcessCommand("CREATE SLL tasks")
	parser.ProcessCommand("CREATE STACK history")
	parser.ProcessCommand("CREATE QUEUE requests")
	parser.ProcessCommand("CREATE TREE scores")
	parser.ProcessCommand("CREATE HASH config")

	// Populate with data
	parser.ProcessCommand("MPUSH users Alice")
	parser.ProcessCommand("MPUSH users Bob")

	parser.ProcessCommand("FPUSH_BACK tasks 'Task 1'")
	parser.ProcessCommand("FPUSH_BACK tasks 'Task 2'")

	parser.ProcessCommand("SPUSH history 'Action 1'")
	parser.ProcessCommand("SPUSH history 'Action 2'")

	parser.ProcessCommand("QPUSH requests 'Req 1'")
	parser.ProcessCommand("QPUSH requests 'Req 2'")

	parser.ProcessCommand("TINSERT scores 100")
	parser.ProcessCommand("TINSERT scores 50")

	parser.ProcessCommand("HINSERT config timeout 30")
	parser.ProcessCommand("HINSERT config retries 3")

	// Verify data exists
	assert.NotNil(t, db.FindArray("users"))
	assert.NotNil(t, db.FindSLL("tasks"))
	assert.NotNil(t, db.FindStack("history"))
	assert.NotNil(t, db.FindQueue("requests"))
	assert.NotNil(t, db.FindTree("scores"))
	assert.NotNil(t, db.FindHashTable("config"))

	// Test file operations
	parser.ProcessCommand("SAVE integration_test.txt")

	// Load into new database
	newDB := NewDatabase()
	newParser := NewCommandParser(newDB)
	newParser.ProcessCommand("LOAD integration_test.txt")

	// Verify loaded structures
	assert.NotNil(t, newDB.FindArray("users"))
	assert.NotNil(t, newDB.FindSLL("tasks"))

	// Cleanup
	os.Remove("integration_test.txt")
}
