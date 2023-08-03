package majorityelement

// LeetCode #169.
// https://leetcode.com/problems/majority-element/

// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times. You
// may assume that the majority element always exists in the array.
//
// Example 1:
// Input: nums = [3,2,3]
// Output: 3
//
// Example 2:
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2
//
// Constraints:
// - n == nums.length
// - 1 <= n <= 5 * 10^4
// - -10^9 <= nums[i] <= 10^9

// func majorityElement(nums []int) int {
//
// }

// Solver is a type that can solve the Majority Element problem.
type Solver struct {
	nums []int
}

// NewSolver returns a new instance of the Solver type.
func NewSolver(nums []int) *Solver {
	return &Solver{nums: nums}
}

// Solve solves the Majority Element problem.
//
// This solutions uses the Boyer-Moore Voting algorithm.
// See https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm.
//
// Time: O(N) since we do a single pass on the input array.
// Space: O(1) since we have constant space for two additional variables.
func (s *Solver) Solve() int {
	// Problem constraints guarantee input nums are not empty.
	candidate, count := s.nums[0], 1
	for i := 1; i < len(s.nums); i++ {
		// Once we reach 0, we've seen as many "minority" elements
		// as the candidate "majority" element. So we switch to the
		// next value as the current candidate.
		if count == 0 {
			candidate = s.nums[i]
		}
		// Increment the count whenever we see a value matching the candidate,
		// and decrement it whenever we see a non-matching value.
		if s.nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
