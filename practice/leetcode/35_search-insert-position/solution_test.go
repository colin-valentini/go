package searchinsertposition

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
		{nums: []int{1, 3, 5, 6}, target: 5, want: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, want: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, want: 4},
		{nums: []int{1, 3, 5, 6}, target: 4, want: 2},
		{nums: []int{1, 3, 5, 6}, target: 0, want: 0},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.nums, testCase.target)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
