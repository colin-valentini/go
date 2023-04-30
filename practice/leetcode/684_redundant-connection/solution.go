package redundantconnection

// LeetCode #684.
// https://leetcode.com/problems/redundant-connection/

// In this problem, a tree is an undirected graph that is connected and has no
// cycles.

// You are given a graph that started as a tree with n nodes labeled from 1 to
// n, with one additional edge added. The added edge has two different vertices
// chosen from 1 to n, and was not an edge that already existed. The graph is
// represented as an array edges of length n where edges[i] = [ai, bi] indicates
// that there is an edge between nodes ai and bi in the graph.

// Return an edge that can be removed so that the resulting graph is a tree of n
// nodes. If there are multiple answers, return the answer that occurs last in
// the input.

// Example 1:
// Input: edges = [[1,2],[1,3],[2,3]]
// Output: [2,3]

// Example 2:
// Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
// Output: [1,4]

// Constraints:
// - n == edges.length
// - 3 <= n <= 1000
// - edges[i].length == 2
// - 1 <= a_i < b_i <= edges.length
// - a_i != b_i
// - There are no repeated edges.
// - The given graph is connected.

// func findRedundantConnection(edges [][]int) []int {
//
// }

type Solver struct {
	edges   [][]int
	parents []int
	ranks   []int
}

func NewSolver(edges [][]int) *Solver {
	return &Solver{edges: edges}
}

// Solve solves the Redundant Connection problem.
//
// This solution uses the UnionFind algorithm to gradually build the graph
// into a representation of "parent" nodes (canonical representatives which
// group each connected sub-graph), and a list of the size of each sub-graph.
// When we try to union two edges which are *already* connected (i.e they have
// the same canonical "parent" representative), we know that the edge is
// unnecessary as the two nodes are already connected in the same sub-graph.
//
// Time: O(N*alpha(N)) where alpha is the Inverse-Ackerman function.
// Space: O(N).
func (s *Solver) Solve() []int {
	s.init()
	for _, edge := range s.edges {
		x, y := edge[0]-1, edge[1]-1 // Convert 1-based to 0-based
		if s.union(x, y) {
			return edge
		}
	}
	// Unreachable based on problem constraints (given n edges and n nodes,
	// there will always be a redundant edge).
	return nil
}

func (s *Solver) init() {
	n := len(s.edges)
	s.parents = make([]int, n)
	s.ranks = make([]int, n)
	for i := 0; i < n; i++ {
		s.parents[i] = i
		s.ranks[i] = 1
	}
}

func (s *Solver) find(node int) int {
	if s.parents[node] == node {
		return node
	}
	s.parents[node] = s.find(s.parents[node])
	return s.parents[node]
}

func (s *Solver) union(x, y int) bool {
	x, y = s.find(x), s.find(y)
	if x == y {
		// We've been asked to union two nodes which are already connected
		// (since they have the same canonical "parent"). This edge is
		// redundant.
		return true
	}
	// Merge the two sub-graphs ("connect" them) by merging them into the
	// larger of the two sub-graphs.
	if s.ranks[x] > s.ranks[y] {
		s.ranks[x] += s.ranks[y]
		s.parents[y] = x
	} else {
		s.ranks[y] += s.ranks[x]
		s.parents[x] = y
	}
	return false
}
