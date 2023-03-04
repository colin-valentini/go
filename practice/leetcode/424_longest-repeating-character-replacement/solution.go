package longestrepeatingcharacterreplacement

// LeetCode #424.
// https://leetcode.com/problems/longest-repeating-character-replacement/

// You are given a string s and an integer k. You can choose any character of
// the string and change it to any other uppercase English character. You can
// perform this operation at most k times.

// Return the length of the longest substring containing the same letter you can
// get after performing the above operations.

// Example 1:
// Input: s = "ABAB", k = 2
// Output: 4
// Explanation: Replace the two 'A's with two 'B's or vice versa.

// Example 2:
// Input: s = "AABABBA", k = 1
// Output: 4
// Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.

// Constraints:
// - 1 <= s.length <= 10^5
// - s consists of only uppercase English letters.
// - 0 <= k <= s.length

// func characterReplacement(s string, k int) int {
//
// }

type Solver struct {
	str string
	k   int
}

func NewSolver(str string, k int) *Solver {
	return &Solver{str: str, k: k}
}

// Solve solves the Longest Repeating Character Replacement problem.
// Time: O(26*N). Space: O(26).
func (s *Solver) Solve() int {
	counts := newCharCounts()
	res := 0
	// For each index position in the string, create a window anchored on the
	// right (end) by the index position. Find a left index position (start)
	// by checking and moving the left anchor incrementally.
	for l, r := 0, 0; r < len(s.str); r++ {
		counts.inc(s.str[r])

		// Make sure this window is valid by checking if
		// we have enough replacements, if not, shift the
		// left pointer forward and try again.
		// Valid iff windowLength - maxCharacterCount <= k.
		for (r-l+1)-counts.maxCount() > s.k {
			counts.dec(s.str[l])
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

type charCounts struct {
	counts [26]int
}

func newCharCounts() *charCounts {
	return &charCounts{counts: [26]int{}}
}

func (c *charCounts) inc(char byte) {
	c.counts[char-'A']++
}

func (c *charCounts) dec(char byte) {
	c.counts[char-'A']--
}

func (c *charCounts) maxCount() int {
	mc := 0
	for _, count := range c.counts {
		mc = max(mc, count)
	}
	return mc
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
