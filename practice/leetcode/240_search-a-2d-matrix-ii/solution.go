package searcha2dmatrixii

// LeetCode #240.
// https://leetcode.com/problems/search-a-2d-matrix-ii/

// func searchMatrix(matrix [][]int, target int) bool {
//
// }

type Solver struct {
	matrix [][]int
	target int
}

func NewSolver(matrix [][]int, target int) *Solver {
	return &Solver{matrix: matrix, target: target}
}

// Solve solves the Search a 2D Matrix II problem.
// Time: O(N+M).
// Space: O(1).
func (s *Solver) Solve() bool {
	n, m := s.shape()
	// Small optimization: check immediately if the value is out of the
	// min and max range of this matrix. We know the min is the top left
	// cell, and the max is the bottom right cell.
	if n == 0 || m == 0 || s.target < s.min() || s.target > s.max() {
		return false
	}
	// Start from the bottom left corner cell. Move as follows:
	//  - If the cell value is too small, go right one column
	//  - If the cell value is too large, go up one row
	//  - If we go out of bounds, the value is not there.
	for r, c := len(s.matrix)-1, 0; 0 <= r && r < n && 0 <= c && c < m; {
		v := s.matrix[r][c]
		if v < s.target {
			c++
		} else if v > s.target {
			r--
		} else {
			return true
		}
	}
	return false
}

func (s *Solver) shape() (int, int) {
	n := len(s.matrix)
	if n == 0 {
		return n, 0
	}
	return n, len(s.matrix[0])
}

func (s *Solver) min() int {
	return s.matrix[0][0]
}

func (s *Solver) max() int {
	return s.matrix[len(s.matrix)-1][len(s.matrix[0])-1]
}
