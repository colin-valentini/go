package leetcode

// LeetCode #102.
// Link: https://leetcode.com/problems/binary-tree-level-order-traversal/

// Given the root of a binary tree, return the level order traversal of its
// nodes' values. (i.e., from left to right, level by level).

// Example 1.
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]

// Example 2.
// Input: root = [1]
// Output: [[1]]

// Example 3.
// Input: root = []
// Output: []

// Constraints:
// - The number of nodes in the tree is in the range [0, 2000].
// - -1000 <= Node.val <= 1000

// LevelOrder is a solution to the binary tree level order traversal problem.
func LevelOrder(root *TreeNode) [][]int {
	return levelOrder(root)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	allLevels := [][]int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		level := []int{}
		for range q {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if len(level) > 0 {
			allLevels = append(allLevels, level)
		}
	}
	return allLevels
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
