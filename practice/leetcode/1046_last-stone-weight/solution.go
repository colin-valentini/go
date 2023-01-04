package laststoneweight

import "sort"

// LeetCode #1046.
// https://leetcode.com/problems/last-stone-weight/

// You are given an array of integers stones where stones[i] is the weight of
// the ith stone.

// We are playing a game with the stones. On each turn, we choose the heaviest
// two stones and smash them together. Suppose the heaviest two stones have
// weights x and y with x <= y. The result of this smash is:
// * If x == y, both stones are destroyed, and
// * If x != y, the stone of weight x is destroyed, and the stone of weight y
//   has new weight y - x.

// At the end of the game, there is **at most one** stone left.

// Return the weight of the last remaining stone. If there are no stones left,
// return 0.

// func lastStoneWeight(stones []int) int {
//
// }

type Solver struct {
	stones []int
}

func NewSolver(stones []int) *Solver {
	return &Solver{stones: stones}
}

func (s *Solver) Solve() int {
	// Sort the slice descending: O(N * log(N))
	sort.Slice(s.stones, func(i, j int) bool {
		return s.stones[i] > s.stones[j]
	})

	for len(s.stones) > 1 {
		y, x := s.stones[0], s.stones[1]
		s.stones = s.stones[2:]
		if x < y {
			insertDescOrder(&(s.stones), y-x)
		}
	}
	if len(s.stones) == 0 {
		return 0
	}
	return s.stones[0]
}

func insertDescOrder(s *[]int, elem int) {
	// TODO: Binary-search-ify this
	for i, v := range *s {
		if elem > v {
			*s = append((*s)[:i+1], (*s)[i:]...)
			(*s)[i] = elem
			return
		}
	}
	*s = append(*s, elem)
}
