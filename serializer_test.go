package dbmsgo

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSerializer(t *testing.T) {
	serializer := NewSerializer()
	assert.NotNil(t, serializer)
}

func TestSerializer_ArrayTextFormat(t *testing.T) {
	serializer := NewSerializer()
	arr := NewArray("test_array")
	arr.PushBack("value1")
	arr.PushBack("value2")

	var buf bytes.Buffer
	err := serializer.SerializeArray(arr, &buf, TEXT)
	assert.NoError(t, err)

	output := buf.String()
	assert.Contains(t, output, "ARRAY test_array 2")
	assert.Contains(t, output, "value1")
	assert.Contains(t, output, "value2")
	assert.True(t, strings.HasSuffix(output, "\n"))
}

func TestSerializer_ArrayBinaryFormat(t *testing.T) {
	serializer := NewSerializer()
	arr := NewArray("test_array")
	arr.PushBack("value1")
	arr.PushBack("value2")

	var buf bytes.Buffer
	err := serializer.SerializeArray(arr, &buf, BINARY)
	assert.NoError(t, err)
	assert.Greater(t, buf.Len(), 0)
}

func TestSerializer_SLLTextFormat(t *testing.T) {
	serializer := NewSerializer()
	sll := NewSinglyLinkedList("test_sll")
	sll.PushBack("value1")
	sll.PushBack("value2")

	var buf bytes.Buffer
	err := serializer.SerializeSLL(sll, &buf, TEXT)
	assert.NoError(t, err)

	output := buf.String()
	assert.Contains(t, output, "SLL test_sll 2")
	assert.Contains(t, output, "value1")
	assert.Contains(t, output, "value2")
}

func TestSerializer_DLLTextFormat(t *testing.T) {
	serializer := NewSerializer()
	dll := NewDoublyLinkedList("test_dll")
	dll.PushBack("value1")
	dll.PushBack("value2")

	var buf bytes.Buffer
	err := serializer.SerializeDLL(dll, &buf, TEXT)
	assert.NoError(t, err)

	output := buf.String()
	assert.Contains(t, output, "DLL test_dll 2")
	assert.Contains(t, output, "value1")
	assert.Contains(t, output, "value2")
}

func TestSerializer_StackTextFormat(t *testing.T) {
	serializer := NewSerializer()
	stack := NewStack("test_stack")
	stack.Push("value1")
	stack.Push("value2")

	var buf bytes.Buffer
	err := serializer.SerializeStack(stack, &buf, TEXT)
	assert.NoError(t, err)

	output := buf.String()
	assert.Contains(t, output, "STACK test_stack 2")
	// Stack serializes in reverse order for LIFO
	assert.True(t, strings.Index(output, "value2") < strings.Index(output, "value1"))
}

func TestSerializer_QueueTextFormat(t *testing.T) {
	serializer := NewSerializer()
	queue := NewQueue("test_queue")
	queue.Push("value1")
	queue.Push("value2")

	var buf bytes.Buffer
	err := serializer.SerializeQueue(queue, &buf, TEXT)
	assert.NoError(t, err)

	output := buf.String()
	assert.Contains(t, output, "QUEUE test_queue 2")
	assert.Contains(t, output, "value1")
	assert.Contains(t, output, "value2")
}

func TestSerializer_TreeTextFormat(t *testing.T) {
	serializer := NewSerializer()
	tree := NewAVLTree("test_tree")
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)

	var buf bytes.Buffer
	err := serializer.SerializeTree(tree, &buf, TEXT)
	assert.NoError(t, err)

	output := buf.String()
	assert.Contains(t, output, "TREE test_tree 3")
	assert.Contains(t, output, "10")
	assert.Contains(t, output, "5")
	assert.Contains(t, output, "15")
}

func TestSerializer_HashTableTextFormat(t *testing.T) {
	serializer := NewSerializer()
	hash := NewHashTable("test_hash")
	hash.Insert("key1", "value1")
	hash.Insert("key2", "value2")

	var buf bytes.Buffer
	err := serializer.SerializeHashTable(hash, &buf, TEXT)
	assert.NoError(t, err)

	output := buf.String()
	assert.Contains(t, output, "HASH test_hash 2")
	assert.Contains(t, output, "key1")
	assert.Contains(t, output, "value1")
	assert.Contains(t, output, "key2")
	assert.Contains(t, output, "value2")
}

func TestSerializer_NilStructures(t *testing.T) {
	serializer := NewSerializer()
	var buf bytes.Buffer

	// Test serializing nil structures
	err := serializer.SerializeArray(nil, &buf, TEXT)
	assert.Error(t, err)

	err = serializer.SerializeSLL(nil, &buf, TEXT)
	assert.Error(t, err)

	err = serializer.SerializeDLL(nil, &buf, TEXT)
	assert.Error(t, err)

	err = serializer.SerializeStack(nil, &buf, TEXT)
	assert.Error(t, err)

	err = serializer.SerializeQueue(nil, &buf, TEXT)
	assert.Error(t, err)

	err = serializer.SerializeTree(nil, &buf, TEXT)
	assert.Error(t, err)

	err = serializer.SerializeHashTable(nil, &buf, TEXT)
	assert.Error(t, err)
}

func TestSerializer_EmptyStructures(t *testing.T) {
	serializer := NewSerializer()

	testCases := []struct {
		name     string
		structure interface{}
		serialize func(interface{}, *bytes.Buffer, SerializationFormat) error
	}{
		{
			name:      "Empty Array",
			structure: NewArray("empty_array"),
			serialize: func(s interface{}, buf *bytes.Buffer, format SerializationFormat) error {
				return serializer.SerializeArray(s.(*Array), buf, format)
			},
		},
		{
			name:      "Empty SLL",
			structure: NewSinglyLinkedList("empty_sll"),
			serialize: func(s interface{}, buf *bytes.Buffer, format SerializationFormat) error {
				return serializer.SerializeSLL(s.(*SinglyLinkedList), buf, format)
			},
		},
		{
			name:      "Empty DLL",
			structure: NewDoublyLinkedList("empty_dll"),
			serialize: func(s interface{}, buf *bytes.Buffer, format SerializationFormat) error {
				return serializer.SerializeDLL(s.(*DoublyLinkedList), buf, format)
			},
		},
		{
			name:      "Empty Stack",
			structure: NewStack("empty_stack"),
			serialize: func(s interface{}, buf *bytes.Buffer, format SerializationFormat) error {
				return serializer.SerializeStack(s.(*Stack), buf, format)
			},
		},
		{
			name:      "Empty Queue",
			structure: NewQueue("empty_queue"),
			serialize: func(s interface{}, buf *bytes.Buffer, format SerializationFormat) error {
				return serializer.SerializeQueue(s.(*Queue), buf, format)
			},
		},
		{
			name:      "Empty Tree",
			structure: NewAVLTree("empty_tree"),
			serialize: func(s interface{}, buf *bytes.Buffer, format SerializationFormat) error {
				return serializer.SerializeTree(s.(*AVLTree), buf, format)
			},
		},
		{
			name:      "Empty HashTable",
			structure: NewHashTable("empty_hash"),
			serialize: func(s interface{}, buf *bytes.Buffer, format SerializationFormat) error {
				return serializer.SerializeHashTable(s.(*HashTable), buf, format)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tc.serialize(tc.structure, &buf, TEXT)
			assert.NoError(t, err)
			assert.Greater(t, buf.Len(), 0)
		})
	}
}

func TestSerializer_WriteReadStringBinary(t *testing.T) {
	serializer := NewSerializer()
	
	testStrings := []string{
		"",
		"hello",
		"world",
		"this is a longer string with spaces",
		"special-chars!@#$%^&*()",
	}

	for _, testStr := range testStrings {
		var buf bytes.Buffer
		
		// Write string
		err := serializer.writeStringBinary(testStr, &buf)
		assert.NoError(t, err)
		
		// Read string back
		result, err := serializer.readStringBinary(&buf)
		assert.NoError(t, err)
		assert.Equal(t, testStr, result)
	}
}

func TestSerializer_WriteReadIntBinary(t *testing.T) {
	serializer := NewSerializer()
	
	testInts := []int{0, 1, -1, 100, -100, 1000, -1000, 123456, -123456}

	for _, testInt := range testInts {
		var buf bytes.Buffer
		
		// Write int
		err := serializer.writeIntBinary(testInt, &buf)
		assert.NoError(t, err)
		
		// Read int back
		result, err := serializer.readIntBinary(&buf)
		assert.NoError(t, err)
		assert.Equal(t, testInt, result)
	}
}
