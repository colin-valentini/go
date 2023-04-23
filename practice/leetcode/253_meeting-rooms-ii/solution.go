package meetingroomsii

// LeetCode #253.
// https://leetcode.com/problems/meeting-rooms-ii/

// Given an array of meeting time intervals `intervals`` where
// `intervals[i] = [start_i, end_i]`, return the minimum number of conference
// rooms required.

// Example 1:
// Input: intervals = [[0,30],[5,10],[15,20]]
// Output: 2

// Example 2:
// Input: intervals = [[7,10],[2,4]]
// Output: 1

// Constraints:
// - 1 <= intervals.length <= 10^4
// - 0 <= start_i < end_i <= 10^6

// func minMeetingRooms(intervals [][]int) int {
//
// }

import (
	"container/heap"
	"sort"
)

type Solver struct {
	intervals [][]int
}

func NewSolver(intervals [][]int) *Solver {
	return &Solver{intervals: intervals}
}

// Solve solves the Meeting Rooms II problem.
//
// This solution first sorts the input intervals by the start time ascending.
// We use a min heap to keep track of rooms that are in use (storing their end
// time).
//
// If we encounter a meeting that is starting after (or at the same time as) a
// meeting at the top of the heap, we can release the meeting at the top of the
// heap. You can think of this like a completed meeting being exited, before
// the next meeting begins in the same room.
//
// The solution value is the peak size of the heap, that is, the maximum number
// of concurrent meetings happening at any given time.
//
// Time: O(N*log(N)). Space: O(N).
func (s *Solver) Solve() int {
	// O(N*log(N)).
	sort.Slice(s.intervals, func(i, j int) bool {
		return s.intervals[i][0] < s.intervals[j][0]
	})

	h := minHeap([]int{})
	heap.Init(&h)
	maxsz := 0

	// O(N*log(N)) in the worst case.
	for _, interval := range s.intervals {
		start, end := interval[0], interval[1]
		if h.Len() > 0 && h.Peek() <= start {
			_ = heap.Pop(&h)
		}
		heap.Push(&h, end)
		maxsz = max(maxsz, h.Len())
	}

	return maxsz
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(v any) {
	*h = append(*h, v.(int))
}

func (h *minHeap) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func (h minHeap) Peek() int {
	return h[0]
}
