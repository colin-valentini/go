package confusingnumber

import (
	"fmt"
	"strconv"
)

// LeetCode #1056.
// https://leetcode.com/problems/confusing-number/

// func confusingNumber(n int) bool {
//
// }

type Solver struct {
	n int
}

func NewSolver(n int) *Solver {
	return &Solver{n: n}
}

// Solve solves the Confusing Number problem.
// Time: O(N). Space: O(N).
// TODO: Do this using math instead of casting to a string
// which is less efficient.
func (s *Solver) Solve() bool {
	// Iterate through the digits in n
	// (1) If we ever encounter a digit that cannot be rotated, stop immediately.
	// (2) Rotate each digit. This is equivalent to mapping it to the rotated
	// digit, and assigning it in mirror position.
	ns := fmt.Sprint(s.n)
	rotated := make([]rune, len(ns))
	for i, digit := range ns {
		rdigit, ok := s.rotate(digit)
		if !ok {
			return false
		}
		rotated[len(rotated)-i-1] = rdigit
	}
	// Ignore the error. It will always be possible to parse.
	m, _ := strconv.Atoi(string(rotated))
	return m != s.n
}

func (s *Solver) rotate(digit rune) (rune, bool) {
	switch digit {
	case '0':
		return '0', true
	case '1':
		return '1', true
	case '6':
		return '9', true
	case '8':
		return '8', true
	case '9':
		return '6', true
	}
	return rune(0), false
}
