package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestZigZagConversion(t *testing.T) {
	testCases := []struct {
		s       string
		numRows int
		want    string
	}{
		{s: "PAYPALISHIRING", numRows: 1, want: "PAYPALISHIRING"},
		{s: "PAYPALISHIRING", numRows: 2, want: "PYAIHRNAPLSIIG"},
		{s: "PAYPALISHIRING", numRows: 3, want: "PAHNAPLSIIGYIR"},
		{s: "PAYPALISHIRING", numRows: 4, want: "PINALSIGYAHRPI"},
		{s: "FooBarBazFizzBuzz", numRows: 1, want: "FooBarBazFizzBuzz"},
		{s: "FooBarBazFizzBuzz", numRows: 2, want: "FoaBzizuzoBraFzBz"},
		{s: "FooBarBazFizzBuzz", numRows: 3, want: "FazzzoBraFzBzoBiu"},
		{s: "FooBarBazFizzBuzz", numRows: 4, want: "FBzorazBoaziuzBFz"},
		{s: "FooBarBazFizzBuzz", numRows: 5, want: "FzzoaFzoBiuBrzBaz"},
	}
	for i, testCase := range testCases {
		solver := newZigZagConversionSolver(testCase.s, testCase.numRows)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Test case %d failed", i+1)
	}
}

func TestZigZagConversionIterator(t *testing.T) {
	// In: s=ABCDEFG, numRows=3
	// A   E
	// B D F
	// C   G
	// Out: AEBDFCG
	s, numRows := "ABCDEFG", 3
	solver := newZigZagConversionSolver(s, numRows)
	it := newZigZagConversionIterator(solver)

	// Calling get before advancing even once should return !ok
	_, _, ok := it.get()
	require.False(t, ok)
	require.True(t, it.hasNext())

	rowIterOrder := []int{0, 1, 2, 1, 0, 1, 2}
	for wantSidx, wantRidx := range rowIterOrder {
		assert.True(t, it.next())
		gotRidx, gotSidx, ok := it.get()
		require.True(t, ok)
		assert.Equal(t, wantRidx, gotRidx)
		assert.Equal(t, wantSidx, gotSidx)
	}
}

func BenchmarkZigZagConversion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newZigZagConversionSolver("Apalindromeisaword,phrase,number,"+
			"orothersequenceofunitsthatcanbereadthesamewayineitherdirection,"+
			"withgeneralallowancesforadjustmentstopunctuationandworddividers.",
			4,
		)
	}
}
