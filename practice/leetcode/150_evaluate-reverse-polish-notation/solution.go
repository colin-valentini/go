package evaluatereversepolishnotation

import (
	"fmt"
	"strconv"
)

// LeetCode #150.
// https://leetcode.com/problems/evaluate-reverse-polish-notation/

// See http://en.wikipedia.org/wiki/Reverse_Polish_notation.

// You are given an array of strings tokens that represents an arithmetic
// expression in a Reverse Polish Notation.

// Evaluate the expression. Return an integer that represents the value of the
// expression.

// Note that:
// - The valid operators are '+', '-', '*', and '/'.
// - Each operand may be an integer or another expression.
// - The division between two integers always truncates toward zero.
// - There will not be any division by zero.
// - The input represents a valid arithmetic expression in a reverse polish
//   notation.
// - The answer and all the intermediate calculations can be represented in a
//   32-bit integer.

// Example 1:
// Input: tokens = ["2","1","+","3","*"]
// Output: 9
// Explanation: ((2 + 1) * 3) = 9

// Example 2:
// Input: tokens = ["4","13","5","/","+"]
// Output: 6
// Explanation: (4 + (13 / 5)) = 6

// Example 3:
// Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
// Output: 22
// Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
// = ((10 * (6 / (12 * -11))) + 17) + 5
// = ((10 * (6 / -132)) + 17) + 5
// = ((10 * 0) + 17) + 5
// = (0 + 17) + 5
// = 17 + 5
// = 22

// Constraints:
// - 1 <= tokens.length <= 104
// - tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the
//   range [-200, 200].

// func evalRPN(tokens []string) int {
//
// }

type Solver struct {
	tokens []string
	stack  *stack
}

func NewSolver(tokens []string) *Solver {
	return &Solver{tokens: tokens, stack: newStack(len(tokens))}
}

// Solve solves the Evaluate Reverse Polish Notation problem.
// Time: O(N+M), where M = the number of operators.
// Space: O(N).
func (s *Solver) Solve() int {
	// Solution outline:
	// (1) Create a stack to track operands.
	// (2) Iterate through the tokens list.
	//  (a) Push each operand (integer) onto the stack.
	//  (b) When we encounter an operator, pop the stack twice.
	//  (c) Perform the calculation using the popped operands
	//  (d) Push the computed value onto the stack.
	// (3) At the end we should have one element on the stack. Pop and return.
	for _, token := range s.tokens {
		if s.isOperator(token) {
			b, a := s.stack.pop(), s.stack.pop()
			s.stack.push(operator(token).operate(a, b))

		} else {
			operand, _ := strconv.Atoi(token)
			s.stack.push(operand)
		}
	}

	return s.stack.pop()
}

func (s *Solver) isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

const (
	add      = operator("+")
	subtract = operator("-")
	multiply = operator("*")
	divide   = operator("/")
)

type operator string

func (o operator) operate(a, b int) int {
	switch o {
	case add:
		return a + b
	case subtract:
		return a - b
	case multiply:
		return a * b
	case divide:
		return a / b
	default:
		panic(fmt.Sprintf("unsupported operator %q", string(o)))
	}
}

type stack struct {
	operands []int
}

func newStack(cap int) *stack {
	return &stack{operands: make([]int, 0, cap)}
}

func (s *stack) pop() int {
	// panics if called on a stack with no elements.
	operand := s.operands[len(s.operands)-1]
	s.operands = s.operands[:len(s.operands)-1]
	return operand
}

func (s *stack) push(operand int) {
	s.operands = append(s.operands, operand)
}
