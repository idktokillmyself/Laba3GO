package dbmsgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	t.Run("FindArray", func(t *testing.T) {
		db := NewDatabase()
		
		arr := NewArray("test_array")
		db.AddArray(arr)
		
		found := db.FindArray("test_array")
		assert.NotNil(t, found)
		assert.Equal(t, "test_array", found.GetName())
		
		notFound := db.FindArray("non_existent")
		assert.Nil(t, notFound)
	})

	t.Run("FindSLL", func(t *testing.T) {
		db := NewDatabase()
		
		sll := NewSinglyLinkedList("test_sll")
		db.AddSLL(sll)
		
		found := db.FindSLL("test_sll")
		assert.NotNil(t, found)
		assert.Equal(t, "test_sll", found.GetName())
	})

	t.Run("FindDLL", func(t *testing.T) {
		db := NewDatabase()
		
		dll := NewDoublyLinkedList("test_dll")
		db.AddDLL(dll)
		
		found := db.FindDLL("test_dll")
		assert.NotNil(t, found)
		assert.Equal(t, "test_dll", found.GetName())
	})

	t.Run("FindStack", func(t *testing.T) {
		db := NewDatabase()
		
		stack := NewStack("test_stack")
		db.AddStack(stack)
		
		found := db.FindStack("test_stack")
		assert.NotNil(t, found)
		assert.Equal(t, "test_stack", found.GetName())
	})

	t.Run("FindQueue", func(t *testing.T) {
		db := NewDatabase()
		
		queue := NewQueue("test_queue")
		db.AddQueue(queue)
		
		found := db.FindQueue("test_queue")
		assert.NotNil(t, found)
		assert.Equal(t, "test_queue", found.GetName())
	})

	t.Run("FindTree", func(t *testing.T) {
		db := NewDatabase()
		
		tree := NewAVLTree("test_tree")
		db.AddTree(tree)
		
		found := db.FindTree("test_tree")
		assert.NotNil(t, found)
		assert.Equal(t, "test_tree", found.GetName())
	})

	t.Run("FindHashTable", func(t *testing.T) {
		db := NewDatabase()
		
		table := NewHashTable("test_hash")
		db.AddHashTable(table)
		
		found := db.FindHashTable("test_hash")
		assert.NotNil(t, found)
		assert.Equal(t, "test_hash", found.GetName())
	})

	t.Run("Cleanup", func(t *testing.T) {
		db := NewDatabase()
		
		db.AddArray(NewArray("test_array"))
		db.AddSLL(NewSinglyLinkedList("test_sll"))
		db.AddStack(NewStack("test_stack"))
		
		assert.NotNil(t, db.FindArray("test_array"))
		assert.NotNil(t, db.FindSLL("test_sll"))
		assert.NotNil(t, db.FindStack("test_stack"))
		
		db.Cleanup()
		
		assert.Nil(t, db.FindArray("test_array"))
		assert.Nil(t, db.FindSLL("test_sll"))
		assert.Nil(t, db.FindStack("test_stack"))
	})
}
