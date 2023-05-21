package combinationsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want: [][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			want: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.candidates, testCase.target)
		got := solver.Solve()
		assert.ElementsMatch(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
