package removeduplicatesfromsortedarrayii

// LeetCode #80.
// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/

// Given an integer array nums sorted in non-decreasing order, remove some
// duplicates in-place such that each unique element appears at most twice. The
// relative order of the elements should be kept the same.
//
// Since it is impossible to change the length of the array in some languages,
// you must instead have the result be placed in the first part of the array
// nums. More formally, if there are k elements after removing the duplicates,
// then the first k elements of nums should hold the final result. It does not
// matter what you leave beyond the first k elements.
//
// Return k after placing the final result in the first k slots of nums.
//
// Do not allocate extra space for another array. You must do this by modifying
// the input array in-place with O(1) extra memory.
//
// Example 1:
// Input: nums = [1,1,1,2,2,3]
// Output: 5, nums = [1,1,2,2,3,_]
// Explanation: Your function should return k = 5, with the first five elements
// of nums being 1, 1, 2, 2 and 3 respectively. It does not matter what you
// leave beyond the returned k (hence they are underscores).
//
// Example 2:
// Input: nums = [0,0,1,1,1,1,2,3,3]
// Output: 7, nums = [0,0,1,1,2,3,3,_,_]
// Explanation: Your function should return k = 7, with the first seven elements
// of nums being 0, 0, 1, 1, 2, 3 and 3 respectively. It does not matter what
// you leave beyond the returned k (hence they are underscores).
//
// Constraints:
// - 1 <= nums.length <= 3 * 10^4
// - -10^4 <= nums[i] <= 10^4
// - nums is sorted in non-decreasing order.

// Solver is a type which can solve the Remove Duplicates from Sorted Array II
// problem.
type Solver struct {
	nums []int
}

// NewSolver returns a new instance of the Solver type.
func NewSolver(nums []int) *Solver {
	return &Solver{nums: nums}
}

// Solve solves the Remove Duplicates from Sorted Array II problem.
//
// Strategy uses a three pointer solution. We use one pointer `i` to always
// track the current element (the most recent unique element). Another pointer
// `k` is constantly anchored to the next write position in the array. Finally,
// we use pointer `j` as our loop variable to "look ahead" and scan the array
// for the next unique element, or the second occurrence of the current element
// at position `i`.
//
// At every position in the array (tracked by `j`), we have two "write"
// conditions:
// 1. The element is brand new (doesn't equal what `i` points to), or
// 2. The element is the second occurrence of what `i` points to
//
// For (1), we set the position at `k` (our "write" pointer to the element at
// `j`, and move `i` to now track this new element).
//
// For (2), we also set our write pointer position to match what `j` points to,
// but don't move `i` (since we're still looking at this element).
//
// Since our write pointer `k` is always one index position ahead of the last
// written value, we can finish by returning `k` to indicate the end of the
// "unique+1" elements.
//
// Time: O(N) since we loop over the array once.
// Space: O(1) since we use constant space for our three pointers.
func (s *Solver) Solve() int {
	i, k := 0, 1
	for j := 1; j < len(s.nums); j++ {
		if s.nums[i] != s.nums[j] {
			// We found a brand new element at position j.
			// Write it out and move i so that it points to this new element.
			// Advance k so that it points to next available position.
			s.nums[k] = s.nums[j]
			i = k
			k++
		} else if s.nums[i] == s.nums[j] && k == i+1 {
			// We found the second occurrence at position j.
			// Write it out but don't move i since this value isn't new.
			// Advance k so that it points to next available position.
			s.nums[k] = s.nums[j]
			k++
		}
	}
	return k
}
