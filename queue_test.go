package dbmsgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("Constructor", func(t *testing.T) {
		queue := NewQueue("test_queue")
		assert.Equal(t, "test_queue", queue.GetName())
		assert.True(t, queue.IsEmpty())
		assert.Equal(t, 0, queue.GetSize())
	})

	t.Run("PushPop", func(t *testing.T) {
		queue := NewQueue("test_queue")
		
		queue.Push("value1")
		assert.False(t, queue.IsEmpty())
		assert.Equal(t, 1, queue.GetSize())
		
		peekVal, err := queue.Peek()
		assert.NoError(t, err)
		assert.Equal(t, "value1", peekVal)
		
		queue.Push("value2")
		assert.Equal(t, 2, queue.GetSize())
		
		peekVal, err = queue.Peek()
		assert.NoError(t, err)
		assert.Equal(t, "value1", peekVal)
		
		popVal, err := queue.Pop()
		assert.NoError(t, err)
		assert.Equal(t, "value1", popVal)
		assert.Equal(t, 1, queue.GetSize())
		
		peekVal, err = queue.Peek()
		assert.NoError(t, err)
		assert.Equal(t, "value2", peekVal)
	})

	t.Run("PopEmpty", func(t *testing.T) {
		queue := NewQueue("test_queue")
		
		_, err := queue.Pop()
		assert.Error(t, err)
	})

	t.Run("PeekEmpty", func(t *testing.T) {
		queue := NewQueue("test_queue")
		
		_, err := queue.Peek()
		assert.Error(t, err)
	})

	t.Run("FIFO", func(t *testing.T) {
		queue := NewQueue("test_queue")
		
		queue.Push("first")
		queue.Push("second")
		queue.Push("third")
		
		val, err := queue.Pop()
		assert.NoError(t, err)
		assert.Equal(t, "first", val)
		
		val, err = queue.Pop()
		assert.NoError(t, err)
		assert.Equal(t, "second", val)
		
		val, err = queue.Pop()
		assert.NoError(t, err)
		assert.Equal(t, "third", val)
		
		assert.True(t, queue.IsEmpty())
	})

	t.Run("Cleanup", func(t *testing.T) {
		queue := NewQueue("test_queue")
		queue.Push("value1")
		queue.Push("value2")
		
		assert.False(t, queue.IsEmpty())
		
		queue.Cleanup()
		
		assert.True(t, queue.IsEmpty())
		assert.Equal(t, 0, queue.GetSize())
	})
}
