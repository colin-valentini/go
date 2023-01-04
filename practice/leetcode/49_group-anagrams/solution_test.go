package groupanagrams

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		strs []string
		want [][]string
	}{
		{
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			strs: []string{},
			want: [][]string{{}},
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.strs)
		got := solver.Solve()
		require.Len(t, got, len(testCase.want))
		// Sort the strings because ordering does not matter for this problem.
		for i := range got {
			sort.Strings(got[i])
			sort.Strings(testCase.want[i])
		}
		assert.ElementsMatch(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
