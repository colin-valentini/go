package searchinrotatedsortedarray

// LeetCode #33.
// https://leetcode.com/problems/search-in-rotated-sorted-array/

// There is an integer array nums sorted in ascending order (with
// distinct values).

// Prior to being passed to your function, nums is possibly rotated at an
// unknown pivot index k (1 <= k < nums.length) such that the resulting array is
// [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]
// (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3
// and become [4,5,6,7,0,1,2].

// Given the array nums after the possible rotation and an integer target, return
// the index of target if it is in nums, or -1 if it is not in nums.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:
// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4

// Example 2:
// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1

// Example 3:
// Input: nums = [1], target = 0
// Output: -1

// Constraints:
// - 1 <= nums.length <= 5000
// - -10^4 <= nums[i] <= 10^4
// - All values of nums are unique.
// - nums is an ascending array that is possibly rotated.
// - -10^4 <= target <= 10^4

// func search(nums []int, target int) int {
//
// }

// Observations:
//  - Whenever there is a rotation, one half of the array will be
//    out of sort order, but the other will still be sorted.
//  - We can check which half is sorted, and then check whether or not
//    the target _should_ be in the sorted half.
//    - If it should, then we can do vanilla binary search in the sorted half.
//    - If it can't be (we always know by bounds checks), we can continue
//      searching the other half (which may or may not be itself sorted).

type Solver struct {
	nums   []int
	target int
}

func NewSolver(nums []int, target int) *Solver {
	return &Solver{nums: nums, target: target}
}

// Solve solves the Search in Rotated Array problem.
// Time: O(log(N)).
// Space: O(1).
func (s *Solver) Solve() int {
	n := len(s.nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2

		// Always check if we got lucky ;)
		if s.nums[mid] == s.target {
			return mid
		}

		// Do some case-by-case analysis. We know that if there _is_ a rotation
		// in the current window, only one half of the window can be out of
		// sort order, and the other is guaranteed to be in order.
		if s.nums[left] <= s.nums[mid] {
			// Left window is sorted.
			if s.nums[left] <= s.target && s.target <= s.nums[mid] {
				// Target can _only_ be in left (sorted) half.
				return s.binarySearch(left, mid)
			} else {
				// We know the left window is sorted, but target is definitely
				// not contained there, so we have to search the right half.
				left = mid + 1
			}
		} else {
			// Right window is sorted.
			if s.nums[mid] <= s.target && s.target <= s.nums[right] {
				// Target can only be in right (sorted) half.
				return s.binarySearch(mid, right)
			} else {
				// We know the right window is sorted, but target is definitely
				// not contained there, so we have to search the left half.
				right = mid - 1
			}
		}
	}
	return -1
}

func (s *Solver) binarySearch(start, end int) int {
	left, right := start, end
	for left <= right {
		mid := left + (right-left)/2
		if s.target < s.nums[mid] {
			right = mid - 1
		} else if s.nums[mid] < s.target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
