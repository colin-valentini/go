package kthlargestelementinanarray

// LeetCode #215.
// https://leetcode.com/problems/kth-largest-element-in-an-array/

import "math"

// Solver is a type that can solve the Kth Largest Element in an Array problem.
type Solver struct {
	nums []int
	k    int
}

// NewSolver returns a pointer to a new instance of the Solver type.
func NewSolver(nums []int, k int) *Solver {
	return &Solver{nums: nums, k: k}
}

// Solve solves the Kth Largest Element in an Array problem.
//
// The solution uses the Quick Select algorithm which is kind of
// like binary search using Quick Sort's partitioning operations.
//
// Time: O(2*N) in the avgerage case, O(N^2) in the worst case.
// Space: O(1).
func (s *Solver) Solve() int {
	if s.k == 1 {
		// First largest number is the max.
		return s.max()
	}
	n := len(s.nums)
	if s.k == n {
		// nth largest numnber is the min.
		return s.min()
	}
	target := n - s.k
	left, right := 0, n-1
	p := s.partition(left, right)
	for ; p != target; p = s.partition(left, right) {
		// If the pivot index we landed on is to the left of the
		// target position, search (parition) the right half.
		// Otherwise we know it was to the right (see loop break
		// condition) so we search (parition) the left half.
		if p < target {
			left = p + 1
		} else {
			right = p - 1
		}
	}
	return s.nums[p]
}

func (s *Solver) partition(left, right int) int {
	// Set pivot index to left, pivot value to value at right.
	p, pv := left, s.nums[right]
	for i := left; i < right; i++ {
		if s.nums[i] < pv {
			s.swap(i, p)
			p++
		}
	}
	// Move the pivot value to the position pointed to by pivot index.
	s.swap(p, right)

	// Return the index at which the array is now sorted with lesser
	// values to the left, and greater values to the right.
	// This index is in it's final sorted position.
	return p
}

func (s *Solver) swap(i, j int) {
	if i != j {
		s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
	}
}

func (s *Solver) min() int {
	m := math.MaxInt
	for _, v := range s.nums {
		if v < m {
			m = v
		}
	}
	return m
}

func (s *Solver) max() int {
	m := math.MinInt
	for _, v := range s.nums {
		if v > m {
			m = v
		}
	}
	return m
}
