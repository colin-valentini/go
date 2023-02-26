package longestsubstringwithoutrepeatingcharacters

// LeetCode #3.
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Given a string s, find the length of the longest  substring without repeating
// characters.

// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a
// substring.

// Constraints:
// - 0 <= s.length <= 5 * 10^4
// - s consists of English letters, digits, symbols and spaces.

// func lengthOfLongestSubstring(s string) int {
//
// }

type Solver struct {
	str string
}

func NewSolver(s string) *Solver {
	return &Solver{str: s}
}

func (s *Solver) Solve() any {
	// 1. Keep a hash map that maps characters to their
	// index position in s.
	// 2. Iterate through the string s keeping two
	//  pointers, left and right, which point to the start
	//  and end of the current substring.
	// 3. At each position i:
	//   a. If s[i] is not in our hash map, add it
	//      and increment right by one (right = i).
	//   b. If s[i] is in the hash map, check the position
	//      of that character, this tells us which character
	//      to "jump to" (i.e. restart the substring window at).
	// 4. Return the substring of s bounded by left and right.
	//
	// Note: Most of this can be done with naive indexing into s
	// only because the string is ASCII (single byte per character).
	if len(s.str) == 0 {
		return 0
	}
	chars := map[rune]int{}
	left, right := 0, 0
	maxLen := 0
	for i, char := range s.str {
		if pos, ok := chars[char]; ok && pos >= left {
			// Already seen this character in our current window.
			// This concludes whatever window we had, so check
			// if it's better than anything we've seen so far.
			if right-left+1 > maxLen {
				maxLen = right - left + 1
			}
			// Next window starts at the character immeidately after
			// wherever the duplicate was first found.
			left = pos + 1
		}
		right = i
		chars[char] = i
	}
	// Check the final substring window.
	if right-left+1 > maxLen {
		maxLen = right - left + 1
	}
	return maxLen
}
