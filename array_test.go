package dbmsgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	t.Run("Constructor", func(t *testing.T) {
		arr := NewArray("test_array")
		assert.Equal(t, "test_array", arr.GetName())
		assert.True(t, arr.IsEmpty())
		assert.Equal(t, 0, arr.Length())
	})

	t.Run("PushBack", func(t *testing.T) {
		arr := NewArray("test_array")
		
		arr.PushBack("value1")
		assert.Equal(t, 1, arr.Length())
		val, err := arr.Get(0)
		assert.NoError(t, err)
		assert.Equal(t, "value1", val)
		
		arr.PushBack("value2")
		assert.Equal(t, 2, arr.Length())
		val, err = arr.Get(1)
		assert.NoError(t, err)
		assert.Equal(t, "value2", val)
	})

	t.Run("Insert", func(t *testing.T) {
		arr := NewArray("test_array")
		arr.PushBack("value1")
		arr.PushBack("value3")
		
		err := arr.Insert(1, "value2")
		assert.NoError(t, err)
		assert.Equal(t, 3, arr.Length())
		
		val, err := arr.Get(1)
		assert.NoError(t, err)
		assert.Equal(t, "value2", val)
		
		val, err = arr.Get(2)
		assert.NoError(t, err)
		assert.Equal(t, "value3", val)
	})

	t.Run("Remove", func(t *testing.T) {
		arr := NewArray("test_array")
		arr.PushBack("value1")
		arr.PushBack("value2")
		arr.PushBack("value3")
		
		err := arr.Remove(1)
		assert.NoError(t, err)
		assert.Equal(t, 2, arr.Length())
		
		val, err := arr.Get(0)
		assert.NoError(t, err)
		assert.Equal(t, "value1", val)
		
		val, err = arr.Get(1)
		assert.NoError(t, err)
		assert.Equal(t, "value3", val)
	})

	t.Run("Replace", func(t *testing.T) {
		arr := NewArray("test_array")
		arr.PushBack("old_value")
		
		err := arr.Replace(0, "new_value")
		assert.NoError(t, err)
		
		val, err := arr.Get(0)
		assert.NoError(t, err)
		assert.Equal(t, "new_value", val)
	})

	t.Run("OutOfBounds", func(t *testing.T) {
		arr := NewArray("test_array")
		
		_, err := arr.Get(0)
		assert.Error(t, err)
		
		err = arr.Insert(1, "value")
		assert.Error(t, err)
		
		err = arr.Remove(0)
		assert.Error(t, err)
		
		err = arr.Replace(0, "value")
		assert.Error(t, err)
	})

	t.Run("Cleanup", func(t *testing.T) {
		arr := NewArray("test_array")
		arr.PushBack("value1")
		arr.PushBack("value2")
		
		assert.False(t, arr.IsEmpty())
		assert.Equal(t, 2, arr.Length())
		
		arr.Cleanup()
		
		assert.True(t, arr.IsEmpty())
		assert.Equal(t, 0, arr.Length())
		
		_, err := arr.Get(0)
		assert.Error(t, err)
	})
}
