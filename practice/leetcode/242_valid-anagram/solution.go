package validanagram

// LeetCode #242.
// https://leetcode.com/problems/valid-anagram/

// Given two strings s and t, return true if t is an anagram of s, and false
// otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a
// different word or phrase, typically using all the original letters exactly
// once.

// Example 1:
// Input: s = "anagram", t = "nagaram"
// Output: true

// Example 2:
// Input: s = "rat", t = "car"
// Output: false

// Constraints:
// - 1 <= s.length, t.length <= 5 * 10^4
// - s and t consist of lowercase English letters.

// Follow up: What if the inputs contain Unicode characters? How would you adapt
// your solution to such a case?

// func isAnagram(s string, t string) bool {
//
// }

type Solver struct {
	// TODO
	s, t string
}

func NewSolver(s, t string) *Solver {
	return &Solver{s: s, t: t}
}

func (s *Solver) Solve() bool {
	// Anagram requires equal length strings as it's a simple re-arrangement.
	if len(s.s) != len(s.t) {
		return false
	}
	// Suffices to check that the value counts (occurrence of each character)
	// of each string is equal.
	return newValueCounts(s.s).equals(newValueCounts(s.t))
}

// Only need 16 bits (unsigned) which supports counts of up to ~65k
// since the problem constraints say that each string is no more than 10_000
// characters long.
type valueCounts map[rune]uint16

func newValueCounts(s string) valueCounts {
	vc := make(valueCounts, len(s))
	for _, char := range s {
		vc[char]++
	}
	return vc
}

func (v valueCounts) equals(other valueCounts) bool {
	if len(v) != len(other) {
		return false
	}
	for char, count := range v {
		if count != other[char] {
			return false
		}
	}
	return true
}
