package largestrectangleinhistogram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		heights []int
		want    int
	}{
		{heights: []int{2, 1, 5, 6, 2, 3}, want: 10},
		{heights: []int{2, 4}, want: 4},
		{heights: []int{2, 1, 2}, want: 3},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.heights)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}

func BenchmarkSolver(b *testing.B) {
	// Brute force:   108.6 ns/op  72 B/op   2 allocs/op
	// Smarter brute: 95.47 ns/op  120 B/op  3 allocs/op
	for i := 0; i < b.N; i++ {
		_ = NewSolver([]int{2, 1, 5, 6, 2, 3}).Solve()
	}
}
