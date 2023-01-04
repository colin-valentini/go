package groupanagrams

// LeetCode #49.
// https://leetcode.com/problems/group-anagrams/

// Given an array of strings strs, group the anagrams together. You can return
// the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a
// different word or phrase, typically using all the original letters exactly
// once.

// Example 1:
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// Example 2:
// Input: strs = [""]
// Output: [[""]]

// Example 3:
// Input: strs = ["a"]
// Output: [["a"]]

// Constraints:
// - 1 <= strs.length <= 104
// - 0 <= strs[i].length <= 100
// - strs[i] consists of lowercase English letters.

// func groupAnagrams(strs []string) [][]string {
//
// }

type Solver struct {
	strs []string
}

func NewSolver(strs []string) *Solver {
	return &Solver{strs: strs}
}

func (s *Solver) Solve() [][]string {
	if len(s.strs) == 0 {
		return [][]string{{}}
	}
	// Maintain an index from character counts (the unique identifier
	// for an anagram) to the list of strings that are anagram equivalent.
	// Can do this only because fixed-size arrays are "comparable" in Go.
	// (cannot do this with slices).
	m := make(map[charCount]*[]string)
	for _, str := range s.strs {
		counts := newCharCount(str)
		if group, ok := m[counts]; ok {
			*group = append(*group, str)
		} else {
			m[counts] = &([]string{str})
		}
	}
	// Transform the groups into a slice of slice of string.
	groups := make([][]string, 0, len(m))
	for _, group := range m {
		groups = append(groups, *group)
	}
	return groups
}

type charCount [26]uint16

func newCharCount(str string) charCount {
	var counts charCount
	for _, char := range str {
		counts[char-'a']++
	}
	return counts
}
