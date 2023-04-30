package redundantconnection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		edges [][]int
		want  []int
	}{
		{
			edges: [][]int{
				{1, 2},
				{1, 3},
				{2, 3},
			},
			want: []int{2, 3},
		},
		{
			edges: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{1, 4},
				{1, 5},
			},
			want: []int{1, 4},
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.edges)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
