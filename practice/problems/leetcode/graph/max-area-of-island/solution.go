package leetcode

// LeetCode #695.
// Link: https://leetcode.com/problems/max-area-of-island/

// You are given an m x n binary matrix grid. An island is a group of 1's
// (representing land) connected 4-directionally (horizontal or vertical.)
// You may assume all four edges of the grid are surrounded by water.

// The area of an island is the number of cells with a value 1 in the island.

// Return the maximum area of an island in grid. If there is no island, return 0.

// Example 1.
// Input: grid = [
// 	[0,0,1,0,0,0,0,1,0,0,0,0,0],
// 	[0,0,0,0,0,0,0,1,1,1,0,0,0],
// 	[0,1,1,0,1,0,0,0,0,0,0,0,0],
// 	[0,1,0,0,1,1,0,0,1,0,1,0,0],
// 	[0,1,0,0,1,1,0,0,1,1,1,0,0],
// 	[0,0,0,0,0,0,0,0,0,0,1,0,0],
// 	[0,0,0,0,0,0,0,1,1,1,0,0,0],
// 	[0,0,0,0,0,0,0,1,1,0,0,0,0],
// ]
// Output: 6
// Explanation: The answer is not 11, because the island must be connected
// 4-directionally.

// Example 2.
// Input: grid = [
//   [0,0,0,0,0,0,0,0],
// ]
// Output: 0

// Constraints:
// - m == grid.length
// - n == grid[i].length
// - 1 <= m, n <= 50
// - grid[i][j] is either 0 or 1.

// MaxAreaOfIsland is a solution to the max area of island problem.
func MaxAreaOfIsland(grid [][]int) int {
	return maxAreaOfIsland(grid)
}

// Procedure:
// - DFS from each cell. Return 0 if the cell is not a 1. Otherwise
//   return 1 + the count returned by the recursive DFS of neighbor cells.
// - We avoid revisiting cells by maintaining a set which tracks which cells
//   have been visited.
func maxAreaOfIsland(grid [][]int) int {
	max := 0
	visited := map[cell]nothing{}
	for r := range grid {
		for c := range grid[r] {
			if m := dfs(grid, r, c, visited); m > max {
				max = m
			}
		}
	}
	return max
}

func dfs(grid [][]int, r, c int, visited map[cell]nothing) int {
	if grid[r][c] != 1 {
		return 0
	}
	if _, ok := visited[cell{r, c}]; ok {
		return 0
	}
	visited[cell{r, c}] = nothing{}
	count := 1
	if 0 <= r-1 {
		count += dfs(grid, r-1, c, visited)
	}
	if r+1 < len(grid) {
		count += dfs(grid, r+1, c, visited)
	}
	if 0 <= c-1 {
		count += dfs(grid, r, c-1, visited)
	}
	if c+1 < len(grid[r]) {
		count += dfs(grid, r, c+1, visited)
	}
	return count
}

type cell struct{ r, c int }
type nothing struct{}
