package subsets

// LeetCode #78.
// https://leetcode.com/problems/subsets/

// Given an integer array nums of unique elements, return all possible
// subsets (the power set).

// The solution set must not contain duplicate subsets. Return the solution in
// any order.

// Example 1:
// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// Example 2:
// Input: nums = [0]
// Output: [[],[0]]

// Constraints:
// - 1 <= nums.length <= 10
// - -10 <= nums[i] <= 10
// - All the numbers of nums are unique.

// func subsets(nums []int) [][]int {
//
// }

// Solve is the type that can solve the Subsets problem.
type Solver struct {
	nums []int
}

// NewSolver returns a new instance of the Solver type.
func NewSolver(nums []int) *Solver {
	return &Solver{nums: nums}
}

// Solve solves the Subsets problem. We construct the power set by taking
// advantage of the fact that once we have the power set of elements A, B
// we can generate the power set of A, B, C by simply adding C to each of those
// sets and including the original set (this effectively toggles inclusion).
//
// Time: O(N * 2^N). Space: O(2^N).
func (s *Solver) Solve() [][]int {
	// Allocate space for all 2^N subsets.
	sets := make([][]int, 0, 2<<len(s.nums))

	// Seed with the empty set.
	sets = append(sets, []int{})

	// Build up the power set.
	for _, num := range s.nums {
		sets = s.extend(sets, num)
	}
	return sets
}

// extend returns a slice of sets equal to: (1) the original sets
// slice, plus (2) every set from sets after adding the given value.
func (s *Solver) extend(sets [][]int, value int) [][]int {
	for _, set := range sets {
		sets = append(sets, append(s.clone(set), value))
	}
	return sets
}

// cloe returns a copy of the given slice.
func (s *Solver) clone(slice []int) []int {
	cp := make([]int, len(slice))
	copy(cp, slice)
	return cp
}
