package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAlienSorted(t *testing.T) {
	testCases := []struct {
		name   string
		words  []string
		order  string
		expect bool
	}{
		{
			name:   "Example 1",
			words:  []string{"hello", "leetcode"},
			order:  "hlabcdefgijkmnopqrstuvwxyz",
			expect: true,
		},
		{
			name:   "Example 2",
			words:  []string{"word", "world", "row"},
			order:  "worldabcefghijkmnpqstuvxyz",
			expect: false,
		},
		{
			name:   "Example 3",
			words:  []string{"apple", "app"},
			order:  "abcdefghijklmnopqrstuvwxyz",
			expect: false,
		},
		{
			name:   "Empty Case",
			words:  []string{},
			order:  "abcdefghijklmnopqrstuvwxyz",
			expect: true,
		},
		{
			name:   "Trivial Case",
			words:  []string{"aliensdontreadenglish"},
			order:  "abcdefghijklmnopqrstuvwxyz",
			expect: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, IsAlienSorted(tc.words, tc.order))
		})
	}
}

func BenchmarkIsAlienSorted(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = IsAlienSorted([]string{
			"alpha",
			"bravo",
			"bravooo",
			"bravoooa",
			"charlie",
			"charly",
			"delta",
			"deltab",
			"deltabba",
			"deltabbba",
			"hotelllllllllllllllllllllllllllllllllllllllllllllllllllll",
			"hotelllllllllllllllllllllllllllllllllllllllllllllllllllllzzzz",
			"hotelllllllllllllllllllllllllllllllllllllllllllllllllllllzzzza",
		}, "abcdefghijklmnopqrstuvwxyz")
	}
}
