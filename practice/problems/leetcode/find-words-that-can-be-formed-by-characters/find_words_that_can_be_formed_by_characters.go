package leetcode

// LeetCode #1160.
// https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/

// You are given an array of strings words and a string chars. A string is good
// if it can be formed by characters from chars (each character can only be used once).
// Return the sum of lengths of all good strings in words.

// Example 1:
// Input: words = ["cat","bt","hat","tree"], chars = "atach"
// Output: 6
// Explanation: The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.

// Example 2:
// Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
// Output: 10
// Explanation: The strings that can be formed are "hello" and "world" so the answer is 5 + 5 = 10.

// Constraints:
// - 1 <= words.length <= 1000
// - 1 <= words[i].length, chars.length <= 100
// - words[i] and chars consist of lowercase English letters.

func CountCharacters(words []string, chars string) int {
	return countCharacters(words, chars)
}

func countCharacters(words []string, chars string) int {
    charCounter := NewCharCounter(chars)
    count := 0
    for _, word := range words {
        if charCounter.copy().canConstruct(word) {
            count += len(word)
        }
    }
    return count
}

type CharCounter map[rune]int

func NewCharCounter(chars string) CharCounter {
    cc := make(CharCounter, len(chars))
    for _, char := range chars {
		// Can do unguarded increment here. See: https://staticcheck.io/docs/checks#S1036.
		cc[char] += 1
    }
    return cc
}

func (cc CharCounter) canConstruct(word string) bool {
	for _, char := range word {
		if remaining, ok := cc.consume(char); !ok || remaining < 0 {
			return false
		}
	}
	return true
}

func (cc CharCounter) copy() CharCounter {
    cp := make(CharCounter, len(cc))
    for k, v := range cc {
        cp[k] = v
    }
    return cp
}

func (cc CharCounter) consume(char rune) (int, bool) {
    _, ok := cc[char]
	if ok {
        cc[char]--
    }
    return cc[char], ok
}
