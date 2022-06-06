package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, LevelOrder(&TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}), "Example 1")
	assert.Equal(t, [][]int{{1}}, LevelOrder(&TreeNode{
		Val: 1,
	}), "Example 2")
	assert.Equal(t, [][]int{}, LevelOrder(nil), "Example 3")
}
