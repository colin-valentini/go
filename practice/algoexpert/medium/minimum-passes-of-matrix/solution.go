package algoexpert

// https://www.algoexpert.io/questions/minimum-passes-of-matrix

// Write a function that takes in an integer matrix of potentially unequal
// height and width and returns the minimum number of passes required to convert
// all negative integers in the matrix to positive integers.

// A negative integer in the matrix can only be converted to a positive integer
// if one or more of its adjacent elements is positive. An adjacent element is
// an element that is to the left, to the right, above, or below the current
// element in the matrix. Converting a negative to a positive simply involves
// multiplying it by -1.

// Note that the 0 value is neither positive nor negative, meaning
// that a 0 can't convert an adjacent negative to a positive.

// A single pass through the matrix involves converting all the negative
// integers that can be converted at a particular point in time. For example,
// consider the following input matrix:

// [
//   [ 0, -2, -1],
//   [-5,  2,  0],
//   [-6, -2,  0],
// ]

// After a first pass, only 3 values can be converted to positives:

// [
//   [0,  2, -1],
//   [5,  2,  0],
//   [-6, 2,  0],
// ]

// After a second pass, the remaining negative values can all be converted to
// positives:

// [
//   [0, 2, 1],
//   [5, 2, 0],
//   [6, 2, 0],
// ]

// Note that the input matrix will always contain at least one element. If the
// negative integers in the input matrix can't all be converted to positives,
// regardless of how many passes are run, your function should return
// -1.

// Sample Input:
// matrix = [
//   [0, -1, -3,  2,  0],
//   [1, -2, -5, -1, -3],
//   [3,  0,  0, -4, -1],
// ]

// Sample Output:
// 3

// MinimumPassesOfMatrix is a solution to the "Minimum Passes of Matrix" problem
func MinimumPassesOfMatrix(matrix [][]int) int {
	return newMinimumPassesOfMatrixSolver(matrix).Solve()
}

type minimumPassesOfMatrixSolver struct {
	n, m                int
	matrix              [][]int
	negCount, flipCount int
	queue, nextQueue    []cell
	passes              int
}

type cell struct {
	row, col int
}

func newMinimumPassesOfMatrixSolver(matrix [][]int) *minimumPassesOfMatrixSolver {
	return &minimumPassesOfMatrixSolver{matrix: matrix}
}

// 1. Do one pass upfront to:
//   - Find all cells with a positive value, and
//   - Count the number of negative values
//
// 2. If there are no negative values, return 0 (nothing to do)
// 3. If there are negatives, but no positives, return -1 (no way to solve)
// 4. Run a breadth-first-search starting from all of the positive cells:
//   - Pop a positive cell off the queue
//   - Find any negative neighbors, put them on the "next" queue
//   - Flip all the negative neighbors we found (count the flips)
//   - Swap the next and current queues
//   - Increment our count of passes
//   - If the current queue is empty, swap the next and current
//   - Repeat all of the above until next is empty
//     5. Return the number of passes less 1 (because the last sequence of pops
//     won't flip any values)
func (s *minimumPassesOfMatrixSolver) Solve() int {
	s.init()

	if s.negCount == 0 {
		return 0 // Nothing to do
	}
	if len(s.queue) == 0 {
		return -1 // No solution
	}
	s.bfs()
	if s.flipCount < s.negCount {
		return -1 // Unable to flip all positives
	}
	return s.passes - 1
}

func (s *minimumPassesOfMatrixSolver) init() {
	// Save the shape of the matrix (assume we'll
	// always have non-zero number of columns)
	s.n, s.m = len(s.matrix), len(s.matrix[0])

	// Do a single pass on the matrix to push the
	// positive cells into our queue, and count the
	// number of negative values in the matrix.
	s.queue = []cell{}
	for r := 0; r < s.n; r++ {
		for c := 0; c < s.m; c++ {
			if s.matrix[r][c] > 0 {
				s.queue = append(s.queue, cell{row: r, col: c})
			} else if s.matrix[r][c] < 0 {
				s.negCount++
			}
		}
	}
}

func (s *minimumPassesOfMatrixSolver) bfs() {
	s.nextQueue = []cell{}
	for len(s.queue) > 0 {
		curr := s.queue[0]
		s.queue = s.queue[1:]
		neighbors := s.negativeNeighbors(curr)
		s.flip(neighbors)
		s.nextQueue = append(s.nextQueue, neighbors...)
		if len(s.queue) == 0 {
			s.queue, s.nextQueue = s.nextQueue, s.queue
			s.passes++
		}
	}
}

func (s *minimumPassesOfMatrixSolver) negativeNeighbors(at cell) []cell {
	neighbors := []cell{}
	if 0 <= at.row-1 && s.matrix[at.row-1][at.col] < 0 {
		neighbors = append(neighbors, cell{row: at.row - 1, col: at.col})
	}
	if at.row+1 < s.n && s.matrix[at.row+1][at.col] < 0 {
		neighbors = append(neighbors, cell{row: at.row + 1, col: at.col})
	}
	if 0 <= at.col-1 && s.matrix[at.row][at.col-1] < 0 {
		neighbors = append(neighbors, cell{row: at.row, col: at.col - 1})
	}
	if at.col+1 < s.m && s.matrix[at.row][at.col+1] < 0 {
		neighbors = append(neighbors, cell{row: at.row, col: at.col + 1})
	}
	return neighbors
}

func (s *minimumPassesOfMatrixSolver) flip(cells []cell) {
	for _, at := range cells {
		s.matrix[at.row][at.col] *= -1
	}
	s.flipCount += len(cells)
}
