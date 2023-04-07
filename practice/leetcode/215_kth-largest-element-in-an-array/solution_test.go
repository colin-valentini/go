package kthlargestelementinanarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		{
			nums: []int{3, 6, 7, 1, 2},
			k:    5,
			want: 1,
		},
		{
			nums: []int{3, 6, 7, 1, 2},
			k:    1,
			want: 7,
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.nums, testCase.k)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
