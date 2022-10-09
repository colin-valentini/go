package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		strs []string
		want string
	}{
		{strs: []string{"flower", "flow", "flight"}, want: "fl"},
		{strs: []string{"dog", "racecar", "car"}, want: ""},
		{strs: []string{"bazzz", "bazzzz", "bazzzzzzzzzz"}, want: "bazzz"},
		{strs: []string{"foo"}, want: "foo"},
		{strs: []string{}, want: ""},
	}
	for i, testCase := range testCases {
		solver := newLongestCommonPrefixSolver(testCase.strs)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Test case %d failed", i)
	}
}
