package removeduplicatesfromsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 1, 2},
			want: []int{1, 2},
		},
		{
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			nums: []int{1, 2, 2, 3, 4, 4, 4, 4, 4, 5, 5, 6, 7, 7},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			nums: []int{0},
			want: []int{0},
		},
		{
			nums: []int{0, 0},
			want: []int{0},
		},
		{
			nums: []int{0, 1},
			want: []int{0, 1},
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.nums)
		got := solver.Solve()
		assert.Len(t, testCase.want, got)
		for j := range testCase.want {
			assert.Equal(t, testCase.want[j], testCase.nums[j], "Failed test case %d: [%d]", i+1, j)
		}
	}
}
