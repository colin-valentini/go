package kclosestpointstoorigin

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		k      int
		points [][]int
		want   [][]int
	}{
		{
			k: 0,
			points: [][]int{
				{1, 3},
				{-2, 2},
			},
			want: [][]int{},
		},
		{
			k: 1,
			points: [][]int{
				{1, 3},
				{-2, 2},
			},
			want: [][]int{
				{-2, 2},
			},
		},
		{
			k: 2,
			points: [][]int{
				{17, 23},
				{23, 17},
				{99, 99},
			},
			want: [][]int{
				{17, 23},
				{23, 17},
			},
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.points, testCase.k)
		got := solver.Solve()
		assert.ElementsMatch(t, testCase.want, got, "Failed test case %d", i+1)
	}
}

func TestDistance(t *testing.T) {
	testCases := []struct {
		ax, ay, bx, by int
		want           float64
	}{
		{ax: 0, ay: 0, bx: 0, by: 0, want: 0},
		{ax: 1, ay: 1, bx: 1, by: 1, want: 0},
		{ax: 1, ay: 0, bx: 0, by: 0, want: 1},
		{ax: 0, ay: 1, bx: 0, by: 0, want: 1},
		{ax: 0, ay: 0, bx: 1, by: 0, want: 1},
		{ax: 0, ay: 0, bx: 0, by: 1, want: 1},
		{ax: 1, ay: 1, bx: 0, by: 0, want: math.Sqrt(2)},
		{ax: 1, ay: 0, bx: 0, by: 1, want: math.Sqrt(2)},
		{ax: 0, ay: 1, bx: 1, by: 0, want: math.Sqrt(2)},
		{ax: 1, ay: 3, bx: 0, by: 0, want: math.Sqrt(10)},
		{ax: -2, ay: 2, bx: 0, by: 0, want: math.Sqrt(8)},
	}
	for i, testCase := range testCases {
		got := distance(testCase.ax, testCase.ay, testCase.bx, testCase.by)
		assert.InDelta(t, testCase.want, got, 0.00001, "Failed test case %d", i)
	}
}

func TestSquared(t *testing.T) {
	testCases := []struct {
		x    int
		want float64
	}{
		{x: 0, want: 0},
		{x: 1, want: 1},
		{x: 2, want: 4},
		{x: 3, want: 9},
		{x: 4, want: 16},
		{x: 5, want: 25},
	}
	for i, testCase := range testCases {
		got := squared(testCase.x)
		assert.InDelta(t, testCase.want, got, 0.00001, "Failed test case %d", i)
	}
}
