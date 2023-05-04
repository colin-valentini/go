package graphvalidtree

// LeetCode #261.
// https://leetcode.com/problems/graph-valid-tree/

// You have a graph of n nodes labeled from 0 to n - 1. You are given an integer
// n and a list of edges where edges[i] = [ai, bi] indicates that there is an
// undirected edge between nodes ai and bi in the graph.

// Return true if the edges of the given graph make up a valid tree, and false
// otherwise.

// Example 1:
// Input: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
// Output: true

// Example 2:
// Input: n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
// Output: false

// Constraints:
// - 1 <= n <= 2000
// - 0 <= edges.length <= 5000
// - edges[i].length == 2
// - 0 <= a_i, b_i < n
// - a_i != b_i
// - There are no self-loops or repeated edges.

// func validTree(n int, edges [][]int) bool {
//
// }

type Solver struct {
	n     int
	edges [][]int
	par   []int
	rank  []int
}

func NewSolver(n int, edges [][]int) *Solver {
	return &Solver{n: n, edges: edges}
}

// Solve solves the Graph Valid Tree problem.
//
// The solution uses the UnionFind algorithm. We keep track of
// the number of connect sub-graphs. When we union two nodes together,
// if we notice that the nodes have the same common "parent", then they are
// already connected, and the edge we're looking at would form a cycle.
// Because a valid tree is acyclic, we know this can't be a tree.
//
// When we're done, the other thing we need to check is that the nodes form
// a single graph (i.e. no disconnected nodes). Otherwise, not a tree.
//
// Time: O(N+N*alpha(N)). Space: O(N).
func (s *Solver) Solve() bool {
	s.init()
	graphs := s.n
	for _, edge := range s.edges {
		x, y := edge[0], edge[1]
		delta := s.union(x, y)
		if delta == 0 {
			return false
		}
		graphs += delta
	}
	return graphs == 1
}

func (s *Solver) init() {
	s.par = make([]int, s.n)
	s.rank = make([]int, s.n)
	for i := 0; i < s.n; i++ {
		s.par[i] = i
		s.rank[i] = 1
	}
}

func (s *Solver) find(node int) int {
	if s.par[node] != node {
		s.par[node] = s.find(s.par[node])
	}
	return s.par[node]
}

func (s *Solver) union(x, y int) int {
	x, y = s.find(x), s.find(y)
	if x == y {
		return 0
	}
	if s.rank[x] > s.rank[y] {
		s.rank[x] += s.rank[y]
		s.par[y] = x
	} else {
		s.rank[y] += s.rank[x]
		s.par[x] = y
	}
	return -1
}
