package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToIntegerAtoi(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2147483647, maxInt32)
	assert.Equal(-2147483648, minInt32)

	assert.Equal(42, StringToIntegerAtoi("42"))
	assert.Equal(-42, StringToIntegerAtoi("-42"))
	assert.Equal(-42, StringToIntegerAtoi("     -42"))
	assert.Equal(0, StringToIntegerAtoi("  - 42"))
	assert.Equal(0, StringToIntegerAtoi("  +  413"))
	assert.Equal(0, StringToIntegerAtoi("+-12"))
	assert.Equal(0, StringToIntegerAtoi("  0000000000")) // Another optimization => skip leading zero calculations
	assert.Equal(4193, StringToIntegerAtoi("4193 with words"))
	assert.Equal(0, StringToIntegerAtoi("words and 987"))
	assert.Equal(-2147483648, StringToIntegerAtoi("-91283472332"))
	assert.Equal(2147483647, StringToIntegerAtoi("91283472332"))
	assert.Equal(0, StringToIntegerAtoi(".1"))
	assert.Equal(12345678, StringToIntegerAtoi("  0000000000012345678"))
	assert.Equal(2147483647, StringToIntegerAtoi("20000000000000000000"))
}