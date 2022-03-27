package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSurroundedRegions(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	SurroundedRegions(board)
	assert.Equal(t, [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}, board, "Example 1")

	board = [][]byte{
		{'X'},
	}
	SurroundedRegions(board)
	assert.Equal(t, [][]byte{
		{'X'},
	}, board, "Example 2")

	board = [][]byte{
		{'O', 'X', 'O', 'O', 'X', 'X'},
		{'O', 'X', 'X', 'X', 'O', 'X'},
		{'X', 'O', 'O', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X', 'X', 'X'},
		{'O', 'O', 'X', 'O', 'X', 'X'},
		{'X', 'X', 'O', 'O', 'O', 'O'},
	}
	SurroundedRegions(board)
	assert.Equal(t, [][]byte{
		{'O', 'X', 'O', 'O', 'X', 'X'},
		{'O', 'X', 'X', 'X', 'O', 'X'},
		{'X', 'O', 'O', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X', 'X', 'X'},
		{'O', 'O', 'X', 'O', 'X', 'X'},
		{'X', 'X', 'O', 'O', 'O', 'O'},
	}, board)
}

func BenchmarkSurroundedRegions(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SurroundedRegions([][]byte{
			{'O', 'X', 'O', 'O', 'X', 'X'},
			{'O', 'X', 'X', 'X', 'O', 'X'},
			{'X', 'O', 'O', 'X', 'O', 'O'},
			{'X', 'O', 'X', 'X', 'X', 'X'},
			{'O', 'O', 'X', 'O', 'X', 'X'},
			{'X', 'X', 'O', 'O', 'O', 'O'},
		})
	}
}
