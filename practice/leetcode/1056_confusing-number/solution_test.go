package confusingnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		n    int
		want bool
	}{
		{n: 6, want: true},
		{n: 89, want: true},
		{n: 11, want: false},
		{n: 916, want: false},
		{n: 902, want: false},
		{n: 8000, want: true},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.n)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
