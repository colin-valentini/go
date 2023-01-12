package validpalindrome

import "unicode"

// LeetCode #125.
// https://leetcode.com/problems/valid-palindrome/

// A phrase is a palindrome if, after converting all uppercase letters into
// lowercase letters and removing all non-alphanumeric characters, it reads the
// same forward and backward. Alphanumeric characters include letters and
// numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

// Example 2:
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.

// Example 3:
// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.

// Constraints:
// -1 <= s.length <= 2 * 10^5
// - s consists only of printable ASCII characters.

// func isPalindrome(s string) bool {
//
// }

type Solver struct {
	str string
}

func NewSolver(str string) *Solver {
	return &Solver{str: str}
}

// Solve solves the Valid Palindrome problem.
// Time: O(N).
// Space: O(1).
func (s *Solver) Solve() bool {
	l, r := 0, len(s.str)-1
	for l <= r {
		lv := s.toLower(s.str[l])
		if !s.isAlphaNumeric(lv) {
			l++
			continue
		}
		rv := s.toLower(s.str[r])
		if !s.isAlphaNumeric(rv) {
			r--
			continue
		}
		if lv != rv {
			return false
		}
		l++
		r--
	}
	return true
}

func (s *Solver) isAlphaNumeric(char rune) bool {
	return ('a' <= char && char <= 'z') || ('0' <= char && char <= '9')
}

func (s *Solver) toLower(char byte) rune {
	return unicode.ToLower(rune(char))
}
