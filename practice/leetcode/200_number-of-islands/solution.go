package leetcode

// LeetCode #200.
// Link: https://leetcode.com/problems/number-of-islands/

// Given an m x n 2D binary grid grid which represents a map of '1's (land) and
// '0's (water), return the number of islands.

// An island is surrounded by water and is formed by connecting adjacent lands
// horizontally or vertically. You may assume all four edges of the grid are all
// surrounded by water.

// Exmample 1.
// Input: grid = [
//   ["1","1","1","1","0"],
//   ["1","1","0","1","0"],
//   ["1","1","0","0","0"],
//   ["0","0","0","0","0"]
// ]
// Output: 1

// Example 2.
// Input: grid = [
//   ["1","1","0","0","0"],
//   ["1","1","0","0","0"],
//   ["0","0","1","0","0"],
//   ["0","0","0","1","1"]
// ]
// Output: 3

// Constraints:
// - m == grid.length
// - n == grid[i].length
// - 1 <= m, n <= 300
// - grid[i][j] is '0' or '1'.

// NumIslands is a solution to the number of islands problem.
func NumIslands(grid [][]byte) int {
	return numIslands(grid)
}

// Procedure:
// (1) Dive into each cell with a "1" value and do a DFS
// (2) Tag the cell as "-1" to denote visited.
// (3) For each of the immediate neighbords of this cell,
//
//	dive into them if they a "1", repeating step (1)
//
// We're able to overwrite the grid values and re-use the existing
// memory because the problem statement says nothing about not
// mutating this space.
func numIslands(grid [][]byte) int {
	count := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c]-'0' == 1 {
				dfs(grid, r, c)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, r, c int) {
	// Terminate if this cell isn't a 1. This is the case when
	// we've either visited it already, or it's a "0".
	if grid[r][c]-'0' != 1 {
		return
	}
	// Tag this cell as visited now to avoid recursive calls
	// repeatedly visiting this cell.
	grid[r][c] = '2'
	// Dive into neighbors.
	if 0 <= r-1 {
		dfs(grid, r-1, c)
	}
	if r+1 < len(grid) {
		dfs(grid, r+1, c)
	}
	if 0 <= c-1 {
		dfs(grid, r, c-1)
	}
	if c+1 < len(grid[r]) {
		dfs(grid, r, c+1)
	}
}
