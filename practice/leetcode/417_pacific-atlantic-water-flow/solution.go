package pacificatlanticwaterflow

// LeetCode #417.
// https://leetcode.com/problems/pacific-atlantic-water-flow/

// func pacificAtlantic(heights [][]int) [][]int {
//
// }

// Solver is a type that can solve the Pacific Atlantic Water Flow problem.
type Solver struct {
	n, m    int
	heights [][]int
}

// NewSolver returns a pointer to a new instance of the Solver type.
func NewSolver(heights [][]int) *Solver {
	return &Solver{n: len(heights), m: len(heights[0]), heights: heights}
}

// Solve solves the Pacific Atlantic Water Flow problem.
//
// The solution use breadth-first search from each ocean
// separately. We start with the boundaries, and work our
// way inwards. We can move to a neighboring cell only if
// that neighbor has a greater or equal height (i.e. that
// cell could flow "into" the current cell).
// Once we have both sets of cells that we found using BFS
// from each ocean, we return the set intersection.
//
// Time: O(N * M) since we visit each cell at most twice.
// Spcae: O(N * M) since we store each cell at most twice.
func (s *Solver) Solve() [][]int {
	cells := s.atlantic().Intersect(s.pacific())
	solution := make([][]int, 0, len(cells))
	for cell := range cells {
		solution = append(solution, []int{cell.r, cell.c})
	}
	return solution
}

func (s *Solver) atlantic() Set {
	q := Queue{}
	// Last row.
	for c := 0; c < s.m; c++ {
		q.Push(Cell{r: s.n - 1, c: c})
	}
	// Last column, except last row.
	for r := 0; r < s.n-1; r++ {
		q.Push(Cell{r: r, c: s.m - 1})
	}
	return s.bfs(q)
}

func (s *Solver) pacific() Set {
	q := Queue{}
	// First row.
	for c := 0; c < s.m; c++ {
		q.Push(Cell{r: 0, c: c})
	}
	// First column except first row.
	for r := 1; r < s.n; r++ {
		q.Push(Cell{r: r, c: 0})
	}
	return s.bfs(q)
}

func (s *Solver) bfs(q Queue) Set {
	visited := make(Set)
	for len(q) > 0 {
		cell := q.Pop()
		if visited.Has(cell) {
			continue
		}
		visited.Add(cell)
		for _, neighbor := range cell.Neighbors() {
			if s.inBounds(neighbor) && s.height(neighbor) >= s.height(cell) {
				q.Push(neighbor)
			}
		}
	}
	return visited
}

func (s *Solver) height(cell Cell) int {
	return s.heights[cell.r][cell.c]
}

func (s *Solver) inBounds(cell Cell) bool {
	return 0 <= cell.r && cell.r < s.n && 0 <= cell.c && cell.c < s.m
}

type Cell struct {
	r, c int
}

func (c Cell) Neighbors() []Cell {
	return []Cell{
		{r: c.r - 1, c: c.c},
		{r: c.r + 1, c: c.c},
		{r: c.r, c: c.c - 1},
		{r: c.r, c: c.c + 1},
	}
}

type Queue []Cell

func (q *Queue) Push(cell Cell) {
	*q = append(*q, cell)
}

func (q *Queue) Pop() Cell {
	cell := (*q)[0]
	*q = (*q)[1:]
	return cell
}

type Set map[Cell]Nothing

func (s Set) Add(v Cell) {
	s[v] = Nothing{}
}

func (s Set) Has(v Cell) bool {
	_, ok := s[v]
	return ok
}

func (s Set) Intersect(other Set) Set {
	intersection := make(Set)
	for cell := range s {
		if other.Has(cell) {
			intersection.Add(cell)
		}
	}
	return intersection
}

type Nothing struct{}
