package removeelement

// LeetCode #27.
// https://leetcode.com/problems/remove-element/

// Given an integer array nums and an integer val, remove all occurrences of val
// in nums in-place. The order of the elements may be changed. Then return the
// number of elements in nums which are not equal to val.
//
// Consider the number of elements in nums which are not equal to val be k, to
// get accepted, you need to do the following things:
// - Change the array nums such that the first k elements of nums contain the
//   elements which are not equal to val. The remaining elements of nums are not
//   important as well as the size of nums.
// - Return k.
//
// Example 1:
// Input: nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2,_,_]
// Explanation: Your function should return k = 2, with the first two elements
// of nums being 2. It does not matter what you leave beyond the returned k
// (hence they are underscores).
//
// Example 2:
// Input: nums = [0,1,2,2,3,0,4,2], val = 2
// Output: 5, nums = [0,1,4,0,3,_,_,_]
// Explanation: Your function should return k = 5, with the first five elements
// of nums containing 0, 0, 1, 3, and 4.
// Note that the five elements can be returned in any order. It does not matter
// what you leave beyond the returned k (hence they are underscores).
//
// Constraints:
// - 0 <= nums.length <= 100
// - 0 <= nums[i] <= 50
// - 0 <= val <= 100

// Solver is a type which can solve the Remove Element problem.
type Solver struct {
	nums []int
	val  int
}

// NewSolver returns a new instance of the Solver type.
func NewSolver(nums []int, val int) *Solver {
	return &Solver{nums: nums, val: val}
}

// Solve solves the Remove Element problem.
//
// Strategy is a two pointer solution. One pointer, w, always tracks
// the next position in the array from the right, that DOES NOT point
// to val (some element on the right that we want on the left). The
// other pointer, k, moves forward from the left, anytime we have a
// value not equal to val. If k does point to val, we swap it with w
// which we know is always not val. The number of times we move k
// forward is the number of non-val elements.
//
// Time: O(N) since our left and right pointers are bounded
// Space: O(1) since we use only constant space for the two pointers.
func (s *Solver) Solve() int {
	// Note that for an empty array, `w = -1`, so we'll never loop
	// and return 0 as desired.
	k, w := 0, len(s.nums)-1
	for k <= w {
		// Write pointer `w`, must always point to non-target from the right.
		if s.nums[w] == s.val {
			w--
			continue
		}
		// When `nums[k] == val``, we know we can exploit this position
		// to store a non-target element from the right (at pointer `w`).
		if s.nums[k] == s.val {
			s.swap(k, w)
		}
		// At this point, `k` points to non-target element, so we
		// increment so that k always points to 1-after a known
		// non-target element (i.e. the size after removal).
		k++
	}
	// Since `k` only moves from the left when we have a confirmed
	// non-target element, the number of times it moves is the size
	// after removal.
	return k
}

func (s *Solver) swap(i, j int) {
	s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
}
