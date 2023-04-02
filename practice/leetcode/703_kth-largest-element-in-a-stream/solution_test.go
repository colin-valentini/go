package kthlargestelementinastream

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		k    int
		nums []int
		adds []int
		want []int
	}{
		{
			k:    3,
			nums: []int{4, 5, 8, 2},
			adds: []int{3, 5, 10, 9, 4},
			want: []int{4, 5, 5, 8, 8},
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.k, testCase.nums)
		require.Len(t, testCase.want, len(testCase.adds))
		for j, add := range testCase.adds {
			want := testCase.want[j]
			assert.Equal(t, want, solver.Add(add), "Failed test case (%d, %d)", i+1, j+1)
		}
	}
}
