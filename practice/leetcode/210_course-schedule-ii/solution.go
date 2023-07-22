package leetcode

// LeetCode #210.
// Link: https://leetcode.com/problems/course-schedule-ii/

// There are a total of numCourses courses you have to take, labeled from 0 to
// numCourses - 1. You are given an array prerequisites where
// prerequisites[i] = [ai, bi] indicates that you must take course bi first if
// you want to take course ai.

// For example, the pair [0, 1], indicates that to take course 0 you have to
// first take course 1. Return the ordering of courses you should take to finish
// all courses. If there are many valid answers, return any of them. If it is
// impossible to finish all courses, return an empty array.

// Example 1.
// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: [0,1]
// Explanation: There are a total of 2 courses to take. To take course 1 you
// should have finished course 0. So the correct course order is [0,1].

// Example 2.
// Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
// Output: [0,2,1,3]
// Explanation: There are a total of 4 courses to take. To take course 3 you
// should have finished both courses 1 and 2. Both courses 1 and 2 should be
// taken after you finished course 0.
// So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

// Example 3.
// Input: numCourses = 1, prerequisites = []
// Output: [0]

// Constraints:
// - 1 <= numCourses <= 2000
// - 0 <= prerequisites.length <= numCourses * (numCourses - 1)
// - prerequisites[i].length == 2
// - 0 <= ai, bi < numCourses
// - ai != bi
// - All the pairs [ai, bi] are distinct.

// FindOrder is a solution to the course schedule ii problem.
func FindOrder(numCourses int, prerequisites [][]int) []int {
	return NewSolver(numCourses, prerequisites).Solve()
}

type Solver struct {
	// The total number of courses (nodes), abbreviated.
	n int

	// An adjcency list mapping each node to
	// a list of nodes that "depend on it".
	// rdeps[A] = [B, C] means that B and C both
	// require A, where A, B, C are valid node ints.
	rdeps [][]int

	// Each element in this slice maintains the
	// current number of remaining (unfulfilled)
	// dependencies.
	// indegree[A] = X means that A requires X
	// nodes to be fulfilled. If X == 0 then A
	// can be processed immediately.
	indegree []int
}

func NewSolver(n int, edges [][]int) *Solver {
	rdeps := make([][]int, n)
	indegree := make([]int, n)
	for _, edge := range edges {
		node, dep := edge[0], edge[1]
		rdeps[dep] = append(rdeps[dep], node)
		indegree[node]++
	}
	return &Solver{
		n:        n,
		rdeps:    rdeps,
		indegree: indegree,
	}
}

func (s *Solver) Solve() []int {
	// Allocate storage for the topological sort ordering.
	order := make([]int, 0, s.n)

	// Allocate queue storage, and initialize it with
	// any nodes that have no dependencies.
	// If a cycle exists, each node in the cycle will
	// have a non-zero indegree, but we can still have
	// a node with zero indegree that is a dependency
	// of another node in a cycle.
	q := make(queue, 0, s.n)
	for node, numDeps := range s.indegree {
		if numDeps == 0 {
			q.push(node)
		}
	}

	// Process each node in the queue. We only ever enqueue
	// a node if all of its dependencies have been met.
	// We never process the same node twice as a reuslt.
	for len(q) > 0 {
		node := q.pop()
		order = append(order, node)

		// Decrement the indegree of all nodes that depend
		// on the current node. If the indegree for any of
		// these becomes zero, we know it has no remaining
		// dependencies and can processed by enqueueing it.
		for _, rdep := range s.rdeps[node] {
			s.indegree[rdep]--
			if s.indegree[rdep] == 0 {
				q.push(rdep)
			}
		}
	}

	// Ordering will not be complete (have all nodes)
	// in the event of a cycle. We might get nothing,
	// or we might have a some nodes outside the cycle.
	// The problem statement requests an empty result
	// in this case.
	if len(order) != s.n {
		return []int{}
	}

	return order
}

type queue []int

func (q *queue) push(value int) {
	*q = append(*q, value)
}

func (q *queue) pop() int {
	// panics if called on an empty queue
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}
