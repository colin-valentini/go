package mergeksortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKSortedLists(t *testing.T) {
	want := makeListNode(t, []int{1, 1, 2, 3, 4, 4, 5, 6})
	assert.Equal(t, want, NewSolver([]*ListNode{
		makeListNode(t, []int{1, 4, 5}),
		makeListNode(t, []int{1, 3, 4}),
		makeListNode(t, []int{2, 6}),
	}).Solve())

	assert.Nil(t, NewSolver([]*ListNode{}).Solve())
	assert.Nil(t, NewSolver([]*ListNode{nil}).Solve())
}

func makeListNode(t *testing.T, values []int) *ListNode {
	t.Helper()
	head := &ListNode{}
	node := head
	for _, val := range values {
		node.Val = val
		node.Next = &ListNode{}
		node = node.Next
	}
	return head
}
