package searcha2dmatrix

// LeetCode #74.
// https://leetcode.com/problems/search-a-2d-matrix/

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

func (s *Solver) Solve() bool {
	// Do a binary search across the rows to find the row
	// where the target should be. Then do a binary search
	// within that row to check if the target is in that row.
	if row := s.findRow(); row >= 0 {
		return s.findCol(row) >= 0
	}
	return false
}

func (s *Solver) findRow() int {
	l, r := 0, len(s.matrix)-1
	for l <= r {
		mid := l + (r-l)/2
		if s.target < s.matrix[mid][0] {
			r = mid - 1
		} else if s.target > s.matrix[mid][len(s.matrix[mid])-1] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func (s *Solver) findCol(row int) int {
	l, r := 0, len(s.matrix[row])-1
	for l <= r {
		mid := l + (r-l)/2
		if s.target < s.matrix[row][mid] {
			r = mid - 1
		} else if s.target > s.matrix[row][mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
