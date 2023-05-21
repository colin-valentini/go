package subsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 2, 3},
			want: [][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
		{
			nums: []int{0},
			want: [][]int{
				{},
				{0},
			},
		},
		{
			nums: []int{1, 2},
			want: [][]int{
				{},
				{1},
				{2},
				{1, 2},
			},
		},
		{
			nums: []int{1, 2, 3, 4},
			want: [][]int{
				{},
				{1},
				{2},
				{3},
				{4},
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{3, 4},
				{1, 2, 3},
				{1, 2, 4},
				{1, 3, 4},
				{2, 3, 4},
				{1, 2, 3, 4},
			},
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.nums)
		got := solver.Solve()
		assert.ElementsMatch(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
