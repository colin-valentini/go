package longestsubstringwithoutrepeatingcharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		str  string
		want int
	}{
		{str: "abcabcbb", want: 3},
		{str: "bbbbb", want: 1},
		{str: "pwwkew", want: 3},
		{str: "dvdf", want: 3},
		{str: " ", want: 1},
		{str: "", want: 0},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.str)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
