package rotatearray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    1,
			want: []int{7, 1, 2, 3, 4, 5, 6},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    2,
			want: []int{6, 7, 1, 2, 3, 4, 5},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    4,
			want: []int{4, 5, 6, 7, 1, 2, 3},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    5,
			want: []int{3, 4, 5, 6, 7, 1, 2},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    6,
			want: []int{2, 3, 4, 5, 6, 7, 1},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    7,
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    14,
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    21,
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    23,
			want: []int{6, 7, 1, 2, 3, 4, 5},
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test Case: %d", i+1), func(t *testing.T) {
			require.GreaterOrEqual(t, testCase.k, 0, "Invalid test case")
			require.Len(t, testCase.want, len(testCase.nums), "Invalid test case")
			NewSolver(testCase.nums, testCase.k).Solve()
			assert.Equal(t, testCase.want, testCase.nums)
		})
	}
}
