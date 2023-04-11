package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	t.Run("Integer slice ascending", func(t *testing.T) {
		in := []any{1, 2, 3, 3, 3, 3, 3, 3, 4, 5, 5, 5, 7, 9}
		assert.Equal(t, 12, BinarySearch(in, intComparator(t, in, 7)))
		assert.Equal(t, 2, BinarySearch(in, intComparator(t, in, 3)))
		assert.Equal(t, -1, BinarySearch(in, intComparator(t, in, 6)))
		assert.Equal(t, -1, BinarySearch(in, intComparator(t, in, 0)))
		assert.Equal(t, -1, BinarySearch(in, intComparator(t, in, 10)))
	})

	t.Run("String slice ascending", func(t *testing.T) {
		in := []any{"b", "b", "c", "d", "d", "d", "e", "f", "h"}
		assert.Equal(t, 0, BinarySearch(in, stringComparator(t, in, "b")))
		assert.Equal(t, 3, BinarySearch(in, stringComparator(t, in, "d")))
		assert.Equal(t, -1, BinarySearch(in, stringComparator(t, in, "g")))
		assert.Equal(t, -1, BinarySearch(in, stringComparator(t, in, "a")))
		assert.Equal(t, -1, BinarySearch(in, stringComparator(t, in, "z")))
	})

	t.Run("Rune slice descending", func(t *testing.T) {
		in := []any{'H', 'G', 'F', 'E', 'D', 'D', 'D', 'B', 'B'}
		assert.Equal(t, 1, BinarySearch(in, runeComparator(t, in, 'G')))
		assert.Equal(t, 4, BinarySearch(in, runeComparator(t, in, 'D')))
		assert.Equal(t, -1, BinarySearch(in, runeComparator(t, in, 'C')))
		assert.Equal(t, -1, BinarySearch(in, runeComparator(t, in, 'I')))
		assert.Equal(t, -1, BinarySearch(in, runeComparator(t, in, 'A')))
	})
}

// intComparator returns a compare function for a slice of int sorted asc.
func intComparator(t *testing.T, in []any, target int) func(i int) int {
	t.Helper()
	return func(i int) int {
		if in[i].(int) == target {
			return 0
		} else if target < in[i].(int) {
			return -1
		} else /* if target > in[i].(int) */ {
			return 1
		}
	}
}

// stringComparator returns a compare function for a slice of string sorted asc.
func stringComparator(t *testing.T, in []any, target string) func(i int) int {
	t.Helper()
	return func(i int) int {
		if in[i].(string) == target {
			return 0
		} else if target < in[i].(string) {
			return -1
		} else /* if target > in[i].(string) */ {
			return 1
		}
	}
}

// runeComparator returns a compare function for a slice of rune sorted desc.
func runeComparator(t *testing.T, in []any, target rune) func(i int) int {
	t.Helper()
	return func(i int) int {
		if in[i].(rune) == target {
			return 0
		} else if target < in[i].(rune) {
			return 1
		} else /* if target > in[i].(rune) */ {
			return -1
		}
	}
}
