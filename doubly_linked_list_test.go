package dbmsgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList(t *testing.T) {
	t.Run("Constructor", func(t *testing.T) {
		dll := NewDoublyLinkedList("test_dll")
		assert.Equal(t, "test_dll", dll.GetName())
		assert.True(t, dll.IsEmpty())
	})

	t.Run("PushFront", func(t *testing.T) {
		dll := NewDoublyLinkedList("test_dll")
		
		dll.PushFront("value1")
		assert.Equal(t, "value1", dll.GetHead().Data)
		assert.Equal(t, "value1", dll.GetTail().Data)
		
		dll.PushFront("value2")
		assert.Equal(t, "value2", dll.GetHead().Data)
		assert.Equal(t, "value1", dll.GetHead().Next.Data)
		assert.Equal(t, "value2", dll.GetTail().Prev.Data)
	})

	t.Run("PushBack", func(t *testing.T) {
		dll := NewDoublyLinkedList("test_dll")
		
		dll.PushBack("value1")
		assert.Equal(t, "value1", dll.GetTail().Data)
		
		dll.PushBack("value2")
		assert.Equal(t, "value2", dll.GetTail().Data)
		assert.Equal(t, "value1", dll.GetTail().Prev.Data)
	})

	t.Run("InsertBefore", func(t *testing.T) {
		dll := NewDoublyLinkedList("test_dll")
		dll.PushBack("value1")
		dll.PushBack("value3")
		
		dll.InsertBefore("value3", "value2")
		assert.Equal(t, "value2", dll.GetHead().Next.Data)
		assert.Equal(t, "value2", dll.GetTail().Prev.Data)
	})

	t.Run("InsertAfter", func(t *testing.T) {
		dll := NewDoublyLinkedList("test_dll")
		dll.PushBack("value1")
		dll.PushBack("value3")
		
		dll.InsertAfter("value1", "value2")
		assert.Equal(t, "value2", dll.GetHead().Next.Data)
		assert.Equal(t, "value2", dll.GetTail().Prev.Data)
	})

	t.Run("DeleteFront", func(t *testing.T) {
		dll := NewDoublyLinkedList("test_dll")
		dll.PushBack("value1")
		dll.PushBack("value2")
		
		dll.DeleteFront()
		assert.Equal(t, "value2", dll.GetHead().Data)
		assert.Nil(t, dll.GetHead().Prev)
	})

	t.Run("DeleteBack", func(t *testing.T) {
		dll := NewDoublyLinkedList("test_dll")
		dll.PushBack("value1")
		dll.PushBack("value2")
		
		dll.DeleteBack()
		assert.Equal(t, "value1", dll.GetTail().Data)
		assert.Nil(t, dll.GetTail().Next)
	})

	t.Run("DeleteByValue", func(t *testing.T) {
		dll := NewDoublyLinkedList("test_dll")
		dll.PushBack("value1")
		dll.PushBack("value2")
		dll.PushBack("value1")
		
		dll.DeleteByValue("value1")
		assert.Equal(t, "value2", dll.GetHead().Data)
		assert.Equal(t, "value2", dll.GetTail().Data)
	})

	t.Run("FindByValue", func(t *testing.T) {
		dll := NewDoublyLinkedList("test_dll")
		dll.PushBack("value1")
		dll.PushBack("value2")
		
		found := dll.FindByValue("value2")
		assert.NotNil(t, found)
		assert.Equal(t, "value2", found.Data)
	})

	t.Run("Cleanup", func(t *testing.T) {
		dll := NewDoublyLinkedList("test_dll")
		dll.PushBack("value1")
		dll.PushBack("value2")
		
		assert.False(t, dll.IsEmpty())
		
		dll.Cleanup()
		
		assert.True(t, dll.IsEmpty())
		assert.Nil(t, dll.GetHead())
		assert.Nil(t, dll.GetTail())
	})
}
