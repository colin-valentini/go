package wallsandgates

import "math"

// LeetCode #286.
// https://leetcode.com/problems/walls-and-gates/

// You are given an m x n grid rooms initialized with these three possible values.

// - -1 A wall or an obstacle.
// - 0 A gate.
// - INF Infinity means an empty room. We use the value 231 - 1 = 2147483647 to
//   represent INF as you may assume that the distance to a gate is less than
//   2147483647.

// Fill each empty room with the distance to its nearest gate. If it is
// impossible to reach a gate, it should be filled with INF.

// Example 1:
// Input: rooms = [
//   [2147483647, -1,         0,          2147483647],
//   [2147483647, 2147483647, 2147483647, -1        ],
//   [2147483647, -1,         2147483647, -1        ],
//   [0,          -1,         2147483647, 2147483647],
// ]
// Output: [
//   [3, -1, 0,  1],
//   [2,  2, 1, -1],
//   [1, -1, 2, -1],
//   [0, -1, 3,  4],
// ]

// Example 2:
// Input: rooms = [[-1]]
// Output: [[-1]]

// Constraints:
// - m == rooms.length
// - n == rooms[i].length
// - 1 <= m, n <= 250
// - rooms[i][j] is -1, 0, or 2^31 - 1.

// func wallsAndGates(rooms [][]int)  {
//
// }

const (
	WALL = -1
	GATE = 0
	INF  = math.MaxInt32
)

// Neighbor cell directions
var directions = []Cell{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type Solver struct {
	rooms [][]int
}

func NewSolver(rooms [][]int) *Solver {
	return &Solver{rooms: rooms}
}

func (s *Solver) Solve() {
	m := len(s.rooms)
	if m == 0 {
		return
	}
	n := len(s.rooms[0])

	var q Queue
	for row := range s.rooms {
		for col := range s.rooms[row] {
			if s.rooms[row][col] == GATE {
				q.Push(Cell{r: row, c: col})
			}
		}
	}

	for len(q) > 0 {
		cell := q.Pop()
		row, col := cell.r, cell.c
		distance := s.rooms[row][col]
		for _, delta := range directions {
			neighbor := Cell{r: row + delta.r, c: col + delta.c}
			if neighbor.r < 0 || neighbor.r >= m || neighbor.c < 0 || neighbor.c >= n {
				// Out-of-bounds. Skip.
				continue
			}
			if s.get(neighbor) != INF {
				// Room is either out of bounds, a wall or gate, or was already
				// processed by previous iterations (will only see INF on first)
				continue
			}
			s.set(neighbor, distance+1)
			q.Push(neighbor)
		}
	}
}

func (s *Solver) get(cell Cell) int {
	return s.rooms[cell.r][cell.c]
}

func (s *Solver) set(cell Cell, value int) {
	// Panics if room is out of bounds.
	s.rooms[cell.r][cell.c] = value
}

type Cell struct {
	r, c int
}

type Queue []Cell

func (q *Queue) Push(cell Cell) {
	*q = append(*q, cell)
}

func (q *Queue) Pop() Cell {
	// Panics if length not checked
	cell := (*q)[0]
	*q = (*q)[1:len(*q)]
	return cell
}
