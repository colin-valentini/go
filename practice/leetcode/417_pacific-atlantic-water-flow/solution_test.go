package pacificatlanticwaterflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		heights [][]int
		want    [][]int
	}{
		{
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			want: [][]int{
				{0, 4},
				{1, 3},
				{1, 4},
				{2, 2},
				{3, 0},
				{3, 1},
				{4, 0},
			},
		},
		{
			heights: [][]int{{1}},
			want:    [][]int{{0, 0}},
		},
		{
			heights: [][]int{
				{10, 10, 10},
				{10, 1, 10},
				{10, 10, 10},
			},
			want: [][]int{
				{0, 0},
				{0, 1},
				{0, 2},
				{1, 0},
				{1, 2},
				{2, 0},
				{2, 1},
				{2, 2},
			},
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.heights)
		got := solver.Solve()
		assert.ElementsMatch(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
