package rottingoranges

// LeetCode #994.
// https://leetcode.com/problems/rotting-oranges/

type Solver struct {
	m, n int
	grid [][]int
}

func NewSolver(grid [][]int) *Solver {
	return &Solver{m: len(grid), n: len(grid[0]), grid: grid}
}

func (s *Solver) Solve() int {
	q := newQueue()
	freshCount := 0
	for r := 0; r < s.m; r++ {
		for c := 0; c < s.n; c++ {
			cell := loc{r: r, c: c}
			if v := s.value(cell); v == rotten {
				q.enqueue(cell)
			} else if v == fresh {
				freshCount++
			}
		}
	}
	minutes := 0
	next := newQueue()
	for q.len() > 0 {
		cell := q.dequeue()
		for _, neighbor := range cell.neighbors() {
			if !s.isInBounds(neighbor) || s.value(neighbor) != fresh {
				continue
			}
			next.enqueue(neighbor)
			s.set(neighbor, rotten)
			freshCount--
		}
		if q.len() == 0 && next.len() > 0 {
			q, next = next, q
			minutes++
		}
	}
	if freshCount > 0 {
		return -1
	}
	return minutes
}

func (s *Solver) isInBounds(cell loc) bool {
	return 0 <= cell.r && cell.r < s.m && 0 <= cell.c && cell.c < s.n
}

func (s *Solver) value(cell loc) int {
	return s.grid[cell.r][cell.c]
}

func (s *Solver) set(cell loc, value int) {
	s.grid[cell.r][cell.c] = value
}

const (
	empty = iota
	fresh
	rotten
)

type loc struct {
	r, c int
}

func (cell loc) neighbors() []loc {
	return []loc{
		{cell.r - 1, cell.c},
		{cell.r + 1, cell.c},
		{cell.r, cell.c - 1},
		{cell.r, cell.c + 1},
	}
}

type queue struct {
	elems []loc
}

func newQueue() *queue {
	return &queue{elems: []loc{}}
}

func (q *queue) enqueue(cell loc) {
	q.elems = append(q.elems, cell)
}

func (q *queue) dequeue() loc {
	// Panics if queue is empty.
	cell := q.elems[0]
	q.elems = q.elems[1:]
	return cell
}

func (q *queue) len() int {
	return len(q.elems)
}
