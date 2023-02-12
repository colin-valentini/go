package pacificatlanticwaterflow

// LeetCode #417.
// https://leetcode.com/problems/pacific-atlantic-water-flow/

// func pacificAtlantic(heights [][]int) [][]int {
//
// }

type Solver struct {
	n, m    int
	heights [][]int
	visited *cache
}

func NewSolver(heights [][]int) *Solver {
	n, m := len(heights), len(heights[0])
	return &Solver{n: n, m: m, heights: heights}
}

func (s *Solver) Solve() [][]int {
	s.visited = newCache()
	results := [][]int{}
	for r := 0; r < s.n; r++ {
		for c := 0; c < s.m; c++ {
			coord := coordinates{row: r, col: c}
			path := make(map[coordinates]nothing)
			if reach := s.flowSearch(coord, path); reach.isBoth() {
				results = append(results, []int{coord.row, coord.col})
			}
		}
	}
	return results
}

func (s *Solver) isOnGrid(coord coordinates) bool {
	return 0 <= coord.row && coord.row < s.n && 0 <= coord.col && coord.col < s.m
}

type reachability struct {
	pacific, atlantic bool
}

func (r reachability) isBoth() bool {
	return r.pacific && r.atlantic
}

func (r reachability) merge(other reachability) reachability {
	return reachability{
		pacific:  r.pacific || other.pacific,
		atlantic: r.atlantic || other.atlantic,
	}
}

type coordinates struct {
	row, col int
}

func (s *Solver) isPacific(coord coordinates) bool {
	return coord.row < 0 || (coord.row < s.n && coord.col < 0)
}

func (s *Solver) isAtlantic(coord coordinates) bool {
	return coord.row >= s.n || (0 <= coord.row && coord.col >= s.m)
}

func (s *Solver) flowSearch(coord coordinates, path map[coordinates]nothing) reachability {
	if s.isPacific(coord) {
		return reachability{pacific: true}
	}
	if s.isAtlantic(coord) {
		return reachability{atlantic: true}
	}
	if cached, ok := s.visited.get(coord); ok {
		return cached
	}
	reach := reachability{}
	path[coord] = nothing{}
	for _, neighbor := range s.flowNeighbors(coord) {
		reach = reach.merge(s.flowSearch(neighbor, path))
	}
	s.visited.set(coord, reach)
	delete(path, coord)
	return reach
}

func (s *Solver) flowNeighbors(coord coordinates) []coordinates {
	coordHeight := s.heights[coord.row][coord.col]
	flowNeighbors := make([]coordinates, 0, 4)
	neighbors := []coordinates{
		{row: coord.row - 1, col: coord.col},
		{row: coord.row + 1, col: coord.col},
		{row: coord.row, col: coord.col - 1},
		{row: coord.row, col: coord.col + 1},
	}
	for _, neighbor := range neighbors {
		// Allow off grid, but filter out greater height neighbors
		if !s.isOnGrid(neighbor) {
			flowNeighbors = append(flowNeighbors, neighbor)
		} else if s.heights[neighbor.row][neighbor.col] <= coordHeight {
			flowNeighbors = append(flowNeighbors, neighbor)
		}
	}
	return flowNeighbors
}

type cache struct {
	reach map[coordinates]reachability
}

func newCache() *cache {
	return &cache{reach: make(map[coordinates]reachability)}
}

func (c *cache) get(coord coordinates) (reachability, bool) {
	r, ok := c.reach[coord]
	return r, ok
}

func (c *cache) set(coord coordinates, reach reachability) {
	c.reach[coord] = reach
}

type nothing struct{}
