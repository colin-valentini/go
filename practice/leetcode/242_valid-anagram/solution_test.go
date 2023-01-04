package validanagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		s, t string
		want bool
	}{
		{s: "anagram", t: "nagaram", want: true},
		{s: "rat", t: "car", want: false},
		{s: "abc", t: "ab", want: false},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.s, testCase.t)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
