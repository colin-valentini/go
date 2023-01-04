package containsduplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		in   []int
		want bool
	}{
		{in: []int{1, 2, 3, 1}, want: true},
		{in: []int{1, 2, 3, 4}, want: false},
		{in: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.in)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
