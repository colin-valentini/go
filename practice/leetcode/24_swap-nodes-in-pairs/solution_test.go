package swapnodesinpairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	head *ListNode
	want []*ListNode
}

func TestSolver(t *testing.T) {
	testCases := []TestCase{
		makeTestCase(t, []int{}, []int{}),
		makeTestCase(t, []int{1}, []int{1}),
		makeTestCase(t, []int{1, 2}, []int{2, 1}),
		makeTestCase(t, []int{1, 2, 3}, []int{2, 1, 3}),
		makeTestCase(t, []int{1, 2, 3, 4}, []int{2, 1, 4, 3}),
		makeTestCase(t, []int{1, 2, 3, 4, 5}, []int{2, 1, 4, 3, 5}),
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.head)
		got := solver.Solve()
		if len(testCase.want) == 0 {
			assert.Nil(t, got, "Failed test case %d: wanted nil", i+1)
			break
		}
		gotNode := got
		for j, wantNode := range testCase.want {
			assert.Same(t, wantNode, gotNode, "Failed test case %d at node %d", i+1, j+1)
			gotNode = gotNode.Next
		}
	}
}

func makeTestCase(t *testing.T, in, out []int) TestCase {
	t.Helper()
	require.Equal(t, len(in), len(out), "Invalid test case: len(in) != len(out)")
	if len(in) == 0 {
		return TestCase{head: nil, want: []*ListNode{}}
	}
	m := map[int]*ListNode{} // Keep track of node pointers by value.
	head := &ListNode{}
	node := head
	for i, v := range in {
		// Set the current node's value.
		node.Val = v
		m[v] = node
		if i == len(in)-1 {
			continue
		}
		// Allocate and advance to the next element.
		node.Next = &ListNode{}
		node = node.Next
	}
	want := make([]*ListNode, len(out))
	for i, v := range out {
		node, ok := m[v]
		require.True(t, ok, "Invalid test case: no in node with value %d", v)
		want[i] = node
	}
	return TestCase{head: head, want: want}
}
