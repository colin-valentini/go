package mergeksortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKSortedLists(t *testing.T) {
	want := makeListNode(t, []int{1, 1, 2, 3, 4, 4, 5, 6})
	got := NewSolver([]*ListNode{
		makeListNode(t, []int{1, 4, 5}),
		makeListNode(t, []int{1, 3, 4}),
		makeListNode(t, []int{2, 6}),
	}).Solve()
	assert.Equal(t, want, got)

	// assert.Nil(t, NewSolver([]*ListNode{}).Solve())
	// assert.Nil(t, NewSolver([]*ListNode{nil}).Solve())
}

func makeListNode(t *testing.T, values []int) *ListNode {
	t.Helper()
	head := &ListNode{}
	node := head
	for _, val := range values {
		node.Next = &ListNode{Val: val}
		node = node.Next
	}
	return head.Next
}
