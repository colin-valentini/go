package removenthnodefromendoflist

// LeetCode #19.
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

// Given the `head` of a linked list, remove the `nth` node from the end of the
// list and return its head.

// Example 1:
// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]

// Example 2:
// Input: head = [1], n = 1
// Output: []

// Example 3:
// Input: head = [1,2], n = 1
// Output: [1]

// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
//
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solver struct {
	head *ListNode
	n    int
}

func NewSolver(head *ListNode, n int) *Solver {
	return &Solver{head: head, n: n}
}

// Solve solves the Remove Nth Node From End of List problem.
// Time: O(N). Space: O(1).
func (s *Solver) Solve() *ListNode {
	// Use a placeholder node just behind the head.
	// We do this so this in case the head pointer is the one
	// we'll remove.
	preHead := &ListNode{Next: s.head}
	left, right := preHead, preHead

	// Advance the forward pointer so that there are n nodes
	// between it and the back pointer.
	for i := 0; i <= s.n; i++ {
		right = right.Next
	}

	// Advance both pointers until the forward one hits the
	// end of the linked list.
	for right != nil {
		left, right = left.Next, right.Next
	}

	// Because the forward pointer is at the end, and the back
	// pointer is exactly n+1 nodes behind it, we know that
	// the element just after the back pointer is the target.
	target := left.Next
	left.Next = target.Next
	if target == s.head {
		s.head = target.Next
	}
	target.Next = nil
	return s.head
}
