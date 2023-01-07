package evaluatereversepolishnotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		tokens []string
		want   int
	}{
		{tokens: []string{
			"2",
			"1",
			"+",
			"3",
			"*",
		}, want: 9},
		{tokens: []string{
			"4",
			"13",
			"5",
			"/",
			"+",
		}, want: 6},
		{tokens: []string{
			"10",
			"6",
			"9",
			"3",
			"+",
			"-11",
			"*",
			"/",
			"*",
			"17",
			"+",
			"5",
			"+",
		}, want: 22},
		{tokens: []string{
			"1",
			"1",
			"-",
		}, want: 0},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.tokens)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}

	// For test coverage.
	assert.Panics(t, func() {
		_ = operator("$").operate(1, 2)
	})
}
