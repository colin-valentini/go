package leetcode

// LeetCode #2133.
// Link: https://leetcode.com/problems/check-if-every-row-and-column-contains-all-numbers/

// An n x n matrix is valid if every row and every column contains all the integers
// from 1 to n (inclusive).

// Given an n x n integer matrix matrix, return true if the matrix is valid.
// Otherwise, return false.

// Example 1:
// Input: matrix = [
// 	[1,2,3],
// 	[3,1,2],
// 	[2,3,1]
// ]
// Output: true
// Explanation: In this case, n = 3, and every row and column contains the numbers
// 1, 2, and 3. Hence, we return true.

// Example 2:
// Input: matrix = [
// 	[1,1,1],
// 	[1,2,3],
// 	[1,2,3]
// ]
// Output: false
// Explanation: In this case, n = 3, but the first row and the first column do not
// contain the numbers 2 or 3. Hence, we return false.

// Constraints:
// * n == matrix.length == matrix[i].length
// * 1 <= n <= 100
// * 1 <= matrix[i][j] <= n

// func checkValid(matrix [][]int) bool {
//
// }

type checkIfEveryRowAndColumnContainsAllNumbersSolver struct {
	matrix [][]int
}

func newCheckIfEveryRowAndColumnContainsAllNumbersSolver(matrix [][]int) *checkIfEveryRowAndColumnContainsAllNumbersSolver {
	return &checkIfEveryRowAndColumnContainsAllNumbersSolver{matrix}
}

func (s *checkIfEveryRowAndColumnContainsAllNumbersSolver) Solve() bool {
	// Input matrix is known to be square by definition.
	N := len(s.matrix)
	for r := 0; r < N; r++ {
		rowCount, colCount := make([]int, N), make([]int, N)
		for c := 0; c < N; c++ {
			rowVal := s.matrix[r][c] - 1
			colVal := s.matrix[c][r] - 1
			if rowCount[rowVal] > 0 || colCount[colVal] > 0 {
				return false
			}
			rowCount[rowVal]++
			colCount[colVal]++
		}
	}
	return true
}
