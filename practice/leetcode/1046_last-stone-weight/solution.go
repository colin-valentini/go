package laststoneweight

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

// Idea:
// - Create a max heap from the stones list
// - At each turn, we pop two elements from the heap
// - Because the max heap always pops the largest value
//   we know that we have the largest two stones.
// - Push the result of smashing both stones back into the heap, if not both destroyed
// - Repeat above until only one stone is left in the heap.
// - Return either the last element in the heap (if one), or zero.

import "container/heap"

type Solver struct {
	stones []int
}

func NewSolver(stones []int) *Solver {
	return &Solver{stones: stones}
}

func (s *Solver) Solve() int {
	h := MaxIntHeap(s.stones)
	heap.Init(&h)
	for h.Len() > 1 {
		x, y := heap.Pop(&h).(int), heap.Pop(&h).(int)
		min, max := s.minMax(x, y)
		if min == max {
			// Both stones destyoed, nothing to do.
			continue
		}
		// Push the fragmented stone back into the heap.
		heap.Push(&h, max-min)
	}
	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(&h).(int)
}

func (s *Solver) minMax(x, y int) (int, int) {
	if x <= y {
		return x, y
	}
	return y, x
}

type MaxIntHeap []int

func (h MaxIntHeap) Len() int {
	return len(h)
}

func (h MaxIntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxIntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxIntHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}

func (h *MaxIntHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}
