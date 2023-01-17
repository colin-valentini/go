package findminimuminrotatedsortedarray

// LeetCode #153.
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

// Suppose an array of length n sorted in ascending order is rotated between 1
// and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:
// - [4,5,6,7,0,1,2] if it was rotated 4 times.
// - [0,1,2,4,5,6,7] if it was rotated 7 times.

// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results
// in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

// Given the sorted rotated array nums of unique elements, return the minimum
// element of this array.

// You must write an algorithm that runs in O(log n) time.

// Example 1:
// Input: nums = [3,4,5,1,2]
// Output: 1
// Explanation: The original array was [1,2,3,4,5] rotated 3 times.

// Example 2:
// Input: nums = [4,5,6,7,0,1,2]
// Output: 0
// Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

// Example 3:
// Input: nums = [11,13,15,17]
// Output: 11
// Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

// Constraints:
// - n == nums.length
// - 1 <= n <= 5000
// - -5000 <= nums[i] <= 5000
// - All the integers of nums are unique.
// - nums is sorted and rotated between 1 and n times.

// func findMin(nums []int) int {
//
// }

// Observations:
// - Mininum element for a sorted array, rotated 1+ times is always at
//   some position i, where nums[i-1] > nums[i].
// - Said another way, the min element always occurs precisely where the sort
//   order invariant breaks down.

// Q: Can we binary search for this?
//  - If our mid value is X, how do we know which window to search next?
//  - nums[mid-1] < nums[mid] only tells us that mid is NOT the min.
//  - A: Compare against the window extrema (left and right elements).
//  - If nums[mid] > nums[right] then the min lies in the upper window.
//  - If nums[mid] < nums[left] then the min lies in the lower window.

//                                       m
// [0,1,2,4,5,6,7] -rotate 1x-> [7,0,1,2,3,4,5,6] (mid value = 3 < 7) => lower
// [0,1,2,4,5,6,7] -rotate 2x-> [6,7,0,1,2,3,4,5] (mid value = 2 < 6) => lower
// [0,1,2,4,5,6,7] -rotate 2x-> [5,6,7,0,1,2,3,4] (mid value = 1 < 5) => lower
// [0,1,2,4,5,6,7] -rotate 3x-> [4,5,6,7,0,1,2,3] (mid value = 0, < 4, < 3)
// [0,1,2,4,5,6,7] -rotate 4x-> [3,4,5,6,7,0,1,2] (mid value = 7 < 2) => upper
// [0,1,2,4,5,6,7] -rotate 5x-> [2,3,4,5,6,7,0,1] (mid value = 6 < 1) => upper
// [0,1,2,4,5,6,7] -rotate 6x-> [1,2,3,4,5,6,7,0] (mid value = 5 < 0) => upper

type Solver struct {
	nums []int
}

func NewSolver(nums []int) *Solver {
	return &Solver{nums: nums}
}

func (s *Solver) Solve() int {
	n := len(s.nums)
	left, right := 0, n-1

	if s.nums[left] < s.nums[right] {
		// No rotation exists.
		return s.nums[left]
	}

	for left <= right {
		mid := left + (right-left)/2
		v := s.nums[mid]
		if mid > 0 && s.nums[mid-1] > v {
			// The minimum value always occurs precisely at position i
			// where nums[i-1] > nums[i] (where the sort order invariant fails).
			return v
		} else if v > s.nums[right] {
			// If the right value is less than the left value,
			// then this range is out of order somewhere in the
			// right window.
			left = mid + 1
		} else {
			// Two cases:
			//  (1) v < s.nums[left] so left window contains min, or
			//  (2) this window is sorted, and we'll terminate at min value.
			right = mid - 1
		}
	}
	return s.nums[left]
}
