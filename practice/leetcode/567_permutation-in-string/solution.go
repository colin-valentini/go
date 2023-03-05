package permutationinstring

// LeetCode #567.
// https://leetcode.com/problems/permutation-in-string/

type Solver struct {
	s1, s2 string
}

func NewSolver(s1, s2 string) *Solver {
	return &Solver{s1: s1, s2: s2}
}

func (s *Solver) Solve() bool {
	if len(s.s1) > len(s.s2) {
		// No possible way for s2 to contain a substring that
		// is a permutation of s1 since s1 has more characters.
		return false
	}

	counts := newCharCounts()
	for i := 0; i < len(s.s1); i++ {
		counts.inc(s.s1[i])
	}

	for left, right := 0, 0; left < len(s.s2); left++ {
		if right < left {
			right = left
		}
		// Move the right anchor of the window as far right as
		// possible.
		for right < len(s.s2) && counts.has(s.s2[right]) {
			counts.dec(s.s2[right])
			right++
		}
		if counts.isEmpty() {
			return true
		}
		if right > left {
			counts.inc(s.s2[left])
		}
	}
	return false
}

type charCounts struct {
	counts [26]int
	total  int
}

func newCharCounts() *charCounts {
	return &charCounts{counts: [26]int{}, total: 0}
}

func (c *charCounts) inc(char byte) {
	c.counts[char-'a']++
	c.total++
}

func (c *charCounts) dec(char byte) {
	c.counts[char-'a']--
	c.total--
}

func (c *charCounts) has(char byte) bool {
	return c.counts[char-'a'] > 0
}

func (c *charCounts) isEmpty() bool {
	return c.total == 0
}
