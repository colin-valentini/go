package sortcolors

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []TestCase{
		NewTestCase(t, []int{2, 0, 2, 1, 1, 0}),
		// NewTestCase(t, []int{2, 0, 1}),
		// NewTestCase(t, []int{}),
		// NewTestCase(t, []int{2, 1, 2, 1, 2, 0, 0}),
		// NewTestCase(t, []int{1, 1, 1, 1, 1, 1}),
		// NewTestCase(t, []int{2, 1, 1, 1, 1, 2}),
		// NewTestCase(t, []int{1, 1, 1, 1, 1, 0}),
	}
	for i, testCase := range testCases {
		NewSolver(testCase.elems).Solve() // Sorts the given slice in place.
		assert.Equal(t, testCase.want, testCase.elems, "Failed test case %d", i+1)
	}
}

type TestCase struct {
	elems []int
	want  []int
}

func NewTestCase(t *testing.T, elems []int) TestCase {
	t.Helper()
	cp := make([]int, len(elems))
	copy(cp, elems)
	sort.Slice(elems, func(i, j int) bool {
		return elems[i] < elems[j]
	})
	return TestCase{elems: cp, want: elems}
}
