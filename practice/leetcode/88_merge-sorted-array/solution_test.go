package mergesortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		m, n  int
		nums1 []int
		nums2 []int
		want  any
	}{
		{
			m:     3,
			n:     3,
			nums1: []int{1, 2, 3, 0, 0, 0},
			nums2: []int{2, 5, 6},
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			m:     1,
			n:     0,
			nums1: []int{1},
			nums2: []int{},
			want:  []int{1},
		},
		{
			m:     0,
			n:     1,
			nums1: []int{0},
			nums2: []int{1},
			want:  []int{1},
		},
		{
			m:     3,
			n:     3,
			nums1: []int{4, 5, 6, 0, 0, 0},
			nums2: []int{1, 2, 3},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			m:     5,
			n:     5,
			nums1: []int{1, 2, 7, 8, 10, 0, 0, 0, 0, 0},
			nums2: []int{2, 5, 6, 9, 11},
			want:  []int{1, 2, 2, 5, 6, 7, 8, 9, 10, 11},
		},
		{
			m:     0,
			n:     0,
			nums1: []int{},
			nums2: []int{},
			want:  []int{},
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.m, testCase.n, testCase.nums1, testCase.nums2)
		solver.Solve()
		assert.Equal(t, testCase.nums1, testCase.want, "Failed test case %d", i+1)
	}
}
