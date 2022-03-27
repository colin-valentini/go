package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfClosedIslands(t *testing.T) {
	assert.Equal(t, 2, NumberOfClosedIslands([][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}), "Example 1")
	assert.Equal(t, 1, NumberOfClosedIslands([][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0},
	}), "Example 2")
	assert.Equal(t, 2, NumberOfClosedIslands([][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}), "Example 3")
	assert.Equal(t, 0, NumberOfClosedIslands([][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}))
	assert.Equal(t, 5, NumberOfClosedIslands([][]int{
		{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
		{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 0},
	}))
}

func BenchmarkNumberOfClosedIslands(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = NumberOfClosedIslands([][]int{
			{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
			{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
			{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
			{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
			{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
			{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
			{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
			{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
			{1, 1, 1, 0, 1, 1, 0, 1, 1, 0},
		})
	}
}