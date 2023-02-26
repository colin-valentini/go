package besttimetobuyandsellstock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			prices: []int{1, 2},
			want:   1,
		},
		{
			prices: []int{2, 1},
			want:   0,
		},
		{
			prices: []int{1},
			want:   0,
		},
		{
			prices: []int{},
			want:   0,
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.prices)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
