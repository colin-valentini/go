package balancedbinarytree

// LeetCode #110.
// https://leetcode.com/problems/balanced-binary-tree/

// Given a binary tree, determine if it is height-balanced.
//
// A height-balanced binary tree is a binary tree in which the
// depth of the two subtrees of every node never differs by more
// than one.

// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
//  func isBalanced(root *TreeNode) bool {
//
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solver struct {
	root *TreeNode
}

func NewSolver(root *TreeNode) *Solver {
	return &Solver{root: root}
}

func (s *Solver) Solve() bool {
	isBalanced, _ := s.solve(s.root)
	return isBalanced
}

func (s *Solver) solve(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}
	leftBalanced, leftHeight := s.solve(node.Left)
	if !leftBalanced {
		return false, 0
	}
	rightBalanced, rightHeight := s.solve(node.Right)
	if !rightBalanced {
		return false, 0
	}
	return s.compare(leftHeight, rightHeight), max(leftHeight, rightHeight) + 1
}

func (s *Solver) compare(leftHeight, rightHeight int) bool {
	return abs(leftHeight-rightHeight) <= 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
