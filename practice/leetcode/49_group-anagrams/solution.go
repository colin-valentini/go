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
	// (1) Iterate over each string in our list of strings.
	// (2) Create a valueCounts type from the string and check any of the
	// previously unique anagrams. If it matches an existing (canonical) anagram
	// append it into the result array at the same location. Otherwise, add a
	// new array to the result, and append this value into the array.
	anagrams := make([][]string, 0, len(s.strs))
	if len(s.strs) == 0 {
		return anagrams
	}
	// Maintain a map from canonical anagrams (canonical here just means the
	// first occurrence of an anagram) to their inserted position in the
	// resulting array.
	seen := make(map[string]uint16, len(s.strs))
	canonical := make(map[string]uint16, len(s.strs))

	// Store the first string in canonical and seen.
	anagrams = append(anagrams, []string{})
	canonical[s.strs[0]] = uint16(0)
	seen[s.strs[0]] = uint16(0)
	next := uint16(1)

	for _, candidate := range s.strs {
		// Already processed this exact candidate.
		if pos, ok := seen[candidate]; ok {
			anagrams[pos] = append(anagrams[pos], candidate)
			continue
		}
		// Check set of canonical candidates.
		for anagram, pos := range canonical {
			if len(candidate) != len(anagram) {
				continue
			}
			// If they're anagrams, insert them into the same position as
			// the canonical one. Else, create new entry.
			// TODO: This is slow because we have to re-create the value
			// counts for the canonical anagram every time. Figure out
			// how to get a hashable representation and store it in the map.
			if newValueCounts(candidate).equals(newValueCounts(anagram)) {
				anagrams[pos] = append(anagrams[pos], candidate)
				seen[candidate] = pos
				break
			}
		}
		// Found this element by matching with some existing canonical.
		if _, ok := seen[candidate]; ok {
			continue
		}
		anagrams = append(anagrams, []string{candidate})
		canonical[candidate] = next
		seen[candidate] = next
		next++
	}
	return anagrams
}

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
