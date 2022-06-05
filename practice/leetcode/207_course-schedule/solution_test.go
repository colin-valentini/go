package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFinish(t *testing.T) {
	assert.Equal(t, true, CanFinish(2, [][]int{{1, 0}}), "Example 1")
	assert.Equal(t, false, CanFinish(2, [][]int{{1, 0}, {0, 1}}), "Example 2")
}
