package searchinsertposition

// LeetCode #35.
// https://leetcode.com/problems/search-insert-position/

// Given a sorted array of distinct integers and a target value, return the
// index if the target is found. If not, return the index where it would be if
// it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:
// Input: nums = [1,3,5,6], target = 5
// Output: 2

// Example 2:
// Input: nums = [1,3,5,6], target = 2
// Output: 1
// Example 3:

// Input: nums = [1,3,5,6], target = 7
// Output: 4

// Constraints:
// - 1 <= nums.length <= 10^4
// - -10^4 <= nums[i] <= 10^4
// - nums contains distinct values sorted in ascending order.
// - -10^4 <= target <= 10^4

// func searchInsert(nums []int, target int) int {
//
// }

type Solver struct {
	nums   []int
	target int
}

func NewSolver(nums []int, target int) *Solver {
	return &Solver{nums: nums, target: target}
}

func (s *Solver) Solve() int {
	n := len(s.nums)

	// Small optimization: if the target is outside the min or max
	// of the array, we know immediately where it would be inserted.
	if s.target < s.nums[0] {
		return 0
	}
	if s.target > s.nums[n-1] {
		return n
	}

	// Binary search for any of the following desired conditions:
	//  (1) mid value is exactly the target value.
	//  (2) target fits between mid value and the value just before mid.
	//  (3) target fits between mid value and the value just after mid.
	for left, right := 0, n-1; left <= right; {
		mid := left + (right-left)/2
		if s.nums[mid] == s.target {
			// Condition (1).
			return mid
		} else if s.target < s.nums[mid] {
			// Condition (2).
			if mid > 0 && s.nums[mid-1] < s.target {
				return mid
			}
			right = mid - 1
		} else if s.nums[mid] < s.target {
			// Conidition (3).
			if mid < n-1 && s.target < s.nums[mid+1] {
				return mid + 1
			}
			left = mid + 1
		}
	}
	// Unreachable.
	return -1
}
