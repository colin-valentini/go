package searchinrotatedsortedarray

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
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, want: 4},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 1, want: 5},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 2, want: 6},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, want: -1},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 4, want: 0},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 5, want: 1},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 6, want: 2},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 7, want: 3},
		{nums: []int{1}, target: 0, want: -1},
		{nums: []int{3, 1, 2}, target: 3, want: 0},
		{nums: []int{3, 1, 2}, target: 1, want: 1},
		{nums: []int{3, 1, 2}, target: 2, want: 2},
		{nums: []int{4, 1, 2}, target: 3, want: -1},
		{nums: []int{3, 1}, target: 0, want: -1},
		{nums: []int{3, 1}, target: 3, want: 0},
		{nums: []int{3, 1}, target: 1, want: 1},
		{nums: []int{2, 4, 6, 8, 10}, target: 3, want: -1},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.nums, testCase.target)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
