package sudokusolver

// LeetCode #37.
// https://leetcode.com/problems/sudoku-solver/

// Write a program to solve a Sudoku puzzle by filling the empty cells.

// A sudoku solution must satisfy **all of the following rules**:
// 1. Each of the digits 1-9 must occur exactly once in each row.
// 1. Each of the digits 1-9 must occur exactly once in each column.
// 1. Each of the digits 1-9 must occur exactly once in each of the 9 3x3
//    sub-boxes of the grid.

// The '.' character indicates empty cells.

// Example 1.
// Input: board = [
//     ["5","3",".",".","7",".",".",".","."],
//     ["6",".",".","1","9","5",".",".","."],
//     [".","9","8",".",".",".",".","6","."],
//     ["8",".",".",".","6",".",".",".","3"],
//     ["4",".",".","8",".","3",".",".","1"],
//     ["7",".",".",".","2",".",".",".","6"],
//     [".","6",".",".",".",".","2","8","."],
//     [".",".",".","4","1","9",".",".","5"],
//     [".",".",".",".","8",".",".","7","9"],
// ]
// Output: [
//     ["5","3","4","6","7","8","9","1","2"],
//     ["6","7","2","1","9","5","3","4","8"],
//     ["1","9","8","3","4","2","5","6","7"],
//     ["8","5","9","7","6","1","4","2","3"],
//     ["4","2","6","8","5","3","7","9","1"],
//     ["7","1","3","9","2","4","8","5","6"],
//     ["9","6","1","5","3","7","2","8","4"],
//     ["2","8","7","4","1","9","6","3","5"],
//     ["3","4","5","2","8","6","1","7","9"],
// ]
// Explanation: The input board is shown above and the only valid solution is
// shown below:
// https://upload.wikimedia.org/wikipedia/commons/thumb/3/31/Sudoku-by-L2G-20050714_solution.svg/250px-Sudoku-by-L2G-20050714_solution.svg.png

// Constraints:
// * board.length == 9
// * board[i].length == 9
// * board[i][j] is a digit or '.'.
// * It is **guaranteed** that the input board has only one solution.

// func solveSudoku(board [][]byte)  {
//
// }

const N = 9
const EmptyByte = byte('.')

type sudokuSolverSolver struct {
	board [][]byte
	index *sudokuSolverIndex
}

func newSudokuSolverSolver(board [][]byte) *sudokuSolverSolver {
	return &sudokuSolverSolver{board: board, index: newSudokuSolverIndex(board)}
}

// Solve solves the Sudoku Solver problem.
func (s *sudokuSolverSolver) Solve() {
	_ = s.solve(0, 0)
}

// Solve the board recursively using backtracking.
func (s *sudokuSolverSolver) solve(row, col int) bool {
	// Negative values for the row and column index indicate
	// we've made it to the end of the board successfully.
	if row < 0 && col < 0 {
		return true
	}

	// If this cell is already not empty, we know it was set
	// in the initial condition, so we just forward it.
	if !s.isEmpty(row, col) {
		nextRow, nextCol := nextCell(row, col)
		return s.solve(nextRow, nextCol)
	}

	// Try different candidate values at this position. If the
	// candidate is invalid, move on. Otherwise forward the call
	// along recursively on the next cell.
	for cand := 1; cand <= N; cand++ {
		if ok := s.index.maybeAdd(cand, row, col); !ok {
			continue
		}
		s.board[row][col] = byte('0' + cand)
		nextRow, nextCol := nextCell(row, col)
		if solved := s.solve(nextRow, nextCol); solved {
			return true
		} else {
			s.index.remove(cand, row, col)
			s.board[row][col] = EmptyByte
		}
	}
	return false
}

func nextCell(row, col int) (int, int) {
	if col+1 < N {
		return row, col + 1
	}
	if row+1 < N {
		return row + 1, 0
	}
	return -1, -1
}

func (s *sudokuSolverSolver) isEmpty(row, col int) bool {
	return s.board[row][col] == EmptyByte
}

type sudokuSolverIndex struct {
	rsets, csets, sbsets [N]*sudokuSolverBitset
}

func newSudokuSolverIndex(board [][]byte) *sudokuSolverIndex {
	index := &sudokuSolverIndex{}
	for r, c := 0, 0; r >= 0 && c >= 0; r, c = nextCell(r, c) {
		v := board[r][c]
		if v == EmptyByte {
			continue
		}
		// Problem constraints guarantee the board has a
		// solution so we can ignore the "ok" variable
		// (i.e. board is _not_ invalid to begin with).
		_ = index.maybeAdd(int(v-'0'), r, c)
	}
	return index
}

func (si *sudokuSolverIndex) maybeAdd(val, row, col int) bool {
	// Returns an "ok" variable to indicate if the addition is valid
	if si.rowSet(row).has(val-1) || si.colSet(col).has(val-1) || si.subBoardSet(row, col).has(val-1) {
		return false
	}
	si.rowSet(row).add(val - 1)
	si.colSet(col).add(val - 1)
	si.subBoardSet(row, col).add(val - 1)
	return true
}

func (si *sudokuSolverIndex) remove(val, row, col int) {
	si.rowSet(row).remove(val - 1)
	si.colSet(col).remove(val - 1)
	si.subBoardSet(row, col).remove(val - 1)
}

func (si *sudokuSolverIndex) rowSet(row int) *sudokuSolverBitset {
	if si.rsets[row] == nil {
		si.rsets[row] = newSudokuSolverBitset()
	}
	return si.rsets[row]
}

func (si *sudokuSolverIndex) colSet(col int) *sudokuSolverBitset {
	if si.csets[col] == nil {
		si.csets[col] = newSudokuSolverBitset()
	}
	return si.csets[col]
}

func (si *sudokuSolverIndex) subBoardSet(row, col int) *sudokuSolverBitset {
	idx := si.subBoard(row, col)
	if si.sbsets[idx] == nil {
		si.sbsets[idx] = newSudokuSolverBitset()
	}
	return si.sbsets[idx]
}

func (si *sudokuSolverIndex) subBoard(row, col int) int {
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if row < 3*(i+1) && col < 3*(j+1) {
				return 3*i + j
			}
		}
	}
	return -1
}

type sudokuSolverBitset struct {
	bitset [N]bool
}

func newSudokuSolverBitset() *sudokuSolverBitset {
	return &sudokuSolverBitset{bitset: [N]bool{}}
}

func (b *sudokuSolverBitset) add(i int) {
	b.bitset[i] = true
}

func (b *sudokuSolverBitset) remove(i int) {
	b.bitset[i] = false
}

func (b *sudokuSolverBitset) has(i int) bool {
	return b.bitset[i]
}
