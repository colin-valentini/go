package mergesortedarray

// LeetCode #88.
// https://leetcode.com/problems/merge-sorted-array/

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing
// order, and two integers m and n, representing the number of elements in nums1
// and nums2 respectively.
//
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//
// The final sorted array should not be returned by the function, but instead be
// stored inside the array nums1. To accommodate this, nums1 has a length of
// m + n, where the first m elements denote the elements that should be merged,
// and the last n elements are set to 0 and should be ignored. nums2 has a
// length of n.
//
// Example 1:
// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming
// from nums1.
//
// Example 2:
// Input: nums1 = [1], m = 1, nums2 = [], n = 0
// Output: [1]
// Explanation: The arrays we are merging are [1] and [].
// The result of the merge is [1].
//
// Example 3:
// Input: nums1 = [0], m = 0, nums2 = [1], n = 1
// Output: [1]
// Explanation: The arrays we are merging are [] and [1].
// The result of the merge is [1].
// Note that because m = 0, there are no elements in nums1. The 0 is only there
// to ensure the merge result can fit in nums1.
//
// Constraints:
// - nums1.length == m + n
// - nums2.length == n
// - 0 <= m, n <= 200
// - 1 <= m + n <= 200
// - -10^9 <= nums1[i], nums2[j] <= 10^9

// func merge(nums1 []int, m int, nums2 []int, n int)  {
//
// }

// Solver is a type which can solve the Merge Sorted Array problem.
type Solver struct {
	m, n  int
	nums1 []int
	nums2 []int
}

// NewSolver returns a new instance of the Solver type.
func NewSolver(m, n int, nums1, nums2 []int) *Solver {
	return &Solver{m: m, n: n, nums1: nums1, nums2: nums2}
}

// Solve solves the Merge Sorted Array problem.
//
// The solution uses a three pointer strategy, and works in reverse. That is,
// we start from the "back" of both arrays, and push the greater of the two
// pointed-to elements into the rearward-most open position which we keep track
// of using a "write" pointer. If one of the array pointers is exhausted, we
// simply fill the remaining space with elements from the other array.
//
// Example:
// nums1 = [1, 3, 5, 0, 0, 0], nums2 = [2, 4, 6]
// w = 5, i = 2, j = 2: s.nums1[i=2] = 5 < s.nums2[j=2] = 6 -> s.nums[w] = 6, j--
// w = 4, i = 2, j = 1: s.nums1[i=2] = 5 > s.nums2[j=1] = 4 -> s.nums[w] = 5, i--
// w = 3, i = 1, j = 1: s.nums1[i=1] = 3 < s.nums2[j=1] = 4 -> s.nums[w] = 4, j--
// w = 2, i = 1, j = 0: s.nums1[i=1] = 3 > s.nums2[j=0] = 2 -> s.nums[w] = 3, i--
// w = 1, i = 0, j = 0: s.nums1[i=0] = 1 < s.nums2[j=0] = 2 -> s.nums[w] = 2, j--
// w = 0, i = 2, j = -1: -> s.nums[w] = 1, i--
//
// Time: O(N+M) since we have single loop of N+M operations.
// Space: O(1) since we only use constant space for the three pointers.
func (s *Solver) Solve() {
	i, j := s.m-1, s.n-1
	for write := len(s.nums1) - 1; write >= 0; write-- {
		if i >= 0 && j >= 0 {
			if s.nums1[i] > s.nums2[j] {
				s.nums1[write] = s.nums1[i]
				i--
			} else {
				s.nums1[write] = s.nums2[j]
				j--
			}
		} else if i >= 0 {
			s.nums1[write] = s.nums1[i]
			i--
		} else /* if j >= 0*/ {
			s.nums1[write] = s.nums2[j]
			j--
		}
	}
}
