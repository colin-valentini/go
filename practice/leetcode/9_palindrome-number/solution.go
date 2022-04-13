package palindrome

// LeetCode #9.
// https://leetcode.com/problems/palindrome-number/

// Given an integer x, return true if x is palindrome integer.
// An integer is a palindrome when it reads the same backward as forward.
// For example, 121 is a palindrome while 123 is not.

// Example 1:
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

// Example 2:
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-.
// Therefore it is not a palindrome.

// Example 3:
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

// Constraints:
// - -2^31 <= x <= 2^31 - 1

// IsPalindrome is a solution to the "palindrome number" problem.
func IsPalindrome(x int) bool {
	return isPalindrome(x)
}

func isPalindrome(x int) bool {
	// Negative numbers are never palindromes because the negative sign
	// can only be positioned once as the leading character.
	if x < 0 {
		return false
	}
	// Single digit, non-negative numbers are trivial palindromes.
	if x < 10 {
		return true
	}
	// Two digit numbers are only a palindromes if they're a multiple of 11.
	if x < 100 {
		return x%11 == 0
	}
	// If the value in the ones digit is zero, then the only way this number is a
	// palindrome is if the number itself is zero (x can't carry a leading zero),
	// but we checked that case above.
	if x%10 == 0 {
		return false
	}
	// Fast paths are exhausted, reverse the first half of the digits and check
	// for numeric equality against the second half. If the number has an odd
	// number has an odd-number of digits, then after the loop completes x
	// will have one fewer digit than rev, so we can drop the trailing ones
	// and compare.
	rev := 0
	for x > rev {
		rev = rev*10 + (x % 10)
		x /= 10
	}
	return x == rev || x == rev/10
}
