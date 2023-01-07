package dailytemperatures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		temps []int
		want  []int
	}{
		{temps: []int{73, 74, 75, 71, 69, 72, 76, 73}, want: []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{temps: []int{30, 40, 50, 60}, want: []int{1, 1, 1, 0}},
		{temps: []int{30, 60, 90}, want: []int{1, 1, 0}},
		{temps: []int{99, 98, 97, 98, 97, 100}, want: []int{5, 4, 1, 2, 1, 0}},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.temps)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
