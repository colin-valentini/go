package algoexpert

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKnapsackProblem(t *testing.T) {
	testCases := []struct {
		capacity int
		items    [][]int
		want     knapsackProblemSolution
	}{
		{
			capacity: 10,
			items: [][]int{
				{1, 2},
				{4, 3},
				{5, 6},
				{6, 7},
			},
			want: knapsackProblemSolution{10, []int{1, 3}},
		},
		{
			capacity: 11,
			items: [][]int{
				{1, 2},
				{4, 3},
				{5, 6},
				{6, 9},
			},
			want: knapsackProblemSolution{10, []int{0, 1, 2}},
		},
		{
			capacity: 200,
			items: [][]int{
				{465, 100},
				{400, 85},
				{255, 55},
				{350, 45},
				{650, 130},
				{1000, 190},
				{455, 100},
				{100, 25},
				{1200, 190},
				{320, 65},
				{750, 100},
				{50, 45},
				{550, 65},
				{100, 50},
				{600, 70},
				{240, 40},
			},
			want: knapsackProblemSolution{1500, []int{3, 12, 14}},
		},
		{
			capacity: 200,
			items: [][]int{
				{465, 100},
				{400, 85},
				{255, 55},
				{350, 45},
				{650, 130},
				{1000, 190},
				{455, 100},
				{100, 25},
				{1200, 190},
				{320, 65},
				{750, 100},
				{50, 45},
				{550, 65},
				{100, 50},
				{600, 70},
				{255, 40},
			},
			want: knapsackProblemSolution{1505, []int{7, 12, 14, 15}},
		},
		{
			capacity: 100,
			items: [][]int{
				{2, 1},
				{70, 70},
				{30, 30},
				{69, 69},
				{100, 100},
			},
			want: knapsackProblemSolution{101, []int{0, 2, 3}},
		},
		{
			capacity: 100,
			items: [][]int{
				{1, 2},
				{70, 70},
				{30, 30},
				{69, 69},
				{99, 100},
			},
			want: knapsackProblemSolution{100, []int{1, 2}},
		},
		{
			capacity: 0,
			items: [][]int{
				{1, 2},
				{70, 70},
				{30, 30},
				{69, 69},
				{100, 100},
			},
			want: knapsackProblemSolution{0, []int{}},
		},
	}
	for i, testCase := range testCases {
		failMsg := "Failed test case %d: "

		want := testCase.want
		got := KnapsackProblem(testCase.items, testCase.capacity)
		require.Len(t, got, 2, failMsg+"return type", i)

		gotValue, ok := got[0].(int)
		require.True(t, ok, failMsg+"total value type", i)
		assert.Equal(t, want.totalValue, gotValue, failMsg+"total value", i)

		gotIndices, ok := got[1].([]int)
		require.True(t, ok, failMsg+"item indices type", i)
		assert.ElementsMatch(t, want.itemIndices, gotIndices, failMsg+"item indices", i)
	}
}
