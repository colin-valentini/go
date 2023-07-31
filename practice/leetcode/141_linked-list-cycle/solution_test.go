package linkedlistcycle

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		head *ListNode
		want bool
	}{
		{head: makeLinkedList(t, []int{3, 2, 0, -4}, 1), want: true},
		{head: makeLinkedList(t, []int{1, 2}, 0), want: true},
		{head: makeLinkedList(t, []int{1}, -1), want: false},
		{head: makeLinkedList(t, []int{1, 2, 3, 4}, -1), want: false},
		{head: makeLinkedList(t, []int{1}, 0), want: true},
		{head: makeLinkedList(t, []int{}, -1), want: false},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.head)
		got := solver.Solve()
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}

func makeLinkedList(t *testing.T, values []int, cycleIndex int) *ListNode {
	t.Helper()
	if len(values) == 0 {
		return nil
	}
	nodes := make([]*ListNode, len(values))
	for i, v := range values {
		nodes[i] = &ListNode{Val: v}
		if i-1 >= 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	if cycleIndex >= 0 {
		require.Less(t, cycleIndex, len(nodes), "Invalid test case params")
		nodes[len(nodes)-1].Next = nodes[cycleIndex]
	}
	return nodes[0]
}
