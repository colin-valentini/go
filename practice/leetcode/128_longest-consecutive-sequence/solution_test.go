package longestconsecutivesequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{nums: []int{100, 4, 200, 1, 3, 2}, want: 4},
		{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, want: 9},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.nums)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
