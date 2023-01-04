package topkfrequentelements

// import "sort"

// LeetCode #347.
// https://leetcode.com/problems/top-k-frequent-elements/

// Given an integer array nums and an integer k, return the k most frequent
// elements. You may return the answer in any order.

// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]

// Example 2:
// Input: nums = [1], k = 1
// Output: [1]

// Constraints:
// - 1 <= nums.length <= 10^5
// - -10^4 <= nums[i] <= 10^4
// - k is in the range [1, the number of unique elements in the array].
// - It is guaranteed that the answer is unique.

// Follow up: Your algorithm's time complexity must be better than O(n log n),
// where n is the array's size.

// func topKFrequent(nums []int, k int) []int {
//
// }

type Solver struct {
	nums []int
	k    int
}

func NewSolver(nums []int, k int) *Solver {
	return &Solver{nums: nums, k: k}
}

func (s *Solver) Solve() []int {
	// (1) Generate a mapping from unique number to the occurrence count.
	m := make(map[int]uint32)
	for _, num := range s.nums {
		m[num]++
	}

	// (2) Generate a bucket sorted slice. Each position in the slice
	// corresponds to the collection of unique numbers which occurred
	// that number of times, e.g. counts[3] = [...all ints that occurred 3x].
	max := uint32(0)
	counts := make([][]int, len(s.nums)+1)
	for num, count := range m {
		if counts[count] == nil {
			counts[count] = []int{num}
		} else {
			counts[count] = append(counts[count], num)
		}
		// Small optimization: keep track of the largest count
		// so we can skip straight to the largest value in the
		// counts slice in the next step.
		if count > max {
			max = count
		}
	}

	// (3) Transform our bucket sorted slice to a flat slice of the
	// k-th largest elements. Do so by iterating backwards from the max.
	res := make([]int, 0, s.k)
	for i := max; i > 0 && len(res) < s.k; i-- {
		if len(counts[i]) > 0 {
			res = append(res, counts[i]...)
		}
	}

	return res
}
