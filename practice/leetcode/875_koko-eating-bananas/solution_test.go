package kokoeatingbananas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		piles []int
		h     int
		want  int
	}{
		{piles: []int{3, 6, 7, 11}, h: 8, want: 4},
		{piles: []int{30, 11, 23, 4, 20}, h: 5, want: 30},
		{piles: []int{30, 11, 23, 4, 20}, h: 6, want: 23},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.piles, testCase.h)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
