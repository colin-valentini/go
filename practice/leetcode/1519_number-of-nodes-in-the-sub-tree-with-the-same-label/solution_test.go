package numberofnodesinthesubtreewiththesamelabel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		n      int
		edges  [][]int
		labels string
		want   []int
	}{
		{
			n: 7,
			edges: [][]int{
				{0, 1},
				{0, 2},
				{1, 4},
				{1, 5},
				{2, 3},
				{2, 6},
			},
			labels: "abaedcd",
			want:   []int{2, 1, 1, 1, 1, 1, 1},
		},
		{
			n: 4, edges: [][]int{
				{0, 1},
				{1, 2},
				{0, 3},
			},
			labels: "bbbb",
			want:   []int{4, 2, 1, 1},
		},
		{
			n: 5,
			edges: [][]int{
				{0, 1},
				{0, 2},
				{1, 3},
				{0, 4},
			},
			labels: "aabab",
			want:   []int{3, 2, 1, 1, 1},
		},
		{
			n: 4,
			edges: [][]int{
				{0, 2},
				{0, 3},
				{1, 2},
			},
			labels: "aeed",
			want:   []int{1, 1, 2, 1},
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.n, testCase.edges, testCase.labels)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
