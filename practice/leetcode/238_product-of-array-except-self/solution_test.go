package productofarrayexceptself

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{nums: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6}},
		{nums: []int{-1, 1, 0, -3, 3}, want: []int{0, 0, 9, 0, 0}},
		{nums: []int{4, 3, 7, 9, 11}, want: []int{2079, 2772, 1188, 924, 756}},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.nums)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
