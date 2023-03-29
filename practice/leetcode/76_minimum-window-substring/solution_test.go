package minimumwindowsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		s, t string
		want string
	}{
		{s: "ADOBECODEBANC", t: "ABC", want: "BANC"},
		{s: "a", t: "a", want: "a"},
		{s: "a", t: "aa", want: ""},
		{s: "ABBBBCBBCBA", t: "ABC", want: "CBA"},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.s, testCase.t)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
