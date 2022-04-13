package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, "d", longestPalindrome("d"))
	assert.Equal(t, "bab", LongestPalindrome("babad"))
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
	assert.Equal(t, "bdb", longestPalindrome("abbdb"))
	assert.Equal(t, "abbba", LongestPalindrome("babbbad"))
	assert.Equal(t, "aaa", longestPalindrome("aaadb"))
	assert.Equal(t, "bbbb", longestPalindrome("abbbbc"))
	assert.Equal(t, "ababababa", longestPalindrome("ababababa"))
}
