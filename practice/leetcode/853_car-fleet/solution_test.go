package carfleet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		target   int
		position []int
		speed    []int
		want     int
	}{
		{
			target:   12,
			position: []int{10, 8, 0, 5, 3},
			speed:    []int{2, 4, 1, 1, 3},
			want:     3,
		},
		{
			target:   10,
			position: []int{3},
			speed:    []int{3},
			want:     1,
		},
		{
			target:   100,
			position: []int{0, 2, 4},
			speed:    []int{4, 2, 1},
			want:     1,
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.target, testCase.position, testCase.speed)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
