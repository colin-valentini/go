package binarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var less = func(x, y int) bool { return x < y }

func TestBinarySearchTree(t *testing.T) {
	t.Run("Insert", func(t *testing.T) {
		tree := NewBinarySearchTree(less)
		two := tree.Insert(2)
		one := tree.Insert(1)
		three := tree.Insert(3)
		four := tree.Insert(4)
		zero := tree.Insert(0)
		require.Equal(t, tree.root, two)
		require.Equal(t, tree.root.left, one)
		require.Equal(t, tree.root.right, three)
		require.Equal(t, tree.root.right.right, four)
		require.Equal(t, tree.root.left.left, zero)
	})
	t.Run("Iter", func(t *testing.T) {
		tree, _ := newTestTree(t)
		it := tree.Iter()
		for _, v := range []int{8, 3, 10, 1, 6, 14, 4, 7, 13} {
			require.True(t, it.next())
			want, got := v, it.value().value
			assert.Equal(t, want, got)
		}
		assert.False(t, it.next())
	})
	t.Run("Find", func(t *testing.T) {
		tree, _ := newTestTree(t)
		assert.Equal(t, tree.root, tree.Find(8))
		assert.Equal(t, tree.root.left, tree.Find(3))
		assert.Equal(t, tree.root.right, tree.Find(10))
		assert.Equal(t, tree.root.left.left, tree.Find(1))
		assert.Equal(t, tree.root.left.right, tree.Find(6))
		assert.Equal(t, tree.root.right.right, tree.Find(14))
		assert.Equal(t, tree.root.left.right.left, tree.Find(4))
		assert.Equal(t, tree.root.left.right.right, tree.Find(7))
		assert.Equal(t, tree.root.right.right.left, tree.Find(13))
	})
	t.Run("Delete", func(t *testing.T) {
		tree, nodes := newTestTree(t)

		tree.Delete(nodes[8])
		assert.Nil(t, tree.Find(8))
		assert.Equal(t, 10, tree.root.value)
		assert.Equal(t, 3, tree.root.left.value)
		assert.Equal(t, 14, tree.root.right.value)

		tree.Delete(nodes[10])
		assert.Nil(t, tree.Find(10))
		assert.Equal(t, 13, tree.root.value)
		assert.Equal(t, 3, tree.root.left.value)
		assert.Equal(t, 14, tree.root.right.value)

		tree.Delete(nodes[13])
		assert.Nil(t, tree.Find(13))
		assert.Equal(t, 14, tree.root.value)
		assert.Equal(t, 3, tree.root.left.value)
		assert.Nil(t, tree.root.right)

		tree.Delete(nodes[14])
		assert.Nil(t, tree.Find(14))
		assert.Equal(t, 3, tree.root.value)
		assert.Equal(t, 1, tree.root.left.value)
		assert.Equal(t, 6, tree.root.right.value)

		tree.Delete(nodes[7])
		assert.Nil(t, tree.Find(7))
		assert.Equal(t, 3, tree.root.value)
		assert.Equal(t, 1, tree.root.left.value)
		assert.Equal(t, 6, tree.root.right.value)

		tree.Delete(nodes[6])
		assert.Nil(t, tree.Find(6))
		assert.Equal(t, 3, tree.root.value)
		assert.Equal(t, 1, tree.root.left.value)
		assert.Equal(t, 4, tree.root.right.value)

		tree.Delete(nodes[3])
		assert.Nil(t, tree.Find(3))
		assert.Equal(t, 4, tree.root.value)
		assert.Equal(t, 1, tree.root.left.value)
		assert.Nil(t, tree.root.right)

		tree.Delete(nodes[1])
		assert.Nil(t, tree.Find(1))
		assert.Equal(t, 4, tree.root.value)
		assert.Nil(t, tree.root.left)

		tree.Delete(nodes[4])
		assert.Nil(t, tree.Find(4))
		assert.Nil(t, tree.root)
	})

	t.Run("successor", func(t *testing.T) {
		tree, nodes := newTestTree(t)
		assert.Equal(t, nodes[10], tree.successor(nodes[8]))
		assert.Equal(t, nodes[4], tree.successor(nodes[3]))
		assert.Equal(t, nodes[13], tree.successor(nodes[10]))
		assert.Nil(t, tree.successor(nodes[14])) // No successor
	})
}

func newTestTree(t *testing.T) (*BinarySearchTree[int], map[int]*BstNode[int]) {
	//        8
	//     ↙     ↘
	//    3       10
	//   ↙ ↘        ↘
	//  1   6       14
	//     ↙ ↘     ↙
	//    4   7   13
	t.Helper()
	tree := NewBinarySearchTree(less)
	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	nodes := make(map[int]*BstNode[int], len(values))
	for _, value := range values {
		nodes[value] = tree.Insert(value)
	}
	return tree, nodes
}
