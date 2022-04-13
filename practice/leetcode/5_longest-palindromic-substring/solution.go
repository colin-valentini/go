package palindrome

// LeetCode #5.
// https://leetcode.com/problems/longest-palindromic-substring/

// Given a string s, return the longest palindromic substring in s.

// Example 1:
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.

// Example 2:
// Input: s = "cbbd"
// Output: "bb"

// Example 3:
// Input: s = "abbbba"
// Output: "abba"

// Contstraints
// - 1 <= s.length <= 1000
// - s consist of only digits and English letters.

// LongestPalindrome is a solution the "longest palindromic substring" problem.
// It is an "expand around the center"
func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}

// Scan through, considering each element the "pivot" element from
// which we erupt left and right scanning for matches.
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	chars := []rune(s)
	longest := []rune("")
	for i := 0; i < len(chars)-1; i++ {
		start, end := findLongestPalindrome(chars, i)
		palindromeLen := end - start + 1
		if palindromeLen > len(longest) {
			longest = chars[start : end+1]
		}
	}
	return string(longest)
}

// findLongestPalindrome returns the start, end indices (inclusive) of
// the substring range that forms the longest palindrome.
func findLongestPalindrome(chars []rune, center int) (int, int) {
	// Compute the single character origin palindrome.
	i, j := longestPalindromeWithCenter(chars, center, center)

	// If a dual character origin is not possible, we're done.
	if chars[center] != chars[center+1] {
		return i, j
	}

	// Compute the dual character origin palindrome, and return the longer.
	n, m := longestPalindromeWithCenter(chars, center, center+1)
	if m-n > j-i {
		return n, m
	}
	return i, j
}

func longestPalindromeWithCenter(chars []rune, center, centerNext int) (int, int) {
	i, j := center, centerNext
	for i-1 >= 0 && j+1 < len(chars) && chars[i-1] == chars[j+1] {
		i--
		j++
	}
	return i, j
}
