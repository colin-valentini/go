package leetcode

// LeetCode #733.
// Link: https://leetcode.com/problems/flood-fill/

// An image is represented by an m x n integer grid image where image[i][j]
// represents the pixel value of the image.

// You are also given three integers sr, sc, and newColor. You should perform
// a flood fill on the image starting from the pixel image[sr][sc].

// To perform a flood fill, consider the starting pixel, plus any pixels connected
// 4-directionally to the starting pixel of the same color as the starting pixel, 
// plus any pixels connected 4-directionally to those pixels (also with the same 
// color), and so on. Replace the color of all of the aforementioned pixels with 
// newColor.

// Return the modified image after performing the flood fill.

// Example 1:
//  --- --- ---       --- --- ---
// | 1 | 1 | 1 |     | 2 | 2 | 2 |
// | 1 | 1 | 0 | --> | 2 | 2 | 0 |
// | 1 | 0 | 1 |     | 2 | 0 | 1 |
//  --- --- ---       --- --- ---

// Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, newColor = 2
// Output: [[2,2,2],[2,2,0],[2,0,1]]
// Explanation: From the center of the image with position (sr, sc) = (1, 1) 
// (i.e., the red pixel), all pixels connected by a path of the same color as 
// the starting pixel (i.e., the blue pixels) are colored with the new color.
// Note the bottom corner is not colored 2, because it is not 4-directionally 
// connected to the starting pixel.

// Example 2:
// Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, newColor = 2
// Output: [[2,2,2],[2,2,2]]

// Constraints:
// - m == image.length
// - n == image[i].length
// - 1 <= m, n <= 50
// - 0 <= image[i][j], newColor < 216
// - 0 <= sr < m
// - 0 <= sc < n

// FloodFill is a solution to the flood fill problem.
func FloodFill(image [][]int, sr, sc, newColor int) [][]int {
	return floodFill(image, sr, sc, newColor)
}

func floodFill(image [][]int, sr, sc, newColor int) [][]int {
	m, n := len(image), len(image[0])
	color := image[sr][sc]

	visited := make(map[cell]nothing, m * n)

	queue := newQueue(m * n)
	queue.push(cell{sr, sc})

	for !queue.isEmpty() {
		cell := queue.pop()
		visited[cell] = nothing{}
		if image[cell.row][cell.col] != color {
			continue
		}
		image[cell.row][cell.col] = newColor
		for _, c := range neighbors(image, cell) {
			// It's critical that we check that if we've already visited 
			// this neighbor otherwise we could push too many cells into
			// the queue and cause a channel deadlock.
			if _, ok := visited[c]; !ok {
				queue.push(c)
			}
		}
	}
    return image
}

func neighbors(image [][]int, c cell) []cell {
	neighbors := make([]cell, 0, 4)
	candidates := []cell{
		{c.row-1, c.col},
		{c.row+1, c.col},
		{c.row, c.col-1},
		{c.row, c.col+1},
	}
	for _, c := range candidates {
		if isInBounds(image, c) {
			neighbors = append(neighbors, c)
		}
	}
	return neighbors
}

func isInBounds(image [][]int, c cell) bool {
	return 0 <= c.row && c.row < len(image) && 0 <= c.col && c.col < len(image[0])
}

type cell struct {
	row int
	col int
}

type queue struct {
	ch chan cell
}

func newQueue(size int) *queue {
	return &queue{ch: make(chan cell, size)}
}

func (q *queue) push(elem cell) {
	// Re-sizing is overly complicated for this use case.
	if len(q.ch) == cap(q.ch) {
		panic("cannot push, will deadlock")
	}
	q.ch <- elem
}

func (q *queue) pop() cell {
	return <- q.ch
}

func (q *queue) isEmpty() bool {
	return len(q.ch) == 0
}

type nothing struct{}