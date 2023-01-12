package validpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		str  string
		want bool
	}{
		{str: "A man, a plan, a canal: Panama", want: true},
		{str: "race a car", want: false},
		{str: " ", want: true},
		{str: "^|a{  b||$ a|*", want: true},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.str)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
