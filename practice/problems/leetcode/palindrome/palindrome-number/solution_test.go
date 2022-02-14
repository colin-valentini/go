package palindrome

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	// Negative numbers should all be non-palindromes
	for x := -1_000; x < 0; x++ {
		assert.False(t, IsPalindrome(x), "negative numbers are not palindromes")
	}
	for x := 0; x < 10; x++ {
		assert.True(t, IsPalindrome(x), "single digit numbers are palindromes")
	}
	for exp, x := 1, 1; x <= 1_111_111; exp, x = exp+1, x+int(math.Pow(10, float64(exp))) {
		for scalar := 1; scalar < 10; scalar++ {
			assert.True(t, IsPalindrome(x * scalar))
		}
	}
	assert.False(t, IsPalindrome(987))
	assert.True(t, IsPalindrome(121))
	assert.True(t, IsPalindrome(4334))
	assert.False(t, IsPalindrome(37985))
}

func BenchmarkIsPalindrome(b *testing.B) {
	x := 333_343_333
	for n := 0; n < b.N; n++ {
		IsPalindrome(x)
	}
} 
