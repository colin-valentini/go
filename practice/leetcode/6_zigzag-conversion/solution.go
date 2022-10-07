package leetcode

// LeetCode #6.
// https://leetcode.com/problems/zigzag-conversion/

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
// of rows like this: (you may want to display this pattern in a fixed font for
// better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R

// And then read line by line: "PAHNAPLSIIGYIR"

// Write the code that will take a string and make this conversion given a
// number of rows:

// string convert(string s, int numRows);

// Example 1:
// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"

// Example 2:
// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I

// Example 3:
// Input: s = "A", numRows = 1
// Output: "A"

// Constraints:
// * 1 <= s.length <= 1000
// * s consists of English letters (lower-case and upper-case), ',' and '.'.
// * 1 <= numRows <= 1000

type zigZagConversionSolver struct {
	s       string
	numRows int
	rows    [][]byte
}

func newZigZagConversionSolver(s string, numRows int) *zigZagConversionSolver {
	return &zigZagConversionSolver{s, numRows, nil}
}

// Solve solves the Zigzag Conversion problem.
func (z *zigZagConversionSolver) Solve() string {
	if z.numRows == 1 {
		return z.s
	}

	// Initialize the rows data structure
	if z.rows == nil {
		z.rows = make([][]byte, z.numRows)
	}

	// Use the iterator to get each character in zig-zag pattern.
	it := newZigZagConversionIterator(z)
	for it.next() {
		ridx, sidx, _ := it.get()
		// Initialize a row if we need to
		if z.rows[ridx] == nil {
			z.rows[ridx] = []byte{}
		}
		// Append the next character into this row.
		// NOTE: This only works because the problem is constrained to ASCII.
		z.rows[ridx] = append(z.rows[ridx], z.s[sidx])
	}

	out := ""
	for _, row := range z.rows {
		out += string(row)
	}
	return out
}

type zigZagConversionIterator struct {
	ridx, sidx int
	inc        int
	solver     *zigZagConversionSolver
}

func newZigZagConversionIterator(s *zigZagConversionSolver) *zigZagConversionIterator {
	return &zigZagConversionIterator{-1, -1, 1, s}
}

func (it *zigZagConversionIterator) hasNext() bool {
	return it.sidx+1 < len(it.solver.s)
}

func (it *zigZagConversionIterator) next() bool {
	if !it.hasNext() {
		return false
	}
	it.sidx++
	// Hitting the zero-th row means we need to start ascending again,
	// and similarly getting tothe (numRows-1)-th row means we flip to descent.
	if it.ridx == 0 {
		it.inc = 1
	} else if it.ridx == it.solver.numRows-1 {
		it.inc = -1
	}
	it.ridx += it.inc
	return true
}

func (it *zigZagConversionIterator) get() (int, int, bool) {
	if it.ridx < 0 {
		return 0, 0, false
	}
	return it.ridx, it.sidx, true
}
