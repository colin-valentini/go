package removenthnodefromendoflist

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		head *ListNode
		n    int
		want []int
	}{
		{
			head: toLinkedList(t, []int{1, 2, 3, 4, 5}),
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			head: toLinkedList(t, []int{1}),
			n:    1,
			want: []int{},
		},
		{
			head: toLinkedList(t, []int{1, 2}),
			n:    1,
			want: []int{1},
		},
		{
			head: toLinkedList(t, []int{1, 2, 3}),
			n:    1,
			want: []int{1, 2},
		},
		{
			head: toLinkedList(t, []int{1, 2, 3}),
			n:    2,
			want: []int{1, 3},
		},
		{
			head: toLinkedList(t, []int{1, 2, 3}),
			n:    3,
			want: []int{2, 3},
		},
	}
	for i, testCase := range testCases {
		solver := NewSolver(testCase.head, testCase.n)
		got := toList(t, solver.Solve())
		require.Equal(t, len(testCase.want), len(got))
		assert.Equal(t, testCase.want, got, "Failed test case %d", i+1)
	}
}

func toLinkedList(t *testing.T, list []int) *ListNode {
	t.Helper()
	require.Greater(t, len(list), 0)
	head := &ListNode{}
	node := head
	for i, v := range list {
		node.Val = v
		if i < len(list)-1 {
			node.Next = &ListNode{}
			node = node.Next
		}
	}
	return head
}

func toList(t *testing.T, head *ListNode) []int {
	t.Helper()
	list := []int{}
	for node := head; node != nil; node = node.Next {
		list = append(list, node.Val)
	}
	return list
}
