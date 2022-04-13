package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloodFill(t *testing.T) {
	assert.Equal(t, [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1},
	}, FloodFill([][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}, 1, 1, 2), "Example 1")

	assert.Equal(t, [][]int{
		{2, 2, 2},
		{2, 2, 2},
	}, FloodFill([][]int{
		{0, 0, 0},
		{0, 0, 0},
	}, 0, 0, 2), "Example 2")

	assert.Equal(t, [][]int{
		{2, 2, 2},
		{2, 2, 2},
	}, FloodFill([][]int{
		{0, 0, 0},
		{0, 0, 0},
	}, 1, 1, 2))
}

func BenchmarkFloodFill(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = FloodFill([][]int{
			{1, 1, 1},
			{1, 1, 0},
			{1, 0, 1},
		}, 1, 1, 2)
	}
}

func TestNeighbors(t *testing.T) {
	image := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	assert.Equal(t, []cell{
		{0, 1},
		{2, 1},
		{1, 0},
		{1, 2},
	}, neighbors(image, cell{1, 1}))
	assert.Equal(t, []cell{
		{1, 0},
		{0, 1},
	}, neighbors(image, cell{0, 0}))
	assert.Equal(t, []cell{
		{2, 2},
		{3, 1},
	}, neighbors(image, cell{3, 2}))
}

func TestIsInBounds(t *testing.T) {
	image := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	for i := range image {
		for j := range image[i] {
			assert.True(t, isInBounds(image, cell{i, j}))
		}
	}
	assert.False(t, isInBounds(image, cell{row: -1, col: 0}))
	assert.False(t, isInBounds(image, cell{row: 0, col: -1}))
	assert.False(t, isInBounds(image, cell{row: -1, col: -1}))
	assert.False(t, isInBounds(image, cell{row: len(image), col: 0}))
	assert.False(t, isInBounds(image, cell{row: 0, col: len(image[0])}))
	assert.False(t, isInBounds(image, cell{row: len(image), col: len(image[0])}))
}

func TestQueue(t *testing.T) {
	q := newQueue(10)
	assert.True(t, q.isEmpty())
	q.push(cell{11, 22})
	assert.False(t, q.isEmpty())
	q.push(cell{33, 44})
	assert.False(t, q.isEmpty())
	assert.Equal(t, cell{11, 22}, q.pop())
	assert.False(t, q.isEmpty())
	assert.Equal(t, cell{33, 44}, q.pop())
	assert.True(t, q.isEmpty())

	q = newQueue(3)
	q.push(cell{1, 2})
	q.push(cell{3, 4})
	q.push(cell{5, 6})
	assert.Panics(t, func() { q.push(cell{7, 8}) })
}
