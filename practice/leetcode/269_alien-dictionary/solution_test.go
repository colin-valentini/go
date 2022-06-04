package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlienOrder(t *testing.T) {
	testCases := []struct {
		name   string
		words  []string
		expect string
	}{
		{
			name:   "Example 1",
			words:  []string{"wrt", "wrf", "er", "ett", "rftt"},
			expect: "wertf",
		},
		{
			name:   "Example 2",
			words:  []string{"z", "x"},
			expect: "zx",
		},
		{
			name:   "Example 3",
			words:  []string{"z", "x", "z"},
			expect: "",
		},
		{
			name:   "Empty Case",
			words:  []string{},
			expect: "",
		},
		{
			name:   "Trivial Case",
			words:  []string{"aliensdontreadenglish"},
			expect: "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, AlienOrder(tc.words))
		})
	}
}