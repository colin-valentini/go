package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestContainerWithMostWater runs test cases for ContainerWithMostWater.
func TestContainerWithMostWater(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, ContainerWithMostWater([]int{1,1}))
	assert.Equal(2, ContainerWithMostWater([]int{1,2,1}))
	assert.Equal(16, ContainerWithMostWater([]int{4,3,2,1,4}))
	assert.Equal(49, ContainerWithMostWater([]int{1,8,6,2,5,4,8,3,7}))
	assert.Equal(999_999_999, ContainerWithMostWater(testHelperHeightSlice(t)))
}

// BenchmarkContainerWithMostWater runs a benchmark for ContainerWithMostWater.
func BenchmarkContainerWithMostWater(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ContainerWithMostWater(benchmarkHelperHeightSlice(b))
	}
}

// testHelperHeightSlice is a benchmark helper around _makeDifficultHeightSlice.
func testHelperHeightSlice(t *testing.T) []int {
	t.Helper()
	return _makeDifficultHeightSlice()
}

// benchmarkHelperHeightSlice is a benchmark helper around _makeDifficultHeightSlice.
func benchmarkHelperHeightSlice(b *testing.B) []int {
	b.Helper()
	return _makeDifficultHeightSlice()
}

// _makeDifficultHeightSlice creates a difficult slice of heights and
// returns it. Intended to be called by test or benchmark helper functions.
func _makeDifficultHeightSlice() []int {
	height := make([]int, 1_000_000)
	for i := range height {
		height[i] = 1
		if i == 499_999 || i == 500_000 {
			height[i] = 999_999_999
		}
	}
	return height
}