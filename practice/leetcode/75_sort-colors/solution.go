package sortcolors

// LeetCode #75.
// https://leetcode.com/problems/sort-colors/

// Given an array nums with n objects colored red, white, or blue, sort them
// in-place so that objects of the same color are adjacent, with the colors in
// the order red, white, and blue.
//
// We will use the integers 0, 1, and 2 to represent the color red, white, and
// blue, respectively.
//
// You must solve this problem without using the library's sort function.
//
// Example 1:
// Input: nums = [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]
//
// Example 2:
// Input: nums = [2,0,1]
// Output: [0,1,2]
//
// Constraints:
// - n == nums.length
// - 1 <= n <= 300
// - nums[i] is either 0, 1, or 2.

// func sortColors(nums []int)  {
//
// }

// Solver is a class which can solve the Sort Colors problem.
type Solver struct {
	elems []int
}

// NewSolver returns a new instance of the Solver type.
func NewSolver(elems []int) *Solver {
	return &Solver{elems: elems}
}

// Solve solves the Sort Colors problem using an arbitrary sorting algorithm.
//
// Selection sort is currently used since I wanted to practice it.
//
// Time: O(N^2) since we have to compare every single possible pair of elements.
// Space: O(1) since we mutate the array in place without auxiliary space.
func (s *Solver) Solve() {
	s.SelectionSort()
}

// SelectionSort is an implementation of the selection sort algorithm.
func (s *Solver) SelectionSort() {
	// At every position in the array, find the minimum in the subarray
	// starting at this position and ending at the last element.
	// Swap the minimum value with the current position.
	for i := range s.elems {
		minIndex := i
		for j := i + 1; j < len(s.elems); j++ {
			if s.less(j, minIndex) {
				minIndex = j
			}
		}
		s.swap(i, minIndex)
	}
}

// Provided as a
func (s *Solver) less(i, j int) bool {
	return s.get(i) < s.get(j)
}

func (s *Solver) get(i int) int {
	// panics on out of bounds index
	return s.elems[i]
}

func (s *Solver) swap(i, j int) {
	s.elems[i], s.elems[j] = s.elems[j], s.elems[i]
}
