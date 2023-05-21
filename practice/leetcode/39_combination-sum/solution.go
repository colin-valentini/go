package combinationsum

// LeetCode #39.
// https://leetcode.com/problems/combination-sum/

// Given an array of distinct integers candidates and a target integer target,
// return a list of all unique combinations of candidates where the chosen
// numbers sum to target. You may return the combinations in any order.

// The same number may be chosen from candidates an unlimited number of times.
// Two combinations are unique if the frequency of at least one of the chosen
// numbers is different.

// The test cases are generated such that the number of unique combinations that
// sum up to target is less than 150 combinations for the given input.

// Example 1:
// Input: candidates = [2,3,6,7], target = 7
// Output: [[2,2,3],[7]]
// Explanation:
// 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
// 7 is a candidate, and 7 = 7.
// These are the only two combinations.

// Example 2:
// Input: candidates = [2,3,5], target = 8
// Output: [[2,2,2,2],[2,3,3],[3,5]]

// Example 3:
// Input: candidates = [2], target = 1
// Output: []

// Constraints:
// - 1 <= candidates.length <= 30
// - 2 <= candidates[i] <= 40
// - All elements of candidates are distinct.
// - 1 <= target <= 40

// func combinationSum(candidates []int, target int) [][]int {
//
// }

// Solver is a type which knows how to solve the Combination Sum problem.
type Solver struct {
	candidates   []int
	target       int
	combinations [][]int
}

// NewSolver returns a new instance of the Solver type.
func NewSolver(candidates []int, target int) *Solver {
	return &Solver{candidates: candidates, target: target}
}

// Solve solves the Combination Sum problem.
//
// The solution uses backtracking to build up combinations of candidates that
// sum to the target. The key to the algorithm is that at each search step,
// we branch into two distinct paths:
//  1. We allow considering the current candidate again
//  2. We never consider the current candidate again
//
// We do this by keeping track of the "current" candidate index position, and
// incrementing this pointer in branch (2).
//
// This way, we're guaranteed that any pathway along the decision tree has a
// unique frequency of each candidate.
//
// Time: O(2^target) since recursive with binary branches and the max number of
// candidates in a combination (the height of the tree) is at most `target`.
//
// Space: O(combinations) since we have to return all solution combinations.
func (s *Solver) Solve() any {
	s.init()
	s.search(0, newSearchStep())
	return s.combinations
}

func (s *Solver) init() {
	s.combinations = [][]int{}
}

func (s *Solver) search(i int, step *searchStep) {
	if i >= len(s.candidates) || step.sum > s.target {
		return
	}
	if step.sum == s.target {
		s.combinations = append(s.combinations, s.clone(step.nums))
		return
	}
	s.search(i, step.push(s.candidates[i]))
	s.search(i+1, step.pop())
}

func (s *Solver) clone(slice []int) []int {
	cp := make([]int, len(slice))
	copy(cp, slice)
	return cp
}

type searchStep struct {
	sum  int
	nums []int
}

func newSearchStep() *searchStep {
	return &searchStep{sum: 0, nums: []int{}}
}

func (s *searchStep) push(num int) *searchStep {
	s.sum += num
	s.nums = append(s.nums, num)
	return s
}

func (s *searchStep) pop() *searchStep {
	s.sum -= s.nums[len(s.nums)-1]
	s.nums = s.nums[0 : len(s.nums)-1]
	return s
}
