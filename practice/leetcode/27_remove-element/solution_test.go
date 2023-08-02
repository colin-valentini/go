package removeelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		nums []int
		val  int
		want int
	}{
		{
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: 2,
		},
		{
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: 5,
		},
		{
			nums: []int{},
			val:  1,
			want: 0,
		},
		{
			nums: []int{1},
			val:  1,
			want: 0,
		},
		{
			nums: []int{2},
			val:  1,
			want: 1,
		},
		{
			nums: []int{1, 2, 3},
			val:  4,
			want: 3,
		},
		{
			nums: []int{1, 1, 1},
			val:  1,
			want: 0,
		},
		{
			nums: []int{1, 1, 3},
			val:  1,
			want: 1,
		},
		{
			nums: []int{2, 1, 3},
			val:  1,
			want: 2,
		},
	}
	for i, testCase := range testCases {
		got := NewSolver(testCase.nums, testCase.val).Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
		for i := 0; i < got; i++ {
			assert.NotEqual(t, testCase.val, testCase.nums[i])
		}
	}
}
