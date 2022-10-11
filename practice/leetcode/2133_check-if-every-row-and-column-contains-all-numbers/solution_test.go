package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfEveryRowAndColumnContainsAllNumbers(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}
	want := true
	got := newCheckIfEveryRowAndColumnContainsAllNumbersSolver(matrix).Solve()
	assert.Equal(t, want, got)

	matrix = [][]int{
		{1, 1, 1},
		{1, 2, 3},
		{1, 2, 3},
	}
	want = false
	got = newCheckIfEveryRowAndColumnContainsAllNumbersSolver(matrix).Solve()
	assert.Equal(t, want, got)
}
