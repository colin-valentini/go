package leetcode

// LeetCode #953.
// https://leetcode.com/problems/verifying-an-alien-dictionary/

// In an alien language, surprisingly, they also use English lowercase letters,
// but possibly in a different order. The order of the alphabet is some
// permutation of lowercase letters.

// Given a sequence of words written in the alien language, and the order of the
// alphabet, return true if and only if the given words are sorted
// lexicographically in this alien language.

// Example 1.
// Input: words = ["hello", "leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
// Output: true
// Explanation: As 'h' comes before 'l' in this language, then the sequence is
// sorted.

// Example 2.
// Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
// Output: false
// Explanation: As 'd' comes after 'l' in this language, then
// words[0] > words[1], hence the sequence is unsorted.

// Example 3.
// Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
// Output: false
// Explanation: The first three characters "app" match, and the second string is
// shorter (in size.) According to lexicographical rules "apple" > "app",
// because 'l' > '∅', where '∅' is defined as the blank character which is less
// than any other character (More info).

// Constraints:
// - 1 <= words.length <= 100
// - 1 <= words[i].length <= 20
// - order.length == 26
// - All characters in words[i] and order are English lowercase letters.

// IsAlienSorted is a solution to the verifying an alien dictionary problem.
func IsAlienSorted(words []string, order string) bool {
	return isAlienSorted(words, order)
}

func isAlienSorted(words []string, order string) bool {
	// An early return guarantees we won't allocate for the ordering map
	// in the trivial case.
	if len(words) <= 1 {
		return true
	}
	indexMap := newIndexMap(order)
	for i, j := 0, 1; j < len(words); i, j = i+1, j+1 {
		if !inOrder(i, j, words, indexMap) {
			return false
		}
	}
	return true
}

func inOrder(i, j int, words []string, indexMap []int) bool {
	a, b := words[i], words[j]
	for k := 0; k < len(a) && k < len(b); k++ {
		// Assume that words are all in the given alphabet.
		l, r := indexMap[a[k]-'a'], indexMap[b[k]-'a']
		if l < r {
			return true
		} else if l > r {
			return false
		}
		// Else continue
	}
	return len(a) <= len(b)
}

// newIndexMap returns a slice which functions as a reverse lookup
// between a rune (character) and it's index position in the alphabet
// ordering.
func newIndexMap(order string) []int {
	indexMap := make([]int, len(order))
	for i, r := range order {
		// Shift the runes to the 0,1,...N by "subtracting" byte value of 'a'
		indexMap[r-'a'] = i
	}
	return indexMap
}
