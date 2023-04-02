package kthlargestelementinastream

// LeetCode #703.
// https://leetcode.com/problems/kth-largest-element-in-a-stream/

// Design a class to find the kth largest element in a stream. Note that it is
// the kth largest element in the sorted order, not the kth distinct element.

// Implement KthLargest class:
// - KthLargest(int k, int[] nums) Initializes the object with the integer k and
//   the stream of integers nums.
// - int add(int val) Appends the integer val to the stream and returns the
//   element representing the kth largest element in the stream.

// Example 1:
// Input
// ["KthLargest", "add", "add", "add", "add", "add"]
// [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
// Output
// [null, 4, 5, 5, 8, 8]

// Explanation
// KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
// kthLargest.add(3);   // return 4
// kthLargest.add(5);   // return 5
// kthLargest.add(10);  // return 5
// kthLargest.add(9);   // return 8
// kthLargest.add(4);   // return 8

// Constraints:
// - 1 <= k <= 10^4
// - 0 <= nums.length <= 10^4
// - -10^4 <= nums[i] <= 10^4
// - -10^4 <= val <= 10^4
// - At most 10^4 calls will be made to add.
// - It is guaranteed that there will be at least k elements in the array when you
// - search for the kth element.

// type KthLargest struct {
//
// }

// func Constructor(k int, nums []int) KthLargest {
//
// }

// func (this *KthLargest) Add(val int) int {
//
// }

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

// 5, 6, 7, 8, 9, k=3
// add(4) -> 4, 5, 6, 7, 8, 9 -> return 7
// add(3) -> 3, 4, 5, 6, 7, 8, 9 -> return 7
// add(10) -> 3, 4, 5, 6, 7, 8, 9, 10 -> return 8

import (
	"container/heap"
)

type Solver struct {
	k int
	h *MinIntHeap
}

func NewSolver(k int, nums []int) *Solver {
	h := MinIntHeap(nums)
	heap.Init(&h)
	for h.Len() > k {
		heap.Pop(&h)
	}
	return &Solver{k: k, h: &h}
}

func (s *Solver) Add(val int) int {
	heap.Push(s.h, val)
	if s.h.Len() > s.k {
		heap.Pop(s.h)
	}
	return (*s.h)[0]
}

func (s *Solver) Solve() any {
	return nil
}

type MinIntHeap []int

func (h MinIntHeap) Len() int {
	return len(h)
}

func (h MinIntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinIntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinIntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}
