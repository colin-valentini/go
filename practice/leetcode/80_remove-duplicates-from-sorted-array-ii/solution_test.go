package removeduplicatesfromsortedarrayii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 1, 1, 2, 2, 3},
			want: []int{1, 1, 2, 2, 3},
		},
		{
			nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			nums: []int{1, 1, 1},
			want: []int{1, 1},
		},
		{
			nums: []int{1, 1, 1, 3},
			want: []int{1, 1, 3},
		},
		{
			nums: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
			want: []int{1, 1, 2},
		},
	}
	for i, testCase := range testCases {
		got := NewSolver(testCase.nums).Solve()
		k := len(testCase.want)
		assert.Equal(t, k, got, "Failed test case %d", i+1)
		assert.Equal(t, testCase.want, testCase.nums[:k], "Failed test case %d", i+1)
	}
}
