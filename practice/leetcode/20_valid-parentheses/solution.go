package validparentheses

// LeetCode #20.
// https://leetcode.com/problems/valid-parentheses/

// Given a string s containing just the characters '(', ')', '{', '}', '[' and
// ']', determine if the input string is valid.

// An input string is valid if:
// - Open brackets must be closed by the same type of brackets.
// - Open brackets must be closed in the correct order.
// - Every close bracket has a corresponding open bracket of the same type.

// Example 1:
// Input: s = "()"
// Output: true

// Example 2:
// Input: s = "()[]{}"
// Output: true
// Example 3:

// Input: s = "(]"
// Output: false

// Constraints:
// - 1 <= s.length <= 104
// - s consists of parentheses only '()[]{}'.

// func isValid(s string) bool {
//
// }

type Solver struct {
	str   string
	stack *stack
}

func NewSolver(str string) *Solver {
	return &Solver{str: str, stack: newStack()}
}

// Solve solves the "Valid Parentheses" problem.
// Time: O(N), Space: O(N)
func (s *Solver) Solve() bool {
	// Iterate through the given string.
	// (1) Every time we see an opened paren/bracket, push it onto the stack.
	// (2) Every time we see a closer paren/bracket, pop the stack.
	// (3) If the popped value matches the type of (closer) paren/bracket,
	// then it's okay. If not, it's invalid.
	for _, char := range s.str {
		if s.isOpener(char) {
			s.stack.push(char)
		} else {
			opener, ok := s.stack.pop()
			if !ok || !s.matchesParen(opener, char) {
				return false
			}
		}
	}
	// Must be no linger, unclosed parens/brackets.
	return s.stack.len() == 0
}

func (s *Solver) isOpener(char rune) bool {
	return char == '(' || char == '[' || char == '{'
}

func (s *Solver) matchesParen(opener, closer rune) bool {
	switch opener {
	case '(':
		return closer == ')'
	case '[':
		return closer == ']'
	case '{':
		return closer == '}'
	default:
		panic("unexpected left paren or bracket " + string(opener))
	}
}

type stack struct {
	elems []rune
}

func newStack() *stack {
	return &stack{elems: []rune{}}
}

func (s *stack) push(elem rune) {
	s.elems = append(s.elems, elem)
}

func (s *stack) pop() (rune, bool) {
	if len(s.elems) == 0 {
		return rune(0), false
	}
	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return elem, true
}

func (s *stack) len() int {
	return len(s.elems)
}
