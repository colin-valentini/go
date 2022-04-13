package leetcode

// LeetCode #1254.
// Link: https://leetcode.com/problems/number-of-closed-islands/

// Given a 2D grid consists of 0s (land) and 1s (water).  An island is a maximal
// 4-directionally connected group of 0s and a closed island is an island
// totally (all left, top, right, bottom) surrounded by 1s.

// Return the number of closed islands.

// Example 1.
// Input: grid = [
// 	 [1,1,1,1,1,1,1,0],
// 	 [1,0,0,0,0,1,1,0],
// 	 [1,0,1,0,1,1,1,0],
// 	 [1,0,0,0,0,1,0,1],
// 	 [1,1,1,1,1,1,1,0],
// ]
// Output: 2
// Explanation:
// Islands in gray are closed because they are completely surrounded by water
// (group of 1s).

// Example 2.
// Input: grid = [
// 	 [0,0,1,0,0],
// 	 [0,1,0,1,0],
// 	 [0,1,1,1,0],
// ]
// Output: 1

// Example 3.
// Input: grid = [
// 	 [1,1,1,1,1,1,1],
//   [1,0,0,0,0,0,1],
//   [1,0,1,1,1,0,1],
//   [1,0,1,0,1,0,1],
//   [1,0,1,1,1,0,1],
//   [1,0,0,0,0,0,1],
//   [1,1,1,1,1,1,1],
// ]
// Output: 2

// Constraints:
// - 1 <= grid.length, grid[0].length <= 100
// - 0 <= grid[i][j] <=1

// NumberOfClosedIslands is a solution to the number of closed islands problem.
func NumberOfClosedIslands(grid [][]int) int {
	return closedIsland(grid)
}

// Prodedure:
// (1) Initiate a DFS from inner cells with a 0 (land), flipping them to 2 as
//     we go to denote that we've visited them.
// (2) Each DFS call will return true if none of the neighbors (or itself)
//     touch the border (which means that the this "island", is not an island).
// (3) If the top-level call returns true, increment our counter of islands.
func closedIsland(grid [][]int) int {
	count := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 0 && dfs(grid, r, c) {
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]int, r, c int) bool {
	// Either it's water or in-progress.
	if grid[r][c] > 0 {
		return true
	}
	// It's definitely not an island (we know it for sure).
	if grid[r][c] < 0 {
		return false
	}
	// Once we get to the border, we know it's not an island.
	if r == 0 || r == len(grid)-1 || c == 0 || c == len(grid[r])-1 {
		return false
	}
	// Tag this cell as an island (for now), and check neighbors. If
	// any neighbor is not an island, we know this isn't an island
	// and we'll update the cell value accordingly.
	grid[r][c] = 1
	if 0 <= r-1 && !dfs(grid, r-1, c) {
		grid[r][c] = -1
		return false
	}
	if r+1 < len(grid) && !dfs(grid, r+1, c) {
		grid[r][c] = -1
		return false
	}
	if 0 <= c-1 && !dfs(grid, r, c-1) {
		grid[r][c] = -1
		return false
	}
	if c+1 < len(grid[r]) && !dfs(grid, r, c+1) {
		grid[r][c] = -1
		return false
	}
	return true
}
