package longestrepeatingcharacterreplacement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		str  string
		k    int
		want int
	}{
		{str: "ABAB", k: 2, want: 4},
		{str: "AABABBA", k: 1, want: 4},
		{str: "ABACADAEAFAG", k: 0, want: 1},
		{str: "ABACADAEAFAG", k: 1, want: 3},
		{str: "ABACADAEAFAG", k: 3, want: 7},
		{str: "ABACADAEAFAG", k: 4, want: 9},
		{str: "ABACADAEAFAG", k: 5, want: 11},
		{str: "ABACADAEAFAG", k: 6, want: 12},
		{str: "", k: 0, want: 0},
		{str: "", k: 1, want: 0},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.str, testCase.k)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
