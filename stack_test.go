package dbmsgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("Constructor", func(t *testing.T) {
		stack := NewStack("test_stack")
		assert.Equal(t, "test_stack", stack.GetName())
		assert.True(t, stack.IsEmpty())
		assert.Equal(t, 0, stack.GetSize())
	})

	t.Run("PushPop", func(t *testing.T) {
		stack := NewStack("test_stack")
		
		stack.Push("value1")
		assert.False(t, stack.IsEmpty())
		assert.Equal(t, 1, stack.GetSize())
		
		peekVal, err := stack.Peek()
		assert.NoError(t, err)
		assert.Equal(t, "value1", peekVal)
		
		stack.Push("value2")
		assert.Equal(t, 2, stack.GetSize())
		
		peekVal, err = stack.Peek()
		assert.NoError(t, err)
		assert.Equal(t, "value2", peekVal)
		
		popVal, err := stack.Pop()
		assert.NoError(t, err)
		assert.Equal(t, "value2", popVal)
		assert.Equal(t, 1, stack.GetSize())
		
		peekVal, err = stack.Peek()
		assert.NoError(t, err)
		assert.Equal(t, "value1", peekVal)
	})

	t.Run("PopEmpty", func(t *testing.T) {
		stack := NewStack("test_stack")
		
		_, err := stack.Pop()
		assert.Error(t, err)
	})

	t.Run("PeekEmpty", func(t *testing.T) {
		stack := NewStack("test_stack")
		
		_, err := stack.Peek()
		assert.Error(t, err)
	})

	t.Run("LIFO", func(t *testing.T) {
		stack := NewStack("test_stack")
		
		stack.Push("first")
		stack.Push("second")
		stack.Push("third")
		
		val, err := stack.Pop()
		assert.NoError(t, err)
		assert.Equal(t, "third", val)
		
		val, err = stack.Pop()
		assert.NoError(t, err)
		assert.Equal(t, "second", val)
		
		val, err = stack.Pop()
		assert.NoError(t, err)
		assert.Equal(t, "first", val)
		
		assert.True(t, stack.IsEmpty())
	})

	t.Run("Cleanup", func(t *testing.T) {
		stack := NewStack("test_stack")
		stack.Push("value1")
		stack.Push("value2")
		
		assert.False(t, stack.IsEmpty())
		
		stack.Cleanup()
		
		assert.True(t, stack.IsEmpty())
		assert.Equal(t, 0, stack.GetSize())
	})
}
