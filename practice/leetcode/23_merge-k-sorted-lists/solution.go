package mergeksortedlists

import "math"

// LeetCode #23.
// https://leetcode.com/problems/merge-k-sorted-lists/

// You are given an array of k linked-lists lists, each linked-list is sorted in
// ascending order.

// Merge all the linked-lists into one sorted linked-list and return it.

// Example 1:
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// merging them into one sorted list:
// 1->1->2->3->4->4->5->6

// Example 2:
// Input: lists = []
// Output: []

// Example 3:
// Input: lists = [[]]
// Output: []

// Constraints:
// * k == lists.length
// * 0 <= k <= 104
// * 0 <= lists[i].length <= 500
// * -104 <= lists[i][j] <= 104
// * lists[i] is sorted in ascending order.
// * The sum of lists[i].length will not exceed 10^4.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  func mergeKLists(lists []*ListNode) *ListNode {
//
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solver struct {
	lists []*ListNode
}

func NewSolver(lists []*ListNode) *Solver {
	return &Solver{lists}
}

// Solve solves the Merge K Sorted Lists problem.
func (s *Solver) Solve() *ListNode {
	if s.lists == nil || len(s.lists) == 0 {
		return nil
	}
	head := &ListNode{}
	node := head
	for {
		minIdx, minVal := -1, math.MaxInt64
		for i, n := range s.lists {
			if n != nil && n.Val < minVal {
				minIdx, minVal = i, n.Val
			}
		}
		if minIdx < 0 {
			break
		}
		node.Next = &ListNode{Val: minVal}
		node = node.Next
		s.lists[minIdx] = s.lists[minIdx].Next
	}
	return head.Next
}
