package binarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBinarySearchTree(t *testing.T) {
	less := func(x, y int) bool { return x < y }
	t.Run("Insert", func(t *testing.T) {
		tree := NewBinarySearchTree(less)
		two := tree.Insert(2)
		one := tree.Insert(1)
		three := tree.Insert(3)
		require.Equal(t, tree.root, two)
		require.Equal(t, tree.root.left, one)
		require.Equal(t, tree.root.right, three)
	})
	t.Run("Iter", func(t *testing.T) {
		tree := NewBinarySearchTree(less)
		_ = tree.Insert(2)
		_ = tree.Insert(1)
		_ = tree.Insert(3)
		want := newTestBinarySearchTree(t)
		it := want.Iter()
		for _, v := range []int{2, 1, 3} {
			require.True(t, it.next())
			want, got := v, it.value().value
			assert.Equal(t, want, got)
		}
	})
	t.Run("Find", func(t *testing.T) {
		tree := NewBinarySearchTree(less)
		_ = tree.Insert(2)
		_ = tree.Insert(1)
		_ = tree.Insert(3)
		assert.Equal(t, tree.root, tree.Find(2))
		assert.Equal(t, tree.root.left, tree.Find(1))
		assert.Equal(t, tree.root.right, tree.Find(3))
	})
	t.Run("Delete", func(t *testing.T) {
		tree := NewBinarySearchTree(less)
		two := tree.Insert(2)
		one := tree.Insert(1)
		three := tree.Insert(3)
		tree.Delete(two)
		assert.Nil(t, tree.Find(two.value))
		assert.Equal(t, three, tree.root)
		tree.Delete(one)
		assert.Nil(t, tree.Find(one.value))
		assert.Equal(t, three, tree.root)
		tree.Delete(three)
		assert.Nil(t, tree.Find(three.value))
		assert.Nil(t, tree.root)
	})
}

// newTestBinarySearchTree returns a tree of the following structure:
// root: {value: 1, left: {value: 2}, right: {value: 3}}
func newTestBinarySearchTree(t *testing.T) *BinarySearchTree[int] {
	t.Helper()
	root := &BstNode[int]{value: 2}
	left := &BstNode[int]{value: 1, parent: root}
	right := &BstNode[int]{value: 3, parent: root}
	root.left, root.right = left, right
	return &BinarySearchTree[int]{
		root: root,
		less: func(x, y int) bool { return x < y },
	}
}
