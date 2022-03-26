package leetcode

// LeetCode #130.
// Link: https://leetcode.com/problems/surrounded-regions/

// Given an m x n matrix board containing 'X' and 'O', capture all regions that
// are 4-directionally surrounded by 'X'.

// A region is captured by flipping all 'O's into 'X's in that surrounded region.

// Example 1.
// Input: board = [
//   ["X","X","X","X"],
//   ["X","O","O","X"],
//   ["X","X","O","X"],
//   ["X","O","X","X"],
// ]
// Output: [
//   ["X","X","X","X"],
//   ["X","X","X","X"],
//   ["X","X","X","X"],
//   ["X","O","X","X"],
// ]
// Explanation: Surrounded regions should not be on the border, which means that
// any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is 
// not on the border and it is not connected to an 'O' on the border will be 
// flipped to 'X'. Two cells are connected if they are adjacent cells connected
// horizontally or vertically.

// Example 2.
// Input: board = [["X"]]
// Output: [["X"]]

// Constraints:
// - m == board.length
// - n == board[i].length
// - 1 <= m, n <= 200
// - board[i][j] is 'X' or 'O'.

// SurroundedRegions is a solution to the surrounded regions problem.
func SurroundedRegions(board [][]byte) {
	solve(board)
}

// Procedure:
// (1) For each border cell, if it's an 'O' value, do a DFS starting from
//     this cell. Each cell that is reachable by DFS (4-directionally) and 
//     is an 'O' value we'll flip to 'R' for "reachable" (these are cells that
//     are un-surroundable given their continuity to the border).
// (3) Iterate over all board cells. If the cell is an 'R' flip it back to an
//     an 'O'. If it's an 'O', flip it to an 'X' (since we obviously didn't
//     reach it from the border).
//
// NOTE: There may be some optimization to determine we do a DFS from the border
// inwards (and record "unsurroundable" cells), or whether we do a DFS out
// from the center (and record "surroundable" cells that don't touch the border).
// How to assess whether or not an inwards or outwards search is faster without
// extra work is unclear, and may actually require more work overall than just
// choosing one method or the other arbitrarily.
func solve(board [][]byte)  {
	for r := range board {
		for c := range board[r] {
			if r == 0 || r == len(board)-1 || c == 0 || c == len(board[r])-1 {
				dfs(board, r, c)
			}
		}
	}
    for r := range board {
		for c := range board[r] {
			switch board[r][c] {
			case 'O':
				board[r][c] = 'X'
			case 'R':
				board[r][c] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, r, c int) {
	if board[r][c] == 'R' {
		return
	}
	if board[r][c] == 'X' {
		return
	}
	board[r][c] = 'R'
	if 0 <= r-1 {
		dfs(board, r-1, c)
	}
	if r+1 < len(board) {
		dfs(board, r+1, c)
	}
	if 0 <= c-1 {
		dfs(board, r, c-1)
	}
	if c+1 < len(board[r]) {
		dfs(board, r, c+1)
	}
}
