package balancedbinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolver(t *testing.T) {
	t.Run("Math", func(t *testing.T) {
		t.Run("Abs", func(t *testing.T) {
			require.Equal(t, 0, abs(0))
			require.Equal(t, 69, abs(-69))
			require.Equal(t, 69, abs(-69))
			require.Equal(t, 420-69, abs(69-420))
			require.Equal(t, 420-69, abs(420-69))
		})
		t.Run("Max", func(t *testing.T) {
			require.Equal(t, 420, max(69, 420))
			require.Equal(t, 420, max(420, 69))
		})
	})
	testCases := []struct {
		root *TreeNode
		want bool
	}{
		{
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: true,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 2},
			},
			want: false,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:  4,
							Left: &TreeNode{Val: 5},
						},
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val:   4,
							Right: &TreeNode{Val: 5},
						},
					},
				},
			},
			want: false,
		},
		{
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val:   4,
							Right: &TreeNode{Val: 5},
						},
					},
				},
			},
			want: false,
		},
		{
			root: &TreeNode{Val: 1},
			want: true,
		},
		{
			root: nil,
			want: true,
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.root)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}
