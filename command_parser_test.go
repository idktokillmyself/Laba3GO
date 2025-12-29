package dbmsgo

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandParser_CreateCommands(t *testing.T) {
	db := NewDatabase()
	parser := NewCommandParser(db)

	// Test CREATE commands
	parser.ProcessCommand("CREATE ARRAY test_array")
	assert.NotNil(t, db.FindArray("test_array"))

	parser.ProcessCommand("CREATE SLL test_sll")
	assert.NotNil(t, db.FindSLL("test_sll"))

	parser.ProcessCommand("CREATE DLL test_dll")
	assert.NotNil(t, db.FindDLL("test_dll"))

	parser.ProcessCommand("CREATE STACK test_stack")
	assert.NotNil(t, db.FindStack("test_stack"))

	parser.ProcessCommand("CREATE QUEUE test_queue")
	assert.NotNil(t, db.FindQueue("test_queue"))

	parser.ProcessCommand("CREATE TREE test_tree")
	assert.NotNil(t, db.FindTree("test_tree"))

	parser.ProcessCommand("CREATE HASH test_hash")
	assert.NotNil(t, db.FindHashTable("test_hash"))

	// Test invalid CREATE commands
	parser.ProcessCommand("CREATE INVALID test_invalid")
	parser.ProcessCommand("CREATE") // insufficient parameters
}

func TestCommandParser_ArrayOperations(t *testing.T) {
	db := NewDatabase()
	parser := NewCommandParser(db)

	// Setup
	parser.ProcessCommand("CREATE ARRAY test_array")

	// Test MPUSH
	parser.ProcessCommand("MPUSH test_array value1")
	arr := db.FindArray("test_array")
	assert.Equal(t, 1, arr.Length())

	// Test MINSERT
	parser.ProcessCommand("MINSERT test_array 0 value0")
	assert.Equal(t, 2, arr.Length())

	// Test MGET
	// Capture output for MGET
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	parser.ProcessCommand("MGET test_array 0")

	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := strings.TrimSpace(buf.String())
	assert.Equal(t, "value0", output)

	// Test MLENGTH
	buf.Reset()
	os.Stdout = w
	parser.ProcessCommand("MLENGTH test_array")
	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output = strings.TrimSpace(buf.String())
	assert.Equal(t, "2", output)

	// Test MDEL
	parser.ProcessCommand("MDEL test_array 0")
	assert.Equal(t, 1, arr.Length())

	// Test MREPLACE
	parser.ProcessCommand("MREPLACE test_array 0 new_value")
	value, _ := arr.Get(0)
	assert.Equal(t, "new_value", value)
}

func TestCommandParser_SLLOperations(t *testing.T) {
	db := NewDatabase()
	parser := NewCommandParser(db)

	parser.ProcessCommand("CREATE SLL test_sll")

	// Test FPUSH_FRONT
	parser.ProcessCommand("FPUSH_FRONT test_sll front_value")
	sll := db.FindSLL("test_sll")
	assert.False(t, sll.IsEmpty())

	// Test FPUSH_BACK
	parser.ProcessCommand("FPUSH_BACK test_sll back_value")
	assert.Equal(t, "back_value", sll.GetTail().Data)

	// Test FINSERT_BEFORE and FINSERT_AFTER
	parser.ProcessCommand("FINSERT_BEFORE back_value middle_before")
	parser.ProcessCommand("FINSERT_AFTER front_value middle_after")

	// Test FGET
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	parser.ProcessCommand("FGET test_sll front_value")
	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := strings.TrimSpace(buf.String())
	assert.Equal(t, "TRUE", output)

	// Test FDEL operations
	parser.ProcessCommand("FDEL_FRONT test_sll")
	parser.ProcessCommand("FDEL_BACK test_sll")
	parser.ProcessCommand("FDEL_VALUE test_sll middle_before")
}

func TestCommandParser_DLLOperations(t *testing.T) {
	db := NewDatabase()
	parser := NewCommandParser(db)

	parser.ProcessCommand("CREATE DLL test_dll")

	// Test LPUSH operations
	parser.ProcessCommand("LPUSH_FRONT test_dll front_value")
	parser.ProcessCommand("LPUSH_BACK test_dll back_value")

	dll := db.FindDLL("test_dll")
	assert.Equal(t, "front_value", dll.GetHead().Data)
	assert.Equal(t, "back_value", dll.GetTail().Data)

	// Test LINSERT operations
	parser.ProcessCommand("LINSERT_BEFORE test_dll back_value before_value")
	parser.ProcessCommand("LINSERT_AFTER test_dll front_value after_value")

	// Test LGET
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	parser.ProcessCommand("LGET test_dll front_value")
	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := strings.TrimSpace(buf.String())
	assert.Equal(t, "TRUE", output)

	// Test LDEL operations
	parser.ProcessCommand("LDEL_FRONT test_dll")
	parser.ProcessCommand("LDEL_BACK test_dll")
	parser.ProcessCommand("LDEL_VALUE test_dll before_value")
}

func TestCommandParser_StackOperations(t *testing.T) {
	db := NewDatabase()
	parser := NewCommandParser(db)

	parser.ProcessCommand("CREATE STACK test_stack")

	// Test SPUSH
	parser.ProcessCommand("SPUSH test_stack value1")
	parser.ProcessCommand("SPUSH test_stack value2")

	stack := db.FindStack("test_stack")
	assert.Equal(t, 2, stack.GetSize())

	// Test SPEEK and SPOP
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	parser.ProcessCommand("SPEEK test_stack")
	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := strings.TrimSpace(buf.String())
	assert.Equal(t, "value2", output)

	buf.Reset()
	os.Stdout = w
	parser.ProcessCommand("SPOP test_stack")
	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output = strings.TrimSpace(buf.String())
	assert.Equal(t, "value2", output)
}

func TestCommandParser_QueueOperations(t *testing.T) {
	db := NewDatabase()
	parser := NewCommandParser(db)

	parser.ProcessCommand("CREATE QUEUE test_queue")

	// Test QPUSH
	parser.ProcessCommand("QPUSH test_queue value1")
	parser.ProcessCommand("QPUSH test_queue value2")

	queue := db.FindQueue("test_queue")
	assert.Equal(t, 2, queue.GetSize())

	// Test QPEEK and QPOP
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	parser.ProcessCommand("QPEEK test_queue")
	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := strings.TrimSpace(buf.String())
	assert.Equal(t, "value1", output)

	buf.Reset()
	os.Stdout = w
	parser.ProcessCommand("QPOP test_queue")
	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output = strings.TrimSpace(buf.String())
	assert.Equal(t, "value1", output)
}

func TestCommandParser_TreeOperations(t *testing.T) {
	db := NewDatabase()
	parser := NewCommandParser(db)

	parser.ProcessCommand("CREATE TREE test_tree")

	// Test TINSERT
	parser.ProcessCommand("TINSERT test_tree 10")
	parser.ProcessCommand("TINSERT test_tree 5")
	parser.ProcessCommand("TINSERT test_tree 15")

	tree := db.FindTree("test_tree")
	assert.Equal(t, 3, tree.CountElements())

	// Test TGET
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	parser.ProcessCommand("TGET test_tree 10")
	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := strings.TrimSpace(buf.String())
	assert.Equal(t, "TRUE", output)

	// Test TDEL
	parser.ProcessCommand("TDEL test_tree 5")
	assert.Equal(t, 2, tree.CountElements())
}

func TestCommandParser_HashTableOperations(t *testing.T) {
	db := NewDatabase()
	parser := NewCommandParser(db)

	parser.ProcessCommand("CREATE HASH test_hash")

	// Test HINSERT
	parser.ProcessCommand("HINSERT test_hash key1 value1")
	parser.ProcessCommand("HINSERT test_hash key2 value2")

	table := db.FindHashTable("test_hash")
	assert.Equal(t, 2, table.GetSize())

	// Test HGET
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	parser.ProcessCommand("HGET test_hash key1")
	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := strings.TrimSpace(buf.String())
	assert.Equal(t, "value1", output)

	// Test HSIZE
	buf.Reset()
	os.Stdout = w
	parser.ProcessCommand("HSIZE test_hash")
	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output = strings.TrimSpace(buf.String())
	assert.Equal(t, "2", output)

	// Test HDEL
	parser.ProcessCommand("HDEL test_hash key1")
	assert.Equal(t, 1, table.GetSize())
}

func TestCommandParser_PrintCommands(t *testing.T) {
	db := NewDatabase()
	parser := NewCommandParser(db)

	// Create test structures
	parser.ProcessCommand("CREATE ARRAY test_array")
	parser.ProcessCommand("MPUSH test_array value1")

	parser.ProcessCommand("CREATE SLL test_sll")
	parser.ProcessCommand("FPUSH_BACK test_sll sll_value")

	parser.ProcessCommand("CREATE DLL test_dll")
	parser.ProcessCommand("LPUSH_BACK test_dll dll_value")

	parser.ProcessCommand("CREATE STACK test_stack")
	parser.ProcessCommand("SPUSH test_stack stack_value")

	parser.ProcessCommand("CREATE QUEUE test_queue")
	parser.ProcessCommand("QPUSH test_queue queue_value")

	parser.ProcessCommand("CREATE TREE test_tree")
	parser.ProcessCommand("TINSERT test_tree 42")

	parser.ProcessCommand("CREATE HASH test_hash")
	parser.ProcessCommand("HINSERT test_hash hash_key hash_value")

	// Test PRINT commands (they should not panic)
	parser.ProcessCommand("PRINT ARRAY test_array")
	parser.ProcessCommand("PRINT SLL test_sll")
	parser.ProcessCommand("PRINT DLL test_dll")
	parser.ProcessCommand("PRINT STACK test_stack")
	parser.ProcessCommand("PRINT QUEUE test_queue")
	parser.ProcessCommand("PRINT TREE test_tree")
	parser.ProcessCommand("PRINT HASH test_hash")
}

func TestCommandParser_FileOperations(t *testing.T) {
	db := NewDatabase()
	parser := NewCommandParser(db)

	// Create test data
	parser.ProcessCommand("CREATE ARRAY test_array")
	parser.ProcessCommand("MPUSH test_array value1")

	// Test SAVE and LOAD commands
	parser.ProcessCommand("SAVE test_db.txt")
	parser.ProcessCommand("SAVE_TEXT test_text.txt")
	parser.ProcessCommand("SAVE_BINARY test_binary.bin")

	// Test LOAD commands (they should not panic)
	parser.ProcessCommand("LOAD test_db.txt")
	parser.ProcessCommand("LOAD_TEXT test_text.txt")
	parser.ProcessCommand("LOAD_BINARY test_binary.bin")

	// Cleanup test files
	os.Remove("test_db.txt")
	os.Remove("test_text.txt")
	os.Remove("test_binary.bin")
}

func TestCommandParser_ErrorCases(t *testing.T) {
	db := NewDatabase()
	parser := NewCommandParser(db)

	// Test operations on non-existent structures
	parser.ProcessCommand("MPUSH nonexistent value")
	parser.ProcessCommand("MGET nonexistent 0")
	parser.ProcessCommand("FPUSH_FRONT nonexistent value")
	parser.ProcessCommand("SPOP nonexistent")

	// Test invalid parameters
	parser.ProcessCommand("CREATE") // no parameters
	parser.ProcessCommand("MPUSH")  // insufficient parameters
	parser.ProcessCommand("MGET array") // missing index

	// Test HELP and unknown commands
	parser.ProcessCommand("HELP")
	parser.ProcessCommand("UNKNOWN_COMMAND")
}

func TestCommandParser_EdgeCases(t *testing.T) {
	db := NewDatabase()
	parser := NewCommandParser(db)

	// Test empty command
	parser.ProcessCommand("")
	parser.ProcessCommand("   ")

	// Test with extra spaces
	parser.ProcessCommand("  CREATE   ARRAY   spaced_array  ")
	assert.NotNil(t, db.FindArray("spaced_array"))

	// Test case sensitivity
	parser.ProcessCommand("create array lower_array")
	assert.NotNil(t, db.FindArray("lower_array"))
}
