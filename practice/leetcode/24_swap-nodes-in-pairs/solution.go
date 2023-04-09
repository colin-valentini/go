package swapnodesinpairs

// LeetCode #24.
// https://leetcode.com/problems/swap-nodes-in-pairs/

// Given a linked list, swap every two adjacent nodes and return its head. You
// must solve the problem without modifying the values in the list's nodes
// (i.e., only nodes themselves may be changed.)

// Example 1:
// Input: head = [1,2,3,4]
// Output: [2,1,4,3]

// Example 2:
// Input: head = []
// Output: []

// Example 3:
// Input: head = [1]
// Output: [1]

// ListNode is given by the problem. It defines a singly-linked list of ints.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Solver is a type which can solve the Swap Nodes in Pairs problem.
type Solver struct {
	head *ListNode
}

// NewSolver returns a pointer to a new instance of the Solver type.
func NewSolver(head *ListNode) *Solver {
	return &Solver{head: head}
}

// Solve solves the Swap Nodes in Pairs problem.
//
// Outline:
// - Do a linked list node swap (i -> j becomes j -> i) at every other node.
// - Need to keep track of the previous node in order to swap correctly.
//
// Time: O(N). Space: O(1).
func (s *Solver) Solve() *ListNode {
	// Check that we can even do this once.
	if s.head == nil || s.head.Next == nil {
		return s.head
	}
	head := s.head.Next // Head after algorithm is finished will be head.next.

	prev, i := (*ListNode)(nil), s.head
	for i != nil && i.Next != nil {
		prev, i = s.swap(prev, i, i.Next)
	}
	s.head = head
	return s.head
}

func (s *Solver) swap(prev, i, j *ListNode) (*ListNode, *ListNode) {
	i.Next, j.Next = j, i
	if prev != nil {
		prev.Next = j
	}
	// Return new prev, i
	return i, i.Next
}
