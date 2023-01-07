package dailytemperatures

// LeetCode #739.
// https://leetcode.com/problems/daily-temperatures/

// Given an array of integers temperatures represents the daily temperatures,
// return an array answer such that answer[i] is the number of days you have to
// wait after the ith day to get a warmer temperature. If there is no future day
// for which this is possible, keep answer[i] == 0 instead.

// Example 1:
// Input: temperatures = [73,74,75,71,69,72,76,73]
// Output: [1,1,4,2,1,1,0,0]

// Example 2:
// Input: temperatures = [30,40,50,60]
// Output: [1,1,1,0]

// Example 3:
// Input: temperatures = [30,60,90]
// Output: [1,1,0]

// Constraints:
// - 1 <= temperatures.length <= 10^5
// - 30 <= temperatures[i] <= 100

// func dailyTemperatures(temperatures []int) []int {
//
// }

type Solver struct {
	temps []int
}

func NewSolver(temps []int) *Solver {
	return &Solver{temps: temps}
}

// Solve solves the Daily Temperatures problem.
// Time: O(N). Space: O(N).
func (s *Solver) Solve() any {
	// Key Insight: Build up elements that are "waiting" to find a next-greater
	// temperature in descending order on a stack. That way, when we find an
	// temperature that is greater than the top of our stack, we can find the
	// distance between the two. Keep popping off elements that were "waiting"
	// to find their next greater element.
	res := make([]int, len(s.temps))
	stack := newStack(len(s.temps))
	for i, temp := range s.temps {
		for top, ok := stack.peek(); ok && top.temp < temp; top, ok = stack.peek() {
			stack.pop()
			res[top.pos] = i - top.pos
		}
		stack.push(i, temp)
	}
	return res
}

type stack struct {
	temps []tempPos
}

func newStack(cap int) *stack {
	return &stack{temps: make([]tempPos, 0, cap)}
}

func (s *stack) push(pos, temp int) {
	s.temps = append(s.temps, tempPos{temp: temp, pos: pos})
}

func (s *stack) pop() {
	s.temps = s.temps[:len(s.temps)-1]
}

func (s *stack) peek() (tempPos, bool) {
	if s.len() == 0 {
		return tempPos{}, false
	}
	return s.temps[len(s.temps)-1], true
}

func (s *stack) len() int {
	return len(s.temps)
}

type tempPos struct {
	temp, pos int
}
