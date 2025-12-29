package dbmsgo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileIO_SaveAndLoadDatabase(t *testing.T) {
	fileIO := NewFileIO()
	db := NewDatabase()

	// Create test data
	arr := NewArray("test_array")
	arr.PushBack("value1")
	arr.PushBack("value2")
	db.AddArray(arr)

	sll := NewSinglyLinkedList("test_sll")
	sll.PushBack("sll_value1")
	sll.PushBack("sll_value2")
	db.AddSLL(sll)

	dll := NewDoublyLinkedList("test_dll")
	dll.PushBack("dll_value1")
	dll.PushBack("dll_value2")
	db.AddDLL(dll)

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
	filename := "test_save_load.txt"
	err := fileIO.SaveDatabaseToFile(db, filename)
	assert.NoError(t, err)

	// Verify file was created
	_, err = os.Stat(filename)
	assert.NoError(t, err)

	// Test LoadDatabaseFromFile
	newDB := NewDatabase()
	err = fileIO.LoadDatabaseFromFile(newDB, filename)
	assert.NoError(t, err)

	// Verify loaded data
	loadedArray := newDB.FindArray("test_array")
	assert.NotNil(t, loadedArray)
	assert.Equal(t, 2, loadedArray.Length())
	val, err := loadedArray.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, "value1", val)

	loadedSLL := newDB.FindSLL("test_sll")
	assert.NotNil(t, loadedSLL)
	assert.False(t, loadedSLL.IsEmpty())

	loadedDLL := newDB.FindDLL("test_dll")
	assert.NotNil(t, loadedDLL)
	assert.False(t, loadedDLL.IsEmpty())

	loadedStack := newDB.FindStack("test_stack")
	assert.NotNil(t, loadedStack)
	assert.Equal(t, 2, loadedStack.GetSize())

	loadedQueue := newDB.FindQueue("test_queue")
	assert.NotNil(t, loadedQueue)
	assert.Equal(t, 2, loadedQueue.GetSize())

	loadedTree := newDB.FindTree("test_tree")
	assert.NotNil(t, loadedTree)
	assert.True(t, loadedTree.CountElements() > 0)

	loadedHash := newDB.FindHashTable("test_hash")
	assert.NotNil(t, loadedHash)
	assert.Equal(t, 2, loadedHash.GetSize())

	// Cleanup
	os.Remove(filename)
}

func TestFileIO_SaveDatabaseToFile_Error(t *testing.T) {
	fileIO := NewFileIO()
	db := NewDatabase()

	// Test with invalid filename (directory)
	err := fileIO.SaveDatabaseToFile(db, "/invalid/path/file.txt")
	assert.Error(t, err)
}

func TestFileIO_LoadDatabaseFromFile_NotFound(t *testing.T) {
	fileIO := NewFileIO()
	db := NewDatabase()

	// Test loading non-existent file
	err := fileIO.LoadDatabaseFromFile(db, "nonexistent_file.txt")
	assert.Error(t, err)
}

func TestFileIO_LoadArray(t *testing.T) {
	fileIO := NewFileIO()
	db := NewDatabase()

	// Test valid array data
	parts := []string{"ARRAY", "test_array", "2", "value1", "value2"}
	fileIO.loadArray(db, parts)

	arr := db.FindArray("test_array")
	assert.NotNil(t, arr)
	assert.Equal(t, 2, arr.Length())
	val, err := arr.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, "value1", val)
}

func TestFileIO_LoadArray_InvalidData(t *testing.T) {
	fileIO := NewFileIO()
	db := NewDatabase()

	// Test insufficient parts
	parts := []string{"ARRAY", "test"}
	fileIO.loadArray(db, parts)
	assert.Nil(t, db.FindArray("test"))

	// Test invalid size
	parts = []string{"ARRAY", "test", "invalid"}
	fileIO.loadArray(db, parts)
	assert.Nil(t, db.FindArray("test"))

	// Test insufficient values
	parts = []string{"ARRAY", "test", "3", "value1"} // size 3 but only 1 value
	fileIO.loadArray(db, parts)
	assert.Nil(t, db.FindArray("test"))
}

func TestFileIO_LoadSLL(t *testing.T) {
	fileIO := NewFileIO()
	db := NewDatabase()

	parts := []string{"SLL", "test_sll", "2", "value1", "value2"}
	fileIO.loadSLL(db, parts)

	sll := db.FindSLL("test_sll")
	assert.NotNil(t, sll)
	assert.False(t, sll.IsEmpty())
	assert.Equal(t, "value1", sll.GetHead().Data)
	assert.Equal(t, "value2", sll.GetTail().Data)
}

func TestFileIO_LoadDLL(t *testing.T) {
	fileIO := NewFileIO()
	db := NewDatabase()

	parts := []string{"DLL", "test_dll", "2", "value1", "value2"}
	fileIO.loadDLL(db, parts)

	dll := db.FindDLL("test_dll")
	assert.NotNil(t, dll)
	assert.False(t, dll.IsEmpty())
	assert.Equal(t, "value1", dll.GetHead().Data)
	assert.Equal(t, "value2", dll.GetTail().Data)
}

func TestFileIO_LoadStack(t *testing.T) {
	fileIO := NewFileIO()
	db := NewDatabase()

	parts := []string{"STACK", "test_stack", "2", "value1", "value2"}
	fileIO.loadStack(db, parts)

	stack := db.FindStack("test_stack")
	assert.NotNil(t, stack)
	assert.Equal(t, 2, stack.GetSize())
	
	// Stack should have values in LIFO order (value2 on top)
	val, err := stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "value2", val)
}

func TestFileIO_LoadQueue(t *testing.T) {
	fileIO := NewFileIO()
	db := NewDatabase()

	parts := []string{"QUEUE", "test_queue", "2", "value1", "value2"}
	fileIO.loadQueue(db, parts)

	queue := db.FindQueue("test_queue")
	assert.NotNil(t, queue)
	assert.Equal(t, 2, queue.GetSize())
	
	// Queue should have values in FIFO order (value1 first)
	val, err := queue.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "value1", val)
}

func TestFileIO_LoadTree(t *testing.T) {
	fileIO := NewFileIO()
	db := NewDatabase()

	parts := []string{"TREE", "test_tree", "3", "10", "5", "15"}
	fileIO.loadTree(db, parts)

	tree := db.FindTree("test_tree")
	assert.NotNil(t, tree)
	assert.Equal(t, 3, tree.CountElements())
	
	// Verify values were inserted
	assert.NotNil(t, tree.Search(10))
	assert.NotNil(t, tree.Search(5))
	assert.NotNil(t, tree.Search(15))
}

func TestFileIO_LoadTree_InvalidNumbers(t *testing.T) {
	fileIO := NewFileIO()
	db := NewDatabase()

	// Test with invalid number
	parts := []string{"TREE", "test_tree", "2", "10", "invalid"}
	fileIO.loadTree(db, parts)

	tree := db.FindTree("test_tree")
	assert.NotNil(t, tree)
	// Should still insert valid numbers
	assert.NotNil(t, tree.Search(10))
}

func TestFileIO_LoadHashTable(t *testing.T) {
	fileIO := NewFileIO()
	db := NewDatabase()

	parts := []string{"HASH", "test_hash", "2", "key1", "value1", "key2", "value2"}
	fileIO.loadHashTable(db, parts)

	hash := db.FindHashTable("test_hash")
	assert.NotNil(t, hash)
	assert.Equal(t, 2, hash.GetSize())
	
	val, found := hash.Search("key1")
	assert.True(t, found)
	assert.Equal(t, "value1", val)
	
	val, found = hash.Search("key2")
	assert.True(t, found)
	assert.Equal(t, "value2", val)
}

func TestFileIO_LoadHashTable_InvalidData(t *testing.T) {
	fileIO := NewFileIO()
	db := NewDatabase()

	// Test insufficient key-value pairs
	parts := []string{"HASH", "test_hash", "2", "key1"} // missing value
	fileIO.loadHashTable(db, parts)

	hash := db.FindHashTable("test_hash")
	assert.Nil(t, hash) // Should not create hash table with invalid data
}

func TestFileIO_LoadFromFileWithMultipleStructures(t *testing.T) {
	fileIO := NewFileIO()
	
	// Create test file with multiple structures
	content := `ARRAY test_array 2 value1 value2
SLL test_sll 2 sll_value1 sll_value2
DLL test_dll 2 dll_value1 dll_value2
STACK test_stack 2 stack_value1 stack_value2
QUEUE test_queue 2 queue_value1 queue_value2
TREE test_tree 3 10 5 15
HASH test_hash 2 key1 value1 key2 value2`

	filename := "test_multiple.txt"
	err := os.WriteFile(filename, []byte(content), 0644)
	assert.NoError(t, err)
	defer os.Remove(filename)

	// Load from file
	db := NewDatabase()
	err = fileIO.LoadDatabaseFromFile(db, filename)
	assert.NoError(t, err)

	// Verify all structures were loaded
	assert.NotNil(t, db.FindArray("test_array"))
	assert.NotNil(t, db.FindSLL("test_sll"))
	assert.NotNil(t, db.FindDLL("test_dll"))
	assert.NotNil(t, db.FindStack("test_stack"))
	assert.NotNil(t, db.FindQueue("test_queue"))
	assert.NotNil(t, db.FindTree("test_tree"))
	assert.NotNil(t, db.FindHashTable("test_hash"))
}

func TestFileIO_LoadFromFileWithInvalidLines(t *testing.T) {
	fileIO := NewFileIO()
	
	// Create test file with invalid and valid lines
	content := `INVALID line
ARRAY test_array 1 value1
ANOTHER invalid
SLL test_sll 1 sll_value1`

	filename := "test_mixed.txt"
	err := os.WriteFile(filename, []byte(content), 0644)
	assert.NoError(t, err)
	defer os.Remove(filename)

	// Load from file
	db := NewDatabase()
	err = fileIO.LoadDatabaseFromFile(db, filename)
	assert.NoError(t, err)

	// Verify only valid structures were loaded
	assert.NotNil(t, db.FindArray("test_array"))
	assert.NotNil(t, db.FindSLL("test_sll"))
}

func TestNewFileIO(t *testing.T) {
	fileIO := NewFileIO()
	assert.NotNil(t, fileIO)
	assert.NotNil(t, fileIO.serializer)
}
