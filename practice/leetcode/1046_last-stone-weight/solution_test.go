package laststoneweight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastStoneWeight(t *testing.T) {
	testCases := []struct {
		stones []int
		want   int
	}{
		{stones: []int{2, 7, 4, 1, 8, 1}, want: 1},
		{stones: []int{1}, want: 1},
		{stones: []int{2, 2}, want: 0},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.stones)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
