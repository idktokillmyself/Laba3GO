package dbmsgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashTable(t *testing.T) {
	t.Run("Constructor", func(t *testing.T) {
		table := NewHashTable("test_hash")
		assert.Equal(t, "test_hash", table.GetName())
		assert.True(t, table.IsEmpty())
		assert.Equal(t, 0, table.GetSize())
	})

	t.Run("InsertSearch", func(t *testing.T) {
		table := NewHashTable("test_hash")
		
		table.Insert("key1", "value1")
		table.Insert("key2", "value2")
		
		assert.False(t, table.IsEmpty())
		assert.Equal(t, 2, table.GetSize())
		
		value, found := table.Search("key1")
		assert.True(t, found)
		assert.Equal(t, "value1", value)
		
		value, found = table.Search("key2")
		assert.True(t, found)
		assert.Equal(t, "value2", value)
	})

	t.Run("UpdateValue", func(t *testing.T) {
		table := NewHashTable("test_hash")
		
		table.Insert("key1", "value1")
		table.Insert("key1", "new_value")
		
		value, found := table.Search("key1")
		assert.True(t, found)
		assert.Equal(t, "new_value", value)
		assert.Equal(t, 1, table.GetSize())
	})

	t.Run("Remove", func(t *testing.T) {
		table := NewHashTable("test_hash")
		
		table.Insert("key1", "value1")
		table.Insert("key2", "value2")
		
		removed := table.Remove("key1")
		assert.True(t, removed)
		assert.Equal(t, 1, table.GetSize())
		
		_, found := table.Search("key1")
		assert.False(t, found)
		
		_, found = table.Search("key2")
		assert.True(t, found)
		
		removed = table.Remove("non_existent")
		assert.False(t, removed)
	})

	t.Run("Resize", func(t *testing.T) {
		table := NewHashTable("test_hash")
		
		for i := 0; i < 20; i++ {
			table.Insert("key"+string(rune(i)), "value"+string(rune(i)))
		}
		
		assert.Equal(t, 20, table.GetSize())
		
		for i := 0; i < 20; i++ {
			value, found := table.Search("key" + string(rune(i)))
			assert.True(t, found)
			assert.Equal(t, "value"+string(rune(i)), value)
		}
	})

	t.Run("Collision", func(t *testing.T) {
		table := NewHashTable("test_hash")
		
		table.Insert("abc", "value1")
		table.Insert("cba", "value2")
		
		value1, found1 := table.Search("abc")
		assert.True(t, found1)
		assert.Equal(t, "value1", value1)
		
		value2, found2 := table.Search("cba")
		assert.True(t, found2)
		assert.Equal(t, "value2", value2)
	})

	t.Run("Cleanup", func(t *testing.T) {
		table := NewHashTable("test_hash")
		table.Insert("key1", "value1")
		table.Insert("key2", "value2")
		table.Insert("key3", "value3")
		
		assert.False(t, table.IsEmpty())
		
		table.Cleanup()
		
		assert.True(t, table.IsEmpty())
		assert.Equal(t, 0, table.GetSize())
		
		table.Insert("new_key", "new_value")
		assert.Equal(t, 1, table.GetSize())
		
		value, found := table.Search("new_key")
		assert.True(t, found)
		assert.Equal(t, "new_value", value)
	})
}
