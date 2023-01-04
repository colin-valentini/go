package laststoneweight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastStoneWeight(t *testing.T) {
	testCases := []struct {
		stones []int
		want   int
	}{
		{stones: []int{2, 7, 4, 1, 8, 1}, want: 1},
		{stones: []int{1}, want: 1},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.stones)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}

func TestInsertDescOrder(t *testing.T) {
	testCases := []struct {
		elem    int
		s, want []int
	}{
		{
			elem: 4,
			s:    []int{},
			want: []int{4},
		},
		{
			elem: 4,
			s:    []int{5, 3, 2, 1},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			elem: 4,
			s:    []int{3, 2, 1},
			want: []int{4, 3, 2, 1},
		},
		{
			elem: 4,
			s:    []int{7, 6, 5},
			want: []int{7, 6, 5, 4},
		},
	}
	for i, testCase := range testCases {
		insertDescOrder(&testCase.s, testCase.elem)
		assert.Equal(t, testCase.want, testCase.s, "Failed test case %d", i+1)
	}
}
