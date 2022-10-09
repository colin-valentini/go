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
	//Not allowed from the problem constraints, but it gives me paranoia.
	if len(s.strs) == 0 {
		return ""
	}
	prefix := s.strs[0]
	// Pairwise comparison: take strings a, b from strs and compare character
	// by character starting from the left to see what the common prefix is.
	for i := 1; i < len(s.strs); i++ {
		prefix = s.prefixAmong(prefix, s.strs[i])
	}
	return prefix
}

func (s *longestCommonPrefixSolver) prefixAmong(a, b string) string {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	next := make([]byte, 0, minLen)
	for j := 0; j < minLen; j++ {
		if a[j] == b[j] {
			next = append(next, a[j])
		} else {
			break
		}
	}
	return string(next)
}
