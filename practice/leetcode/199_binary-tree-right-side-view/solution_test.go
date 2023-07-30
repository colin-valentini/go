package binarytreerightsideview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 4},
				},
			},
			want: []int{1, 3, 4},
		},
		{
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 3},
			},
			want: []int{1, 3},
		},
		{
			root: nil,
			want: []int{},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  6,
						Left: &TreeNode{Val: 8},
					},
				},
			},
			want: []int{1, 3, 6, 8},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{Val: 3},
			},
			want: []int{1, 3, 4, 5},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 6},
					},
					Right: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 7},
					},
				},
			},
			want: []int{1, 2, 4, 7},
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.root)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
