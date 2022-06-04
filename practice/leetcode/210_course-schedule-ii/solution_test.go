package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOrder(t *testing.T) {
	assert.Equal(t, []int{0, 1}, FindOrder(2, [][]int{{1, 0}}), "Example 1")
	assert.Equal(t, []int{0, 1, 2, 3}, FindOrder(4, [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2},
	}), "Example 2")
	assert.Equal(t, []int{0}, FindOrder(1, [][]int{}), "Example 3")
	assert.Equal(t, []int{2, 3, 1, 0, 4, 5}, FindOrder(6, [][]int{
		{0, 1},
		{0, 2},
		{1, 3},
		{3, 2},
		{4, 0},
		{5, 0},
	}))
	assert.Equal(t, []int{}, FindOrder(2, [][]int{
		{1, 0},
		{0, 1},
	}))
	assert.Equal(t, []int{}, FindOrder(4, [][]int{
		{1, 0},
		{2, 1},
		{3, 2},
		{0, 3},
	}))
}
