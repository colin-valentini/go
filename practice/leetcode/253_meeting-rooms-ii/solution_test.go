package meetingroomsii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		intervals [][]int
		want      int
	}{
		{
			intervals: [][]int{
				{0, 30},
				{5, 10},
				{15, 20},
			},
			want: 2,
		},
		{
			intervals: [][]int{
				{7, 10},
				{2, 4},
			},
			want: 1,
		},
		{
			intervals: [][]int{
				{0, 5},
				{0, 10},
				{10, 30},
				{10, 15},
				{15, 30},
			},
			want: 2,
		},
		{
			intervals: [][]int{
				{0, 5},
				{0, 10},
				{0, 10},
				{0, 10},
				{5, 10},
				{10, 15},
				{15, 20},
			},
			want: 4,
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.intervals)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
