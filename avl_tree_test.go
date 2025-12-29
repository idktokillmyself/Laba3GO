package dbmsgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLTree(t *testing.T) {
	t.Run("Constructor", func(t *testing.T) {
		tree := NewAVLTree("test_tree")
		assert.Equal(t, "test_tree", tree.GetName())
		assert.True(t, tree.IsEmpty())
	})

	t.Run("InsertSearch", func(t *testing.T) {
		tree := NewAVLTree("test_tree")
		
		tree.Insert(50)
		tree.Insert(30)
		tree.Insert(70)
		tree.Insert(20)
		tree.Insert(40)
		
		assert.False(t, tree.IsEmpty())
		assert.NotNil(t, tree.Search(50))
		assert.NotNil(t, tree.Search(30))
		assert.Nil(t, tree.Search(100))
	})

	t.Run("Remove", func(t *testing.T) {
		tree := NewAVLTree("test_tree")
		
		tree.Insert(50)
		tree.Insert(30)
		tree.Insert(70)
		
		tree.Remove(30)
		assert.Nil(t, tree.Search(30))
		assert.NotNil(t, tree.Search(50))
		assert.NotNil(t, tree.Search(70))
	})

	t.Run("CountElements", func(t *testing.T) {
		tree := NewAVLTree("test_tree")
		
		assert.Equal(t, 0, tree.CountElements())
		
		tree.Insert(50)
		tree.Insert(30)
		tree.Insert(70)
		
		assert.Equal(t, 3, tree.CountElements())
		
		tree.Remove(30)
		assert.Equal(t, 2, tree.CountElements())
	})

	t.Run("SaveTree", func(t *testing.T) {
		tree := NewAVLTree("test_tree")
		
		tree.Insert(50)
		tree.Insert(30)
		tree.Insert(70)
		tree.Insert(20)
		tree.Insert(40)
		
		saved := tree.SaveTree()
		assert.Equal(t, 5, len(saved))
		
		for i := 1; i < len(saved); i++ {
			assert.True(t, saved[i] > saved[i-1])
		}
	})

	t.Run("DuplicateInsert", func(t *testing.T) {
		tree := NewAVLTree("test_tree")
		
		tree.Insert(50)
		tree.Insert(50)
		
		assert.Equal(t, 1, tree.CountElements())
	})

	t.Run("RemoveNonExistent", func(t *testing.T) {
		tree := NewAVLTree("test_tree")
		tree.Insert(50)
		tree.Insert(30)
		
		assert.NotPanics(t, func() {
			tree.Remove(100)
		})
		assert.Equal(t, 2, tree.CountElements())
	})

	t.Run("Cleanup", func(t *testing.T) {
		tree := NewAVLTree("test_tree")
		tree.Insert(50)
		tree.Insert(30)
		tree.Insert(70)
		
		assert.False(t, tree.IsEmpty())
		
		tree.Cleanup()
		
		assert.True(t, tree.IsEmpty())
		assert.Equal(t, 0, tree.CountElements())
	})
}
