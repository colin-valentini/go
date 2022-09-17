package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumPassesOfMatrix(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		want   int
	}{
		{
			matrix: [][]int{
				{0, -1, -3, 2, 0},
				{1, -2, -5, -1, -3},
				{3, 0, 0, -4, -1},
			},
			want: 3,
		},
		{
			matrix: [][]int{
				{1},
			},
			want: 0,
		},
		{
			matrix: [][]int{
				{1, 0, 0, -2, -3},
				{-4, -5, -6, -2, -1},
				{0, 0, 0, 0, -1},
				{1, 2, 3, 0, -2},
			},
			want: 7,
		},
		{
			matrix: [][]int{
				{1, 0, 0, -2, -3},
				{-4, -5, -6, -2, -1},
				{0, 0, 0, 0, -1},
				{1, 2, 3, 0, 3},
			},
			want: 4,
		},
		{
			matrix: [][]int{
				{1, 0, 0, -2, -3},
				{-4, -5, -6, -2, -1},
				{0, 0, 0, 0, -1},
				{-1, 0, 3, 0, 3},
			},
			want: -1,
		},
		{
			matrix: [][]int{
				{-1},
			},
			want: -1,
		},
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: 0,
		},
		{
			matrix: [][]int{
				{-1, -9, 0, -1, 0},
				{-9, -4, -5, 4, -8},
				{2, 0, 0, -3, 0},
				{0, -17, -4, 2, -5},
			},
			want: 3,
		},
		{
			matrix: [][]int{
				{-2, -3, -4, -1, -9},
				{-4, -3, -4, -1, -2},
				{-6, -7, -2, -1, -1},
				{0, 0, 0, 0, -3},
				{1, -2, -3, -6, -1},
			},
			want: 12,
		},
		{
			matrix: [][]int{
				{-1, 2, 3},
				{4, 5, 6},
			},
			want: 1,
		},
		{
			matrix: [][]int{
				{-1, 0, 3},
				{0, -5, -6},
			},
			want: -1,
		},
		{
			matrix: [][]int{
				{-1, 0, 3},
				{0, -5, -6},
			},
			want: -1,
		},
		{
			matrix: [][]int{
				{0, 0, -1, -2},
				{-2, -3, 4, -1},
				{-2, -3, 1, -3},
				{-14, -15, 2, 0},
				{0, 0, 0, 0},
				{1, -1, -1, -1},
			},
			want: 3,
		},
		{
			matrix: [][]int{
				{0, 0, -1, -2},
				{-2, -3, 4, -1},
				{-2, -3, 1, -3},
				{-14, -15, 2, 0},
				{0, 0, 0, 0},
				{1, -1, -1, 1},
			},
			want: 2,
		},
		{
			matrix: [][]int{
				{-2, 0, -2, 1},
				{-2, -1, -1, -1},
			},
			want: 5,
		},
	}

	for i, testCase := range testCases {
		got := MinimumPassesOfMatrix(testCase.matrix)
		assert.Equal(t, testCase.want, got, "Failed test case %d", i)
	}
}
