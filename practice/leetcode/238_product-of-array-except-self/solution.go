package productofarrayexceptself

// LeetCode #238.
// https://leetcode.com/problems/product-of-array-except-self/

// Given an integer array nums, return an array answer such that answer[i] is
// equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
// integer.

// You must write an algorithm that runs in O(n) time and without using the
// division operation.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

// Example 2:
// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]

// Constraints:
// - 2 <= nums.length <= 10^5
// - -30 <= nums[i] <= 30
// - The product of any prefix or suffix of nums is guaranteed to fit in a
//   32-bit integer.

// Follow up: Can you solve the problem in O(1) extra space complexity? (The
// output array does not count as extra space for space complexity analysis.)

// func productExceptSelf(nums []int) []int {
//
// }

type Solver struct {
	nums []int
}

func NewSolver(nums []int) *Solver {
	return &Solver{nums: nums}
}

func (s *Solver) Solve() []int {
	// Create two arrays:
	//  (1) running product accumulating from the "left"
	//  (2) running product accumulating from the "right"
	// The final value at each position is the product of the left and
	// right running products. We can actually do this in one pass.
	// [    1,   2,   3,     4]
	// [    _,   1, 1*2, 1*2*3]
	// [2*3*4, 3*4,   4,     _]
	n := len(s.nums)
	res := make([]int, n)
	left, right := make([]int, n), make([]int, n)
	left[0], right[n-1] = 1, 1

	for i := 0; i < n; i++ {
		j := n - i - 1
		if i > 0 {
			left[i] = left[i-1] * s.nums[i-1]
		}
		if j < n-1 {
			right[j] = right[j+1] * s.nums[j+1]
		}
	}
	for i := range res {
		res[i] = left[i] * right[i]
	}

	return res
}
