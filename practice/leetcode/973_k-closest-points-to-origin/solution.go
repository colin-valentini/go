package kclosestpointstoorigin

// LeetCode #973.
// https://leetcode.com/problems/k-closest-points-to-origin/

// Given an array of points where points[i] = [xi, yi] represents a point on
// the X-Y plane and an integer k, return the k closest points to the origin
// (0, 0).

// The distance between two points on the X-Y plane is the Euclidean distance
// (i.e., √(x1 - x2)2 + (y1 - y2)2).

// You may return the answer in any order. The answer is guaranteed to be unique
// (except for the order that it is in).

// Example 1:
// Input: points = [[1,3],[-2,2]], k = 1
// Output: [[-2,2]]
// Explanation:
// The distance between (1, 3) and the origin is sqrt(10).
// The distance between (-2, 2) and the origin is sqrt(8).
// Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
// We only want the closest k = 1 points from the origin, so the answer is just
// [[-2,2]].

// Example 2:
// Input: points = [[3,3],[5,-1],[-2,4]], k = 2
// Output: [[3,3],[-2,4]]
// Explanation: The answer [[-2,4],[3,3]] would also be accepted.

// Constraints:
// - 1 <= k <= points.length <= 10^4
// - -10^4 < xi, yi < 10^4

// func kClosest(points [][]int, k int) [][]int {
//
// }

import (
	"container/heap"
	"math"
)

type Solver struct {
	k      int
	points [][]int
}

func NewSolver(points [][]int, k int) *Solver {
	return &Solver{points: points, k: k}
}

// Solve solves the K Closest Points to Origin problem.
//
// This solution creates a min heap of points in the plane.
// It then pops off enough elements to fulfill k (or less than k if
// we were provided with len(points) = N < k). It relies on the property
// that all min heaps pop the minimum value.
//
// Time: O(N+k*log(N)), Space: O(N).
// The reason the time complexity is linear in N is because heap construction
// is linear (see https://stackoverflow.com/a/18742428). We also need to pop
// at most k elements, which is a log(N) heap operation.
//
// Note: there are better solutions using quick select (see
// https://leetcode.com/problems/k-closest-points-to-origin/solutions/220235/java-three-solutions-to-this-classical-k-th-problem/?orderBy=most_votes)
func (s *Solver) Solve() [][]int {
	if s.k <= 0 {
		return [][]int{}
	}
	h := pointHeap{}
	for _, p := range s.points {
		h.Push(newPoint(p[0], p[1]))
	}
	heap.Init(&h)
	res := make([][]int, 0, s.k)
	for h.Len() > 0 && len(res) < s.k {
		p := heap.Pop(&h).(point)
		res = append(res, []int{p.x, p.y})
	}
	return res
}

type point struct {
	x, y int
	dist float64 // distance from the origin
}

func newPoint(x, y int) point {
	return point{x: x, y: y, dist: distance(x, y, 0, 0)}
}

func distance(ax, ay, bx, by int) float64 {
	// Don't strictly need the square root here because we don't return it.
	// However, it's easier to understand, and more realistic if we do :)
	return math.Sqrt(squared(bx-ax) + squared(by-ay))
}

func squared(x int) float64 {
	return math.Pow(float64(x), 2)
}

type pointHeap []point

func (h pointHeap) Len() int {
	return len(h)
}

func (h pointHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h pointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pointHeap) Push(v any) {
	*h = append(*h, v.(point))
}

func (h *pointHeap) Pop() any {
	v := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return v
}
