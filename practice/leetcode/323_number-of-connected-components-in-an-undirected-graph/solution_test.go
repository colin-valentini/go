package numberofconnectedcomponentsinanundirectedgraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		n     int
		edges [][]int
		want  int
	}{
		{
			n: 5,
			edges: [][]int{
				{0, 1},
				{1, 2},
				{3, 4},
			},
			want: 2,
		},
		{
			n: 5,
			edges: [][]int{
				{0, 1},
				{1, 2},
				{2, 3},
				{3, 4},
			},
			want: 1,
		},
		{
			n: 5,
			edges: [][]int{
				{0, 1},
				{2, 1},
				{2, 0},
				{2, 4},
			},
			want: 2,
		},
		{
			n: 5,
			edges: [][]int{
				{0, 1},
				{0, 2},
				{1, 2},
				{2, 3},
				{2, 4},
			},
			want: 1,
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.n, testCase.edges)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
