package leetcode

// LeetCode #133.
// Link: https://leetcode.com/problems/clone-graph/

// Given a reference of a node in a connected undirected graph.
// Return a deep copy (clone) of the graph.
// Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

// class Node {
//     public int val;
//     public List<Node> neighbors;
// }

// Test case format:

// For simplicity, each node's value is the same as the node's index (1-indexed). For example,
// the first node with val == 1, the second node with val == 2, and so on. The graph is represented
// in the test case using an adjacency list.

// An adjacency list is a collection of unordered lists used to represent a finite graph. Each list
// describes the set of neighbors of a node in the graph.

// The given node will always be the first node with val = 1. You must return the copy of the given
// node as a reference to the cloned graph.

// Example 1:
// Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
// Output: [[2,4],[1,3],[2,4],[1,3]]
// Explanation: There are 4 nodes in the graph.
// 1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
// 2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
// 3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
// 4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).

// Example 2:
// Input: adjList = [[]]
// Output: [[]]
// Explanation: Note that the input contains one empty list. The graph consists of only one node with
// val = 1 and it does not have any neighbors.

// Example 3:
// Input: adjList = []
// Output: []
// Explanation: This an empty graph, it does not have any nodes.

// Example 4:
// Input: adjList = [[2],[1]]
// Output: [[2],[1]]

// Constraints:
// - The number of nodes in the graph is in the range [0, 100].
// - 1 <= Node.val <= 100
// - Node.val is unique for each node.
// - There are no repeated edges and no self-loops in the graph.
// - The Graph is connected and all nodes can be visited starting from the given node.

type Node struct {
	Val       int
	Neighbors []*Node
}

// CloneGraph is a solution to the clone graph problem.
func CloneGraph(node *Node) *Node {
	return deepClone(node, make(map[int]*Node))
}

// deepClone performs a deep clone starting from the given node and
// uses the given clones map to cache allocations.
func deepClone(node *Node, clones map[int]*Node) *Node {
	// There should be no nil nodes, but proceeding without this
	// guard would be foolish in practice. We also have an early
	// return if we've already allocated a copy of the given node
	if node == nil {
		return nil
	} else if clone, ok := clones[node.Val]; ok {
		return clone
	}

	// Allocate a new node and stash it in our index
	clone := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0, len(node.Neighbors)),
	}
	clones[node.Val] = clone

	// Recursively clone neighboring nodes
	for _, neighborNode := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, deepClone(neighborNode, clones))
	}

	return clone
}
