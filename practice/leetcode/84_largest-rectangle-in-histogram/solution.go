package largestrectangleinhistogram

// LeetCode #84.
// https://leetcode.com/problems/largest-rectangle-in-histogram/

// func largestRectangleArea(heights []int) int {
//
// }

type Solver struct {
	heights []int
}

func NewSolver(heights []int) *Solver {
	return &Solver{heights: heights}
}

// Solve solves the Largest Rectangle in Histogram problem.
// Time: O(N).
// Space: O(N).
func (s *Solver) Solve() int {
	max := 0
	stack := newStack(len(s.heights))
	for i, height := range s.heights {
		// Anything already on the stack that is taller than the current height
		// cannot produce a rectangle that covers the current position.
		// Pop all those elements off, compute the area, compare it to our best
		// and then push the current height onto the stack, but consider it's
		// start position to be the start position of the last element popped
		// (the furthest "left") as the start position of this rectangle.
		pos := i
		for rec, ok := stack.peek(); ok && rec.height >= height; rec, ok = stack.peek() {
			area := rec.height * (i - rec.pos)
			if area > max {
				max = area
			}
			pos = rec.pos
			stack.pop()
		}
		stack.push(rectangle{height: height, pos: pos})
	}
	// Any elements that remain on the stack need to be popped and
	// the area from them calculated and compared for better results.
	for top, ok := stack.peek(); ok; top, ok = stack.peek() {
		area := top.height * (len(s.heights) - top.pos)
		if area > max {
			max = area
		}
		stack.pop()
	}
	return max
}

type stack struct {
	recs []rectangle
}

func newStack(cap int) *stack {
	return &stack{recs: make([]rectangle, 0, cap)}
}

func (s *stack) push(rec rectangle) {
	s.recs = append(s.recs, rec)
}

func (s *stack) peek() (rectangle, bool) {
	if len(s.recs) == 0 {
		return rectangle{}, false
	}
	return s.recs[len(s.recs)-1], true
}

func (s *stack) pop() {
	s.recs = s.recs[:len(s.recs)-1]
}

// rectangle represents a rectangle with height `height`, starting from `pos`.
type rectangle struct {
	height, pos int
}
