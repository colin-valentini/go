package kokoeatingbananas

// LeetCode #875.
// https://leetcode.com/problems/koko-eating-bananas/

// Koko loves to eat bananas. There are n piles of bananas, the ith pile has
// piles[i] bananas. The guards have gone and will come back in h hours.

// Koko can decide her bananas-per-hour eating speed of k. Each hour, she
// chooses some pile of bananas and eats k bananas from that pile. If the pile
// has less than k bananas, she eats all of them instead and will not eat any
// more bananas during this hour.

// Koko likes to eat slowly but still wants to finish eating all the bananas
// before the guards return.

// Return the minimum integer k such that she can eat all the bananas within h
// hours.

// Example 1:
// Input: piles = [3,6,7,11], h = 8
// Output: 4

// Example 2:
// Input: piles = [30,11,23,4,20], h = 5
// Output: 30

// Example 3:
// Input: piles = [30,11,23,4,20], h = 6
// Output: 23

// Constraints:
// - 1 <= piles.length <= 10^4
// - piles.length <= h <= 10^9
// - 1 <= piles[i] <= 10^9

// func minEatingSpeed(piles []int, h int) int {
//
// }

// Observations:
//  - We can reorder the piles
//  - If h = n, then we know that we have only one hour per pile
//  - If h > n, then we know that we can eat at least one pile over more than
//    one hour
//  - h is like a pool resource, we can dip into it to consume part of (or all)
//    of any given pile.
//  - Upper bound is clearly, k = max(piles), which minimizes the _time_ to eat
//    all of the piles (n hours), but doesn't minimize the rate. There is no
//    benefit of eating bananas faster than k = max(piles).
//  - Want to have time to eat as close as possible to h.
//  - Invariant: sum(ttf(i) for all i) <= h where ttf(i) = piles[i] / k

// Idea: binary search for k between 1 and h.
//  - If the total time-to-finish (TTF) for the midpoint (our k) is
//    not enough time to finish all the piles (ttf(k) > h), then we need
//    a larger k (faster rate), so we can search the upper window.
//  - If the total TTF for the midpoint (k) is fast enough to finish all the
//    piles, then it's a solution, but there could be a better (smaller) k
//    so we can search in the lower window.

// h = 8
//  0  1  2  3
// [3, 6, 7, 11]
//  1  1  1  1   at k = 11
//  1  1  1  2   at k = 10
//  1  1  1  2   at k = 9
//  1  1  1  2   at k = 8
//  1  1  1  2   at k = 7 => ttf = 1+1+1+2 = 5
//  1  1  2  2   at k = 6 => ttf = 1+1+2+2 = 6
//  1  2  2  3   at k = 5 => ttf = 1+2+2+3 = 8
//  1  2  2  3   at k = 4 => ttf = 1+2+2+3 = 8
//  1  2  3  4   at k = 3 (but h = 8, and this would take 10 hours)

// h = 6
// piles = [30, 11, 23, 4, 20]
// max = 30
// k = 15 => [2, 1, 2, 1, 2] => ttf = 2+1+2+1+2 = 8
// k = 16 => [2, 1, 2, 1, 2] => ttf = 8
// ...
// k = 23 => [2, 1, 1, 1, 1] => ttf = 2+1+1+1+1 = 6

type Solver struct {
	piles []int
	h     int
}

func NewSolver(piles []int, h int) *Solver {
	return &Solver{piles: piles, h: h}
}

// Solve solves the Koko Eating Bananas problem.
// Time: O(N + log(N)).
// Space: O(1).
func (s *Solver) Solve() int {
	max := s.maxPile()
	if s.h == len(s.piles) {
		// There are only enough hours to spend one hour per pile
		// so the max pile size is the only possible answer.
		return max
	}

	// Binary search for the min.
	min := max
	left, right := 1, max
	for left <= right {
		k := left + (right-left)/2
		time := s.ttf(k)
		if time > s.h {
			// This choice of k resulted in not enough time to
			// finish all the piles. So we need to eat faster.
			left = k + 1
		} else {
			// This choice of k enabled consuming all piles within
			// the time limit, but there could still be a smaller k.
			if k < min {
				min = k
			}
			right = k - 1
		}

	}
	return min
}

func (s *Solver) maxPile() int {
	max := 0
	for _, size := range s.piles {
		if size > max {
			max = size
		}
	}
	return max
}

func (s *Solver) ttf(k int) int {
	t := 0
	for _, size := range s.piles {
		t += size / k
		if size%k > 0 {
			t++
		}
	}
	return t
}
