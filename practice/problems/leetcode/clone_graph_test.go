package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloneGraph(t *testing.T) {
	runTestCase(t, [][]int{{}})
	runTestCase(t, [][]int{})
	runTestCase(t, [][]int{{2}, {1}})
	runTestCase(t, [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
}

func runTestCase(t *testing.T, adjList [][]int) {
	expected := newGraph(adjList)
	actual := cloneGraph(expected)
	assertGraphEquality(t, expected, actual, map[int]struct{}{})
}

// newGraph creates a new graph data structure from the given adjcency 
// list in two passes. The first pass allocates each node in memory, 
// while the second pass establishes the edges between nodes in the graph. 
// This constructor assumes 1-based indexing on the node values (see 
// the problem description for more details).
func newGraph(adjList [][]int) *Node {
	nodesByVal := map[int]*Node{}
	for i, neighborVals := range adjList {
		nodeVal := i+1
		nodesByVal[nodeVal] = &Node{
			Val:       nodeVal,
			Neighbors: make([]*Node, len(neighborVals)),
		}
	}

	for i, neighborVals := range adjList {
		node := nodesByVal[i+1]
		for j, val := range neighborVals {
			node.Neighbors[j] = nodesByVal[val]
		}
	}

	return nodesByVal[1]
}

// assertGraphEquality runs assertions to test that graph b is a true copy of graph a.
func assertGraphEquality(t *testing.T, a, b *Node, visited map[int]struct{}) {
	if _, ok := visited[a.Val]; ok {
		return
	}
	assert.NotSame(t, a, b)
	assert.Equal(t, a.Val, b.Val)
	assert.Equal(t, len(a.Neighbors), len(b.Neighbors))
	visited[a.Val] = struct{}{}
	for i, expectedNeighbor := range a.Neighbors {
		actualNeighbor := b.Neighbors[i]
		assertGraphEquality(t, expectedNeighbor, actualNeighbor, visited)
	}
}
