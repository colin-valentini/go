package minimumwindowsubstring

// LeetCode #76.
// https://leetcode.com/problems/minimum-window-substring/

// Given two strings s and t of lengths m and n respectively, return the minimum
// window substring of s such that every character in t (including duplicates)
// is included in the window. If there is no such substring, return the empty
// string "".

// The testcases will be generated such that the answer is unique.

// Example 1:
// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C'
// from string t.

// Example 2:
// Input: s = "a", t = "a"
// Output: "a"
// Explanation: The entire string s is the minimum window.

// Example 3:
// Input: s = "a", t = "aa"
// Output: ""
// Explanation: Both 'a's from t must be included in the window.
// Since the largest window of s only has one 'a', return empty string.

// Constraints:
// - m == s.length
// - n == t.length
// - 1 <= m, n <= 10^5
// - s and t consist of uppercase and lowercase English letters.

// Follow up: Could you find an algorithm that runs in O(m + n) time?

// func minWindow(s string, t string) string {
//
// }

type Solver struct {
	s, t string
}

func NewSolver(s, t string) *Solver {
	return &Solver{s: s, t: t}
}

func (s *Solver) Solve() string {
	if len(s.t) > len(s.s) {
		// No way to produce a window of s that contains all characters
		// at their frequency from t since t contains more characters than s.
		return ""
	}

	// Data structure: hash map in the form of a 26*2 fixed sized array.
	// The keys are the characters (mapped to 0, ..., 52-1 using ASCII value).
	// The values are count (frequency of occurrence in t) of that key.
	// We also need to maintain a queue of characters in the current window,
	// so that we know where to "jump" the left pointer forward to next.
	// When we encounter a character that occurs in t, we add it to the
	// window's queue. When we need to move left forward, we dequeue.
	counts := newCharCounts()
	for i := range s.t {
		counts.inc(s.t[i])
	}

	winCounts := counts.clone()
	hasSubstr := false // Need this because we track best window by left, right.
	bestLeft, bestRight := 0, len(s.s)-1

	// Keep track of the indices we observed that matched a
	// character in t.
	q := newQueue()
	for l, r := 0, 0; r < len(s.s); r++ {
		if counts.has(s.s[r]) {
			q.enqueue(r)
			winCounts.dec(s.s[r])
		} else if q.len() == 0 {
			// queue is only completely empty when our current
			// window has no characters from t in it.
			l++
			continue
		}
		// Need to search subwindows while the current window has
		// all the characters because a subwindow will be smaller.
		for winCounts.hasAll() {
			// Window contains t fully. Take note of it
			// if it's better than the last best window we saw.
			// Then dequeue so that we know where to push left to.
			winSize := r - l + 1
			if winSize < bestRight-bestLeft+1 {
				bestLeft, bestRight = l, r
			}
			hasSubstr = true
			_ = q.dequeue()
			winCounts.inc(s.s[l]) // Add back the left pointer's char.
			if q.len() > 0 {
				l = q.peek()
			} else {
				// Only happens if len(t) == 1
				l++
			}
		}
	}
	if !hasSubstr {
		return ""
	}
	return string(s.s[bestLeft : bestRight+1])
}

type charCounts struct {
	// First 26 characters are A-Z, next 26 are a-z
	counts [52]int
}

func newCharCounts() *charCounts {
	return &charCounts{counts: [52]int{}}
}

func (c *charCounts) inc(char byte) {
	// Note: does not check correctness of argument.
	if char <= 'Z' {
		c.counts[char-'A']++
	} else {
		c.counts[char-'a'+26]++
	}
}

func (c *charCounts) dec(char byte) {
	// Note: does not check correctness of argument.
	if char <= 'Z' {
		c.counts[char-'A']--
	} else {
		c.counts[char-'a'+26]--
	}
}

func (c *charCounts) has(char byte) bool {
	// Note: does not check correctness of argument.
	if char <= 'Z' {
		return c.counts[char-'A'] > 0
	} else {
		return c.counts[char-'a'+26] > 0
	}
}

func (c *charCounts) hasAll() bool {
	for _, count := range c.counts {
		if count > 0 {
			return false
		}
	}
	return true
}

func (c *charCounts) clone() *charCounts {
	cp := c.counts
	return &charCounts{counts: cp}
}

type queue struct {
	elems []int
}

func newQueue() *queue {
	return &queue{elems: []int{}}
}

func (q *queue) enqueue(elem int) {
	q.elems = append(q.elems, elem)
}

func (q *queue) dequeue() int {
	// Panics if length not checked.
	elem := q.elems[0]
	q.elems = q.elems[1:q.len()]
	return elem
}

func (q *queue) peek() int {
	// Panics if length not checked.
	return q.elems[0]
}

func (q *queue) len() int {
	return len(q.elems)
}
