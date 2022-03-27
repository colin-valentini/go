package leetcode

// LeetCode #1020.
// Link: https://leetcode.com/problems/number-of-enclaves/

// You are given an m x n binary matrix grid, where 0 represents a sea cell and
// 1 represents a land cell.

// A move consists of walking from one land cell to another adjacent 
// (4-directionally) land cell or walking off the boundary of the grid.

// Return the number of land cells in grid for which we cannot walk off the 
// boundary of the grid in any number of moves.

// Example 1.
// Input: grid = [
// 	[0,0,0,0],
// 	[1,0,1,0],
// 	[0,1,1,0],
// 	[0,0,0,0],
// ]
// Output: 3
// Explanation: There are three 1s that are enclosed by 0s, and one 1 that is 
// not enclosed because its on the boundary.

// Example 2.
// Input: grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
// Output: 0
// Explanation: All 1s are either on the boundary or can reach the boundary.

// Constraints:
// - m == grid.length
// - n == grid[i].length
// - 1 <= m, n <= 500
// - grid[i][j] is either 0 or 1.
func NumberOfEnclaves(grid [][]int) int {
	return numEnclaves(grid)
}

// Procedure:
// (1) DFS from each border cell that is a 1. Any cells connected
//     4-directionally to a border cell is reset to the value 2 in 
//     the grid.
// (2) Count every cell in the grid that has value 1 after we did (1).
func numEnclaves(grid [][]int) int {
	for r := range grid {
		for c := range grid[r] {
			if r == 0 || r == len(grid)-1 || c == 0 || c == len(grid[r])-1 {
				dfs(grid, r, c)
			}
		}
	}
	count := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]int, r, c int) {
	if grid[r][c] != 1 {
		return
	}
	grid[r][c] = -1
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