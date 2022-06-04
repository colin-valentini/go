package leetcode

// LeetCode #269.
// https://leetcode.com/problems/alien-dictionary/

// There is a new alien language which uses the latin alphabet. However, the 
// order among letters are unknown to you. You receive a list of non-empty words
// from the dictionary, where words are sorted lexicographically by the rules of
// this new language. Derive the order of letters in this language.

// Example 1:
// Input: ["wrt", "wrf", "er", "ett", "rftt"]
// Output: "wertf"

// Example 2:
// Input: ["z", "x"]
// Output: "zx"

// Example 3:
// Input: ["z", "x", "z"]
// Output: "" 
// Explanation: The order is invalid, so return "".

// Constraints:
// - All letters are in English, lowercase.
// - If a is a prefix of b, then a must appear before b in the given dictionary.
// - If the order is invalid, return an empty string.
// - If there may be multiple valid order of letters, return any one of them.

func AlienOrder(words []string) string {
	return alienOrder(words)
}

func alienOrder(words []string) string {
	return ""
}