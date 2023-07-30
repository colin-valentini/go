package binarytreerightsideview

// LeetCode #199.
// https://leetcode.com/problems/binary-tree-right-side-view/

// Given the root of a binary tree, imagine yourself standing on the right side
// of it, return the values of the nodes you can see ordered from top to bottom.
//
// Example 1:
// Input: root = [1,2,3,null,5,null,4]
//      1
//    /    \
//   2      3
//    \      \
//     5      4
// Output: [1,3,4]
//
// Example 2:
// Input: root = [1,null,3]
//   1
//    \
//     3
// Output: [1,3]
//
// Example 3:
// Input: root = []
// Output: []

// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
//  func rightSideView(root *TreeNode) []int {
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

// Solve solves the Binary Tree Right Side View problem.
//
// Strategy:
//   - Perform a breadth-first search (BFS)
//   - Use two queues to separate the current layer from the next
//   - When we pop a node from the queue, we add it's children to
//     the "next" queue (next layer) in left to right order.
//   - Whenever the current queue (layer) is empty, the last node
//     we popped off is the "right most" node since we pushed all
//     nodes in left to right order.
//
// Time: O(N) since we do a BFS and touch every node once.
// Space: O(N) since, at worst, every node is in the RHS
// (which happens when the tree is a linked-list).
func (s *Solver) Solve() []int {
	if s.root == nil {
		return []int{}
	}
	rhs := []int{}
	q, nq := queue{s.root}, queue{}
	for len(q) > 0 {
		node := q.pop()
		if node.Left != nil {
			nq.push(node.Left)
		}
		if node.Right != nil {
			nq.push(node.Right)
		}
		if len(q) == 0 {
			rhs = append(rhs, node.Val)
			q, nq = nq, q // Move to the next layer by swapping
		}
	}
	return rhs
}

type queue []*TreeNode

func (q *queue) push(node *TreeNode) {
	*q = append(*q, node)
}

func (q *queue) pop() *TreeNode {
	// panics if called on empty queue
	node := (*q)[0]
	*q = (*q)[1:]
	return node
}
