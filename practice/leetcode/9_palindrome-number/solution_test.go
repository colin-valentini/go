package leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome(1_001))
	// Negative numbers should all be non-palindromes
	for x := -10_000; x < 0; x++ {
		assert.False(t, IsPalindrome(x), "negative numbers are not palindromes")
	}
	for x := 0; x < 10; x++ {
		assert.True(t, IsPalindrome(x), "single digit numbers are palindromes")
	}
	for exp, x := 1, 1; x <= 1_111_111; exp, x = exp+1, x+int(math.Pow(10, float64(exp))) {
		for scalar := 1; scalar < 10; scalar++ {
			res := IsPalindrome(x * scalar)
			assert.True(t, res)
		}
	}
	assert.False(t, IsPalindrome(987))
	assert.True(t, IsPalindrome(121))
	assert.True(t, IsPalindrome(4_334))
	assert.False(t, IsPalindrome(37_985))
	assert.True(t, IsPalindrome(998_777_666_777_899))
}

func BenchmarkIsPalindrome(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPalindrome(222_333_343_333_222)
	}
}
