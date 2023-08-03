package majorityelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			nums: []int{6, 5, 5},
			want: 5,
		},
		{
			nums: []int{6, 6, 6, 7, 7},
			want: 6,
		},
	}
	for i, testCase := range testCases {
		got := NewSolver(testCase.nums).Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
