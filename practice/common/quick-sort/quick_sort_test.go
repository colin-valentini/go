package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	// Integers.
	in := []any{2, 3, 4, 5, 1}
	QuickSort(in, func(i, j int) bool {
		return in[i].(int) < in[j].(int)
	})
	assert.Equal(t, []any{1, 2, 3, 4, 5}, in, "Sorts ints ascending")
	QuickSort(in, func(i, j int) bool {
		return in[i].(int) > in[j].(int)
	})
	assert.Equal(t, []any{5, 4, 3, 2, 1}, in, "Sorts ints descending")

	// Strings.
	in = []any{"z", "c", "b", "e", "a", "x"}
	QuickSort(in, func(i, j int) bool {
		return in[i].(string) < in[j].(string)
	})
	assert.Equal(t, []any{"a", "b", "c", "e", "x", "z"}, in, "Sorts strings ascending")
	QuickSort(in, func(i, j int) bool {
		return in[i].(string) > in[j].(string)
	})
	assert.Equal(t, []any{"z", "x", "e", "c", "b", "a"}, in, "Sorts strings descending")
}
