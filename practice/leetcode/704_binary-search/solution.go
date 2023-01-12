package binarysearch

// LeetCode #704.
// https://leetcode.com/problems/binary-search/

// Given an array of integers nums which is sorted in ascending order, and an
// integer target, write a function to search target in nums. If target exists,
// then return its index. Otherwise, return -1.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:
// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4

// Example 2:
// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1

// Constraints:
// - 1 <= nums.length <= 10^4
// - -104 < nums[i], target < 10^4
// - All the integers in nums are unique.
// - nums is sorted in ascending order.

// func search(nums []int, target int) int {
//
// }

type Solver struct {
	nums   []int
	target int
}

func NewSolver(nums []int, target int) *Solver {
	return &Solver{nums: nums, target: target}
}

// Solve solves the Binary Search problem.
// Time: O(N log(N)).
// Spcae: O(1).
func (s *Solver) Solve() int {
	l, r := 0, len(s.nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if s.target < s.nums[mid] {
			r = mid - 1
		} else if s.target > s.nums[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
