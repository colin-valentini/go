package containsduplicate

// LeetCode #217.
// https://leetcode.com/problems/contains-duplicate/

// Given an integer array nums, return true if any value appears at least twice
// in the array, and return false if every element is distinct.

// Example 1:
// Input: nums = [1,2,3,1]
// Output: true

// Example 2:
// Input: nums = [1,2,3,4]
// Output: false

// Example 3:
// Input: nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true

// Constraints:
// 1 <= nums.length <= 105
// -109 <= nums[i] <= 109

// func containsDuplicate(nums []int) bool {
//
// }

type Solver struct {
	// TODO
	nums []int
}

func NewSolver(nums []int) *Solver {
	return &Solver{nums: nums}
}

func (s *Solver) Solve() bool {
	// Use a set to maintain knowledge of which values have been encountered
	// already. Iterate over the array. If we find a value that is already in
	// the set, return true. If we get to the end of the array, return false.
	set := make(map[int]nothing, len(s.nums))
	for _, v := range s.nums {
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = nothing{}
	}
	return false
}

type nothing struct{}
