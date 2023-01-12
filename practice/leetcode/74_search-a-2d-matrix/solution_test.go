package searcha2dmatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 3,
			want:   true,
		},
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 13,
			want:   false,
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.matrix, testCase.target)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
