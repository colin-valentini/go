package findminimuminrotatedsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{nums: []int{3, 4, 5, 1, 2}, want: 1},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, want: 0},
		{nums: []int{11, 13, 15, 17}, want: 11},
		{nums: []int{5, 1, 2, 3, 4}, want: 1},
		{nums: []int{3, 1, 2}, want: 1},
		{nums: []int{1}, want: 1},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.nums)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
