package longestconsecutivesequence

// LeetCode #128.
// https://leetcode.com/problems/longest-consecutive-sequence/

// Given an unsorted array of integers nums, return the length of the longest
// consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.

// Example 1:
// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
// Therefore its length is 4.

// Example 2:
// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9

// Constraints:
// - 0 <= nums.length <= 10^5
// - -10^9 <= nums[i] <= 10^9

// func longestConsecutive(nums []int) int {
//
// }

type Solver struct {
	nums    []int
	visited map[int]bool
}

func NewSolver(nums []int) *Solver {
	return &Solver{nums: nums}
}

func (s *Solver) Solve() int {
	// (1) Generate a set of known values from the given collection.
	s.visited = make(map[int]bool, len(s.nums))
	for _, num := range s.nums {
		s.visited[num] = false
	}

	// (2) Fan out from each known value. Starting from any available value
	// we do recursively find the length of the descending, consecutive sequence
	// and the same for the ascending sequence. Tag each of the values we visit
	// along the way as "visited" so that we know we can skip them (starting
	// from any consecutive value always finds and returns the length of the all
	// neighboring consecutive values in the chain, so we never have to do this
	// again.
	max := 0
	for num := range s.visited {
		if m := s.dfs(num); m > max {
			max = m
		}
	}

	return max
}

func (s *Solver) dfs(num int) int {
	if !s.hasNum(num) || s.visited[num] {
		return 0
	}
	s.visited[num] = true
	return s.dfsDec(num-1) + 1 + s.dfsInc(num+1)
}

func (s *Solver) dfsDec(num int) int {
	if !s.hasNum(num) {
		return 0
	}
	s.visited[num] = true
	return s.dfsDec(num-1) + 1
}

func (s *Solver) dfsInc(num int) int {
	if !s.hasNum(num) {
		return 0
	}
	s.visited[num] = true
	return 1 + s.dfsInc(num+1)
}

func (s *Solver) hasNum(num int) bool {
	_, ok := s.visited[num]
	return ok
}
