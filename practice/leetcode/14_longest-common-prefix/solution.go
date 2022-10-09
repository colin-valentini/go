package leetcode

// LeetCode #14.
// https://leetcode.com/problems/longest-common-prefix/

// Write a function to find the longest common prefix string amongst an array of
// strings.

// If there is no common prefix, return an empty string "".

// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"

// Example 2:
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

// Constraints:
// * 1 <= strs.length <= 200
// * 0 <= strs[i].length <= 200
// * strs[i] consists of only lowercase English letters.

// func longestCommonPrefix(strs []string) string {
//
// }

type longestCommonPrefixSolver struct {
	strs []string
}

func newLongestCommonPrefixSolver(strs []string) *longestCommonPrefixSolver {
	return &longestCommonPrefixSolver{strs}
}

// Solve solves the Longest Common Prefix problem.
func (s *longestCommonPrefixSolver) Solve() string {
	// Paranoia.
	if len(s.strs) == 0 {
		return ""
	}
	// Arbitrarily choose the first element for convenience.
	// We know the longest common prefix is no greater than any randomly
	// chosen element anyway (by definition).
	elem := s.strs[0]
	for i := 0; i < len(elem); i++ {
		// Check the character at position i against the character at that
		// position in every other string.
		for j := 1; j < len(s.strs); j++ {
			other := s.strs[j]
			// When we find a character that doesn't match, or a string that
			// isn't long enough, we're done.
			if i >= len(other) || other[i] != elem[i] {
				return string(elem[0:i])
			}
		}
	}
	// All strings were the same in this case.
	return elem
}
