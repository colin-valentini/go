package wallsandgates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		rooms [][]int
		want  [][]int
	}{
		{
			rooms: [][]int{
				{INF, -1, 0, INF},
				{INF, INF, INF, -1},
				{INF, -1, INF, -1},
				{0, -1, INF, INF},
			},
			want: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			rooms: [][]int{{-1}},
			want:  [][]int{{-1}},
		},
		{
			rooms: [][]int{
				{INF, -1, 0, INF},
				{INF, INF, INF, -1},
				{INF, -1, -1, -1},
				{0, -1, INF, INF},
			},
			want: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, -1, -1},
				{0, -1, INF, INF},
			},
		},
		{
			rooms: [][]int{},
			want:  [][]int{},
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.rooms)
		solver.Solve()
		assert.Equal(t, testCase.want, testCase.rooms, "Failed test case %d", i+1)
	}
}
