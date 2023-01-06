package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		str  string
		want bool
	}{
		{str: "()", want: true},
		{str: "()[]{}", want: true},
		{str: "(]", want: false},
		{str: "((", want: false},
		{str: "[)]", want: false},
		{str: "([()])", want: true},
		{str: "", want: true},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.str)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}

func TestStack(t *testing.T) {
	s := newStack()
	s.push('a')
	a, ok := s.pop()
	require.True(t, ok)
	assert.Equal(t, 'a', a)

	s.push('b')
	s.push('c')
	s.push('d')
	s.push('e')
	assert.Equal(t, 4, s.len())

	e, ok := s.pop()
	require.True(t, ok)
	assert.Equal(t, 'e', e)
	assert.Equal(t, 3, s.len())

	s.push('f')
	assert.Equal(t, 4, s.len())
	f, ok := s.pop()
	require.True(t, ok)
	assert.Equal(t, 'f', f)
	assert.Equal(t, 3, s.len())

	d, ok := s.pop()
	require.True(t, ok)
	assert.Equal(t, 'd', d)
	assert.Equal(t, 2, s.len())

	c, ok := s.pop()
	require.True(t, ok)
	assert.Equal(t, 'c', c)
	assert.Equal(t, 1, s.len())

	b, ok := s.pop()
	require.True(t, ok)
	assert.Equal(t, 'b', b)
	assert.Equal(t, 0, s.len())

	_, ok = s.pop()
	assert.False(t, ok)
	assert.Equal(t, 0, s.len())
}
