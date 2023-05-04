package graphvalidtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		n     int
		edges [][]int
		want  bool
	}{
		{
			n: 5,
			edges: [][]int{
				{0, 1},
				{0, 2},
				{0, 3},
				{1, 4},
			},
			want: true,
		},
		{
			n: 5,
			edges: [][]int{
				{0, 1},
				{1, 2},
				{2, 3},
				{1, 3},
				{1, 4},
			},
			want: false,
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.n, testCase.edges)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
