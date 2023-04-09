package taskscheduler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		tasks []byte
		n     int
		want  int
	}{
		{
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     2,
			want:  8,
		},
		{
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     0,
			want:  6,
		},
		{
			tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
			n:     2,
			want:  16,
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.tasks, testCase.n)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
