package linkedlistcycle

// LeetCode #141.
// https://leetcode.com/problems/linked-list-cycle/

// Given head, the head of a linked list, determine if the linked list has a
// cycle in it.
//
// There is a cycle in a linked list if there is some node in the list that can
// be reached again by continuously following the next pointer. Internally, pos
// is used to denote the index of the node that tail's next pointer is connected
// to. Note that pos is not passed as a parameter.
//
// Return true if there is a cycle in the linked list. Otherwise, return false.
//
// Example 1:
// Input: head = [3,2,0,-4], pos = 1
// Output: true
// Explanation: There is a cycle in the linked list, where the tail connects to
// the 1st node (0-indexed).
//
// Example 2:
// Input: head = [1,2], pos = 0
// Output: true
// Explanation: There is a cycle in the linked list, where the tail connects to
// the 0th node.
//
// Example 3:
// Input: head = [1], pos = -1
// Output: false
// Explanation: There is no cycle in the linked list.

// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
// func hasCycle(head *ListNode) bool {
//
// }

// ListNode is the definition for a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Solver is a class which can solve the Linked List Cycle problem.
type Solver struct {
	head *ListNode
}

// NewSolver returns a new instance of the Solver type.
func NewSolver(head *ListNode) *Solver {
	return &Solver{head: head}
}

// Solve solves the Linked List Cycle problem.
//
// The solution uses the "tortoise and hare" approach whereby
// we maintain two pointers: "fast" and "slow". The slow pointer
// moves through the list one node at at time, while the fast
// pointer moves two nodes at a time.
//
// If there is no cycle, the fast pointer will obviously reach
// the end of the list first, and so we're done in O(N) time.
// If there is a cycle, the slow and fast pointers both enter
// the cycle and eventually meet (which clearly cannot happen
// without a cycle since the fast one is moving, well, faster).
//
// Note: the choice of fast taking 2 steps is to minimize the
// work we do (the number of "hops"). The same approach will
// succeed for 3, 4, 5, ... steps per fast move, but takes
// longer.
//
// Space: O(1) since we only ever allocate two pointers.
// Time: O(N) since in the worst case, the last node links back
// to the very first node, and we wait for the slow pointer to
// enter the cycle (at which point the fast pointer will meet it
// after having moved 2N times.
func (s *Solver) Solve() bool {
	slow, fast := s.head, s.head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
