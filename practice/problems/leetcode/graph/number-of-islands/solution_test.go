package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	assert.Equal(t, 1, NumIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}), "Example 1")
	assert.Equal(t, 3, NumIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}), "Example 2")
	assert.Equal(t, 0, NumIslands([][]byte{
		{'0'},
	}))
	assert.Equal(t, 1, NumIslands([][]byte{
		{'1'},
	}))
	assert.Equal(t, 0, NumIslands([][]byte{
		{'0'},
		{'0'},
	}))
	assert.Equal(t, 1, NumIslands([][]byte{
		{'1'},
		{'1'},
	}))
	assert.Equal(t, 0, NumIslands([][]byte{
		{'0', '0'},
	}))
	assert.Equal(t, 1, NumIslands([][]byte{
		{'1', '1'},
	}))
	assert.Equal(t, 0, NumIslands([][]byte{
		{'0', '0'},
		{'0', '0'},
		{'0', '0'},
	}))
	assert.Equal(t, 1, NumIslands([][]byte{
		{'1', '1'},
		{'1', '1'},
		{'1', '1'},
	}))
	assert.Equal(t, 5, NumIslands([][]byte{
		{'1', '0'},
		{'0', '1'},
		{'1', '0'},
		{'0', '1'},
		{'1', '0'},
	}))
	assert.Equal(t, 1, NumIslands([][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}))
}

func BenchmarkNumIslands(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = NumIslands([][]byte{
			{'1', '1', '1'},
			{'0', '1', '0'},
			{'1', '1', '1'},
		})
	}
}