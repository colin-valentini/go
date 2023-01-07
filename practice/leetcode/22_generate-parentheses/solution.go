package generateparentheses

// LeetCode #22.
// https://leetcode.com/problems/generate-parentheses/

// Given n pairs of parentheses, write a function to generate all combinations
// of well-formed parentheses.

// Example 1:
// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]

// Example 2:
// Input: n = 1
// Output: ["()"]

// Constraints:
// - 1 <= n <= 8

// func generateParenthesis(n int) []string {
//
// }

type Solver struct {
	n      int
	stack  *stack
	combos []string
}

func NewSolver(n int) *Solver {
	// TODO: Make the capacity of res equal to the combinatorial
	// count of valid parentheses (requires some math).
	// For now, we can expect for large N, there are at least twice
	// as many combinations as paren pairs.
	return &Solver{n: n, combos: make([]string, 0, 2*n), stack: newStack(n)}
}

// Solve solves the Generate Parentheses problem.
// Time: O(M), where M = the number of combinations.
// Space: O(2N*M).
func (s *Solver) Solve() []string {
	// Key insight: Keep track of the number of opening and closing parens.
	//  - If we have more openers than closers, then we can add a closer.
	//  - If we've added n openers, we can only close.
	//  - If we've added n openers, and n closers, we're done.
	//
	// Solution: recursively build each permutation by branching using
	// the above conditions to check which type of paren (opener vs. closer)
	// we can add at each point. Use a global stack to keep track of string
	// being built (instead of strings which incurs copy penalty).
	s.recursivelyBuild(0, 0)
	return s.combos
}

func (s *Solver) recursivelyBuild(openerCount, closerCount int) {
	if openerCount == s.n && closerCount == s.n {
		s.combos = append(s.combos, s.stack.String())
		return
	}
	if closerCount < openerCount {
		s.stack.push(')')
		s.recursivelyBuild(openerCount, closerCount+1)
		s.stack.pop()
	}
	if openerCount < s.n {
		s.stack.push('(')
		s.recursivelyBuild(openerCount+1, closerCount)
		s.stack.pop()
	}
}

type stack struct {
	parens []rune
}

func newStack(n int) *stack {
	return &stack{parens: make([]rune, 0, 2*n)}
}

func (s *stack) push(paren rune) {
	s.parens = append(s.parens, paren)
}

func (s *stack) pop() {
	// Panics if called on an empty stack.
	s.parens = s.parens[:len(s.parens)-1]
}

func (s *stack) String() string {
	return string(s.parens)
}
