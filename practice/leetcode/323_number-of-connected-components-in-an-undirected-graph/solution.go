package numberofconnectedcomponentsinanundirectedgraph

// LeetCode #323.
// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/

// You have a graph of n nodes. You are given an integer n and an array edges
// where edges[i] = [a_i, b_i] indicates that there is an edge between a_i and
// b_i in the graph.

// Return the number of connected components in the graph.

// Example 1:
// Input: n = 5, edges = [[0,1],[1,2],[3,4]]
// Output: 2

// Example 2:
// Input: n = 5, edges = [[0,1],[1,2],[2,3],[3,4]]
// Output: 1

// Constraints:
// - 1 <= n <= 2000
// - 1 <= edges.length <= 5000
// - edges[i].length == 2
// - 0 <= a_i <= b_i < n
// - a_i != b_i
// - There are no repeated edges.

type Solver struct {
	n     int
	edges [][]int
	reps  []int
	sizes []int
	count int
}

func NewSolver(n int, edges [][]int) *Solver {
	return &Solver{n: n, edges: edges}
}

func (s *Solver) Solve() int {
	s.init()
	for _, edge := range s.edges {
		s.union(edge[0], edge[1])
	}
	return s.count
}

func (s *Solver) init() {
	s.count = s.n
	s.reps, s.sizes = make([]int, s.n), make([]int, s.n)
	for i := 0; i < s.n; i++ {
		s.reps[i] = i
		s.sizes[i] = 1
	}
}

func (s *Solver) find(node int) int {
	if s.reps[node] == node {
		return node
	}
	// Recursively find the canonical representative for this node.
	s.reps[node] = s.find(s.reps[node])
	return s.reps[node]
}

func (s *Solver) union(nodeX, nodeY int) {
	nodeX, nodeY = s.find(nodeX), s.find(nodeY)
	if nodeX == nodeY {
		// Nothing to do, nodes are in the same connected component.
		return
	}
	if s.sizes[nodeX] > s.sizes[nodeY] {
		s.reps[nodeY] = nodeX
		s.sizes[nodeX] += s.sizes[nodeY]
	} else {
		s.reps[nodeX] = nodeY
		s.sizes[nodeY] += s.sizes[nodeX]
	}
	// Merged two components, decrement count of connected components.
	s.count--
}
