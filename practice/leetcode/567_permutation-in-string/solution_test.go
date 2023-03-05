package permutationinstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		s1, s2 string
		want   bool
	}{
		{s1: "ab", s2: "eidbaooo", want: true},
		{s1: "ab", s2: "eidboaoo", want: false},
		{s1: "aabc", s2: "eidbaabc", want: true},
		{s1: "abc", s2: "ab", want: false},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.s1, testCase.s2)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
