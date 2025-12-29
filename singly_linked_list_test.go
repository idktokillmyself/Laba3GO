package dbmsgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList(t *testing.T) {
	t.Run("Constructor", func(t *testing.T) {
		sll := NewSinglyLinkedList("test_sll")
		assert.Equal(t, "test_sll", sll.GetName())
		assert.True(t, sll.IsEmpty())
	})

	t.Run("PushFront", func(t *testing.T) {
		sll := NewSinglyLinkedList("test_sll")
		
		sll.PushFront("value1")
		assert.False(t, sll.IsEmpty())
		assert.Equal(t, "value1", sll.GetHead().Data)
		
		sll.PushFront("value2")
		assert.Equal(t, "value2", sll.GetHead().Data)
	})

	t.Run("PushBack", func(t *testing.T) {
		sll := NewSinglyLinkedList("test_sll")
		
		sll.PushBack("value1")
		assert.Equal(t, "value1", sll.GetHead().Data)
		assert.Equal(t, "value1", sll.GetTail().Data)
		
		sll.PushBack("value2")
		assert.Equal(t, "value2", sll.GetTail().Data)
		assert.Equal(t, "value1", sll.GetHead().Data)
	})

	t.Run("InsertBefore", func(t *testing.T) {
		sll := NewSinglyLinkedList("test_sll")
		sll.PushBack("value1")
		sll.PushBack("value3")
		
		sll.InsertBefore("value3", "value2")
		assert.Equal(t, "value2", sll.GetHead().Next.Data)
	})

	t.Run("InsertAfter", func(t *testing.T) {
		sll := NewSinglyLinkedList("test_sll")
		sll.PushBack("value1")
		sll.PushBack("value3")
		
		sll.InsertAfter("value1", "value2")
		assert.Equal(t, "value2", sll.GetHead().Next.Data)
	})

	t.Run("DeleteFront", func(t *testing.T) {
		sll := NewSinglyLinkedList("test_sll")
		sll.PushBack("value1")
		sll.PushBack("value2")
		
		sll.DeleteFront()
		assert.Equal(t, "value2", sll.GetHead().Data)
	})

	t.Run("DeleteBack", func(t *testing.T) {
		sll := NewSinglyLinkedList("test_sll")
		sll.PushBack("value1")
		sll.PushBack("value2")
		
		sll.DeleteBack()
		assert.Equal(t, "value1", sll.GetTail().Data)
	})

	t.Run("DeleteByValue", func(t *testing.T) {
		sll := NewSinglyLinkedList("test_sll")
		sll.PushBack("value1")
		sll.PushBack("value2")
		sll.PushBack("value3")
		
		sll.DeleteByValue("value2")
		assert.Equal(t, "value3", sll.GetHead().Next.Data)
	})

	t.Run("FindByValue", func(t *testing.T) {
		sll := NewSinglyLinkedList("test_sll")
		sll.PushBack("value1")
		sll.PushBack("value2")
		
		found := sll.FindByValue("value2")
		assert.NotNil(t, found)
		assert.Equal(t, "value2", found.Data)
		
		notFound := sll.FindByValue("value3")
		assert.Nil(t, notFound)
	})

	t.Run("Cleanup", func(t *testing.T) {
		sll := NewSinglyLinkedList("test_sll")
		sll.PushBack("value1")
		sll.PushBack("value2")
		
		assert.False(t, sll.IsEmpty())
		
		sll.Cleanup()
		
		assert.True(t, sll.IsEmpty())
		assert.Nil(t, sll.GetHead())
		assert.Nil(t, sll.GetTail())
	})
}
