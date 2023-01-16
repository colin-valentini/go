package numberofnodesinthesubtreewiththesamelabel

// LeetCode #1519.
// https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/

// func countSubTrees(n int, edges [][]int, labels string) []int {
//
// }

type Solver struct {
	n      int
	edges  [][]int
	labels string
	adj    [][]int
	res    []int
}

func NewSolver(n int, edges [][]int, labels string) *Solver {
	return &Solver{n: n, edges: edges, labels: labels}
}

// Solve solves the Number of Nodes in the Sub-Tree With the Same Label problem.
// Time: O(N).
// Space: O(26*N).
func (s *Solver) Solve() []int {
	// Create an adjacency list from the given edges. Then traverse the tree
	// in a depth-first manner from the root. At each node, we ask each child
	// node for their character counts, and return the aggregate to the caller.
	// These character count arrays (of size 26 for each lower-case English
	// character), make it trivial for each node to compute their own answer.
	s.res = make([]int, s.n)
	s.adj = s.adjacencyList()
	_ = s.dfs(0, -1)
	return s.res
}

func (s *Solver) adjacencyList() [][]int {
	adj := make([][]int, s.n)
	for _, pair := range s.edges {
		x, y := pair[0], pair[1]
		if adj[x] == nil {
			adj[x] = make([]int, 0, 3)
		}
		if adj[y] == nil {
			adj[y] = make([]int, 0, 3)
		}
		adj[x] = append(adj[x], y)
		adj[y] = append(adj[y], x)
	}
	return adj
}

func (s *Solver) dfs(node, parent int) [26]int {
	var counts [26]int
	counts[s.labels[node]-'a']++

	neighbors := s.adj[node]
	for _, neighbor := range neighbors {
		if neighbor == parent {
			continue
		}
		for i, count := range s.dfs(neighbor, node) {
			counts[i] += count
		}
	}
	s.res[node] = counts[s.labels[node]-'a']
	return counts
}
