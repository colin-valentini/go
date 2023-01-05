package topkfrequentelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		want []int
	}{
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{nums: []int{1}, k: 1, want: []int{1}},
		{nums: []int{1, 1, 2, 2, 3, 3}, k: 3, want: []int{1, 2, 3}},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.nums, testCase.k)
		got := solver.Solve()
		assert.ElementsMatch(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
