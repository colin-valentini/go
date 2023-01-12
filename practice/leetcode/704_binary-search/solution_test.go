package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, want: 4},
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, want: -1},
		{nums: []int{}, target: 69, want: -1},
		{nums: []int{1, 2, 3}, target: 2, want: 1},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.nums, testCase.target)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
