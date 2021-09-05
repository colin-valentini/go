package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2147483647, maxInt32)
	assert.Equal(-2147483648, minInt32)

	assert.Equal(42, myAtoi("42"))
	assert.Equal(-42, myAtoi("-42"))
	assert.Equal(-42, myAtoi("     -42"))
	assert.Equal(0, myAtoi("  - 42"))
	assert.Equal(0, myAtoi("  +  413"))
	assert.Equal(0, myAtoi("+-12"))
	assert.Equal(0, myAtoi("  0000000000")) // Another optimization => skip leading zero calculations
	assert.Equal(4193, myAtoi("4193 with words"))
	assert.Equal(0, myAtoi("words and 987"))
	assert.Equal(-2147483648, myAtoi("-91283472332"))
	assert.Equal(2147483647, myAtoi("91283472332"))
	assert.Equal(0, myAtoi(".1"))
	assert.Equal(12345678, myAtoi("  0000000000012345678"))
	assert.Equal(2147483647, myAtoi("20000000000000000000"))
}