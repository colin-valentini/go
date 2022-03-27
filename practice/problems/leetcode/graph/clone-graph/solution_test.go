package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCloneGraph runs test cases for CloneGraph.
func TestCloneGraph(t *testing.T) {
	runTestCase(t, [][]int{{}})
	runTestCase(t, [][]int{})
	runTestCase(t, [][]int{{2}, {1}})
	runTestCase(t, [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
}

func runTestCase(t *testing.T, adjList [][]int) {
	expected := newGraph(t, adjList)
	actual := CloneGraph(expected)
	assertGraphEquality(t, expected, actual, map[int]struct{}{})
}

// newGraph creates a new graph data structure from the given adjcency
// list in two passes. The first pass allocates each node in memory,
// while the second pass establishes the edges between nodes in the graph.
// This constructor assumes 1-based indexing on the node values (see
// the problem description for more details).
func newGraph(t *testing.T, adjList [][]int) *Node {
	t.Helper()
	nodesByVal := map[int]*Node{}
	for i, neighborVals := range adjList {
		nodeVal := i + 1
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
	if a == nil {
		require.Nil(t, b, "a is nil, but b is not")
		return
	} else if b == nil {
		assert.Fail(t, "a is not nil, but b is nil")
	}
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
