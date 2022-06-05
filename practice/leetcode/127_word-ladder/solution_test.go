package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLadderLength(t *testing.T) {
	assert.Equal(t, 5, LadderLength("hit", "cog", []string{
		"hot",
		"dot",
		"dog",
		"lot",
		"log",
		"cog",
	}), "Example 1")
	assert.Equal(t, 0, LadderLength("hit", "cog", []string{
		"hot",
		"dot",
		"dog",
		"lot",
		"log",
	}), "Example 2")
	assert.Equal(t, 0, LadderLength("hit", "biz", []string{
		"hot",
		"dot",
		"dog",
		"lot",
		"biz",
	}), "No Solution")
}
