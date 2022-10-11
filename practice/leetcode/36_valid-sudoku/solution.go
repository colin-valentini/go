package leetcode

// LeetCode #36.
// https://leetcode.com/problems/valid-sudoku/

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be
// validated according to the following rules:
// * Each row must contain the digits 1-9 without repetition.
// * Each column must contain the digits 1-9 without repetition.
// * Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9
//   without repetition.

// Note:
// * A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// * Only the filled cells need to be validated according to the mentioned rules.

// Example 1.
// Input: board =
// {
//  {"5","3",".",".","7",".",".",".","."},
//  {"6",".",".","1","9","5",".",".","."},
//  {".","9","8",".",".",".",".","6","."},
//  {"8",".",".",".","6",".",".",".","3"},
//  {"4",".",".","8",".","3",".",".","1"},
//  {"7",".",".",".","2",".",".",".","6"},
//  {".","6",".",".",".",".","2","8","."},
//  {".",".",".","4","1","9",".",".","5"},
//  {".",".",".",".","8",".",".","7","9"},
// }
// Output: true

// Example 2:
// Input: board =
// {
//  {"8","3",".",".","7",".",".",".","."},
//  {"6",".",".","1","9","5",".",".","."},
//  {".","9","8",".",".",".",".","6","."},
//  {"8",".",".",".","6",".",".",".","3"},
//  {"4",".",".","8",".","3",".",".","1"},
//  {"7",".",".",".","2",".",".",".","6"},
//  {".","6",".",".",".",".","2","8","."},
//  {".",".",".","4","1","9",".",".","5"},
//  {".",".",".",".","8",".",".","7","9"},
// }
// Output: false
// Explanation: Same as Example 1, except with the 5 in the top left corner being
// modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

// Constraints:
// * board.length == 9
// * board[i].length == 9
// * board[i][j] is a digit 1-9 or '.'.

// func isValidSudoku(board [][]byte) bool {
//
// }

type validSudokuSolver struct {
	board     [][]byte
	subBoards [][]int
}

func newValidSudokuSolver(board [][]byte) *validSudokuSolver {
	return &validSudokuSolver{board, make([][]int, len(board))}
}

// Solve solves the Valid Sudoku problem.
// The basic premise of the algorithm is to do one nest for-loop pass,
// and check if we ever see duplicates in the current row, column and
// sub-board.
func (s *validSudokuSolver) Solve() bool {
	N := len(s.board)
	for r := 0; r < N; r++ {
		rowCounts, colCounts := make([]int, N), make([]int, N)
		for c := 0; c < N; c++ {
			if rv, ok := s.valueAt(r, c); ok {
				if rowCounts[rv-1] > 0 {
					return false
				}
				rowCounts[rv-1]++
				sub := s.subBoardAt(r, c)
				if sub[rv-1] > 0 {
					return false
				}
				sub[rv-1]++
			}
			if cv, ok := s.valueAt(c, r); ok {
				if colCounts[cv-1] > 0 {
					return false
				}
				colCounts[cv-1]++
			}
		}
	}
	return true
}

func (s *validSudokuSolver) subBoardAt(r, c int) []int {
	// Actually only ever need 3 slices to track the sub-board
	// values since every row spans at most 3 sub-boards and
	// we're moving row-by-row. This is easier to reason about
	// though.
	idx := [][]int{
		{0, 0, 0, 1, 1, 1, 2, 2, 2},
		{0, 0, 0, 1, 1, 1, 2, 2, 2},
		{0, 0, 0, 1, 1, 1, 2, 2, 2},
		{3, 3, 3, 4, 4, 4, 5, 5, 5},
		{3, 3, 3, 4, 4, 4, 5, 5, 5},
		{3, 3, 3, 4, 4, 4, 5, 5, 5},
		{6, 6, 6, 7, 7, 7, 8, 8, 8},
		{6, 6, 6, 7, 7, 7, 8, 8, 8},
		{6, 6, 6, 7, 7, 7, 8, 8, 8},
	}[r][c]
	if s.subBoards[idx] == nil {
		s.subBoards[idx] = make([]int, len(s.board))
	}
	return s.subBoards[idx]
}

func (s *validSudokuSolver) isFilled(r, c int) bool {
	return s.board[r][c] != '.'
}

func (s *validSudokuSolver) valueAt(r, c int) (int, bool) {
	if s.isFilled(r, c) {
		return int(s.board[r][c] - '0'), true
	}
	return 0, false
}
