package rotatearray

// LeetCode #189.
// https://leetcode.com/problems/rotate-array/

// Given an integer array nums, rotate the array to the right by k steps, where
// k is non-negative.

// Example 1:
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]

// Example 2:
// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
// Explanation:
// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]

// Constraints:
// - 1 <= nums.length <= 10^5
// - -2^31 <= nums[i] <= 2^31 - 1
// - 0 <= k <= 10^5

// Follow up:
// - Try to come up with as many solutions as you can. There are at least three
//   different ways to solve this problem.
// - Could you do it in-place with O(1) extra space?

// func rotate(nums []int, k int)  {
//
// }

// Solver is a type which can solve the Rotate Array problem.
type Solver struct {
	nums []int
	k    int
}

// NewSolver returns a pointer to a new instance of the Solver type.
func NewSolver(nums []int, k int) *Solver {
	return &Solver{nums: nums, k: k % len(nums)}
}

// Solve solves the Rotate Array problem.
//
// Outline:
//   - We know every element will move k positions ahead (moduluo n).
//   - We can move an element at position i to position (i+k)%n
//   - Just before we move it, we take note of what was at position (i+k)%n
//   - We then move that value (at position j = (i+k)%n) to position (j+k)%n
//   - By doing this we're effectively following the moves one at a time
//
// Caveat: if k evenly divides n we won't "cover" all of the elements by
// simply by starting at index 0, and following the next position.
//
// Time: O(N).
// Space: O(1).
func (s *Solver) Solve() {
	n := len(s.nums)
	if s.k%n == 0 {
		// Rotation by a multiple of the length of the input array
		// is a no-op (rotation results in the same array).
		return
	}

	// moves tracks the number of times we moved an element to its
	// final position. We know that we only need to make n moves.
	moves := 0

	// We can't simply start at 0 and move all of the elements
	// sequentially following their next position. Why? Because
	// if k evenly divides n, then we'll eventually return to 0
	// without moving all of the elements (creating a cycle).
	// We guard against this by wrapping our "move" logic inside
	// another loop that simply starts from the next element.
	//
	// Example (n % k == 0): n = 8, k = 2
	// 0 -> 2 -> 4 -> 6 -> 0 (moves == 4 < n == 8, cycle)
	//
	// If k does not evenly divide n, then it takes exactly n
	// iterations to return to the origin.
	//
	// Example (n % k != 0): n = 8, k = 3
	// 0 -> 3 -> 6 -> 1 -> 4 -> 7 -> 2 -> 5 (moves == 8 == n, done)
	for origin := 0; origin < n; origin++ {
		src, srcVal := origin, s.nums[origin]

		// This loop is the Go equivalent of a do-while loop.
		// We always need to perform at least one iteration.
		for ok := true; ok; ok = src != origin {
			// Move the source value to it's destination which is
			// exactly k elements ahead (moduluo n for wrapping).
			dest := (src + s.k) % n
			destVal := s.nums[dest]
			s.nums[dest] = srcVal
			src, srcVal = dest, destVal
			moves++
			if moves == n {
				// All elements moved to their final position.
				return
			}
		}
	}
}
