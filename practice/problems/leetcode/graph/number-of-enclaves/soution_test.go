package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfEnclaves(t *testing.T) {
	assert.Equal(t, 3, NumberOfEnclaves([][]int{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}), "Example 1")
	assert.Equal(t, 0, NumberOfEnclaves([][]int{
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}), "Example 2")
}

func BenchmarkNumberOfEnclaves(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = NumberOfEnclaves([][]int{
			{0, 0, 0, 0},
			{1, 0, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		})
	}
}