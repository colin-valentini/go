package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMatrix2D(t *testing.T) {
	type Foo struct {
		a int
		b string
		c map[string]int
	}
	// Initialize an empty 2x3 matrix
	n, m := 2, 3
	matrix := NewMatrix2D[Foo](n, m)

	want := Foo{a: 69, b: "bar", c: map[string]int{"baz": 420}}
	assert.True(t, matrix.SetValue(n-1, m-1, want))
	got, ok := matrix.Value(n-1, m-1)
	require.True(t, ok)
	assert.Equal(t, want, got)
}

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
