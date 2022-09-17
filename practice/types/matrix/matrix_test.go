package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntMatrix2D(t *testing.T) {
	in := [][]int{
		{1, 2, 3},
		{11, 22, 33},
		{111, 222, 333},
		{1111, 2222, 3333},
	}
	matrix := NewMatrix2DFromSlice(in)

	for r := range in {
		for c := range in[r] {
			want := in[r][c]
			got, ok := matrix.Value(r, c)
			require.True(t, ok)
			assert.Equal(t, want, got)

			want = in[r][c] + 666
			require.True(t, matrix.SetValue(r, c, want))

			got, ok = matrix.Value(r, c)
			require.True(t, ok)
			assert.Equal(t, want, got)
		}
	}

	_, ok := matrix.Value(len(in)+1, 0)
	assert.False(t, ok)
	_, ok = matrix.Value(0, len(in[0])+1)
	assert.False(t, ok)

	assert.False(t, matrix.SetValue(len(in)+1, 0, 69420))
	assert.False(t, matrix.SetValue(0, len(in[0])+1, 69240))

	n, m := matrix.Shape()
	assert.Equal(t, len(in), n)
	assert.Equal(t, len(in[0]), m)
}

func TestMatrix2DIterator(t *testing.T) {
	in := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	matrix := NewMatrix2DFromSlice(in)
	it := matrix.Iter()

	want := 0
	assert.True(t, it.HasNext())
	for it.Next() {
		got, ok := it.Value()
		assert.True(t, ok)
		assert.Equal(t, want, got)
		want++
	}
	assert.False(t, it.HasNext())
	assert.False(t, it.Next())
}
