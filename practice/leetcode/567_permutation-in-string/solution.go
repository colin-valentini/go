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

	counts := [26]int{}
	for i := range s.s1 {
		counts[s.s1[i]-'a']++
	}

	left := 0
	for right := range s.s2 {
		counts[s.s2[right]-'a']--
		if counts == [26]int{} {
			// Permutation exists when the counts array is exactly
			// zero everywhere, meaning our counts "balance out"
			// and the current window contains exactly the same
			// characters (and frequency of them) as s1.
			return true
		}
		if right+1 >= len(s.s1) {
			// Extending would create a window of greater size
			// than the permutation we're looking for, so shift
			// left forward, and add back whatever character left
			// pointed to in s2.
			counts[s.s2[left]-'a']++
			left++
		}
	}
	return false
}
