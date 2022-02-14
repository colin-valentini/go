package palindrome

import "math"

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

func IsPalindrome(x int) bool {
	return isPalindrome(x)
}


func isPalindrome(x int) bool {
	// Negative numbers are never palindromes because the negative sign 
	// can only be positioned as the leading character.
    if x < 0 {
        return false
    }
    // Single digit, non-negative numbers are trivial palindromes.
    if x < 10 {
        return true
    }
    // Two digit numbers are only a palindromes if they're a multiple of 11.
    if x < 100 {
        return x % 11 == 0
    }
	// If the value in the ones digit is zero, then the only way this number is a
	// palindrome is if the number itself is zero (x can't carry a leading zero),
	// but we checked that case above.
	if x % 10 == 0 {
		return false
	}
	// Fast paths are exhausted, reverse half the digits and check for equality.
	n := lenDecimalDigits(x)
	return hasSymmetricDigits(x, n)
}

// lenDecimalDigits returns the number of decimal digits in the given number.
func lenDecimalDigits(x int) int {
	return int(math.Log10(float64(x))) + 1
}

// nthPowerOfTen returns 10^N, and caches it for subsequent invocations.
func nthPowerOfTen(n int) int {
	return int(math.Pow(10, float64(n)))
}

// hasSymmetricDigits
func hasSymmetricDigits(x, nd int) bool {
	if nd % 2 == 1 {
		midpower := (nd/2)+1
		midcoeff := nthPowerOfTen(midpower)
		left := (x - (x % midcoeff)) / midcoeff
		right := x % nthPowerOfTen(midpower-1)
		return left == reverseDigits(right, midpower-1)
	}
	midpower := nd/2
	midcoeff := nthPowerOfTen(midpower)
	left := (x - (x % midcoeff)) / midcoeff
	right := x % midcoeff
	return left == reverseDigits(right, midpower)
}

// reverseDigits reverses the digits of x, given that x is nd digits long.
func reverseDigits(x, nd int) int {
    // Iterate over each digit from right to left, and multipling
	// that digit by the power of ten at the position symmetric to it.
    rev := 0
    for i := 1; i <= nd; i++ {
        digit := nthDigit(x, i)
        rev += digit * int(math.Pow(10, float64(nd-i)))
    }
    return rev
}

// nthDigit returns the n-th decimal digit of x using 1-based indexing.
func nthDigit(x, k int) int {
    a := x % nthPowerOfTen(k)
    return (a - (a % nthPowerOfTen(k-1))) / nthPowerOfTen(k-1)
}