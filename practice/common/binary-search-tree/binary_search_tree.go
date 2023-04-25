package binarysearchtree

// BinarySearchTree is a simple (non-comprehensive) implementation of the
// binary search tree data type. It's use for non-distinct values is untested.
// See https://en.wikipedia.org/wiki/Binary_search_tree.
type BinarySearchTree[T any] struct {
	root *BstNode[T]
	less func(x, y T) bool
}

type BstNode[T any] struct {
	parent      *BstNode[T]
	left, right *BstNode[T]
	value       T
}

func newBstNode[T any](value T) *BstNode[T] {
	return &BstNode[T]{value: value}
}

func NewBinarySearchTree[T comparable](less func(x, y T) bool) *BinarySearchTree[T] {
	return &BinarySearchTree[T]{less: less}
}

func (tree *BinarySearchTree[T]) Insert(value T) *BstNode[T] {
	var parent *BstNode[T]
	newNode := newBstNode(value)
	for node := tree.root; node != nil; {
		parent = node
		if tree.less(newNode.value, node.value) {
			node = node.left
		} else /* if tree.less(x.value, z.value) */ {
			node = node.right
		}
	}
	newNode.parent = parent
	if parent == nil {
		tree.root = newNode
	} else if tree.less(newNode.value, parent.value) {
		parent.left = newNode
	} else /* if tree.less(y.value, z.value) */ {
		parent.right = newNode
	}
	return newNode
}

func (tree *BinarySearchTree[T]) Find(target T) *BstNode[T] {
	node := tree.root
	for node != nil {
		if tree.less(target, node.value) == tree.less(node.value, target) {
			return node
		} else if tree.less(target, node.value) {
			node = node.left
		} else /* if tree.less(node.value, target) */ {
			node = node.right
		}
	}
	return nil
}

// Delete removes the given node from the binary search tree.
// See https://en.wikipedia.org/wiki/Binary_search_tree#Deletion.
func (tree *BinarySearchTree[T]) Delete(target *BstNode[T]) {
	if target.left == nil {
		tree.shiftNodes(target, target.right)
		return
	} else if target.right == nil {
		tree.shiftNodes(target, target.left)
		return
	} // else target has both left and right children.
	successor := tree.successor(target)
	if successor.parent != target {
		tree.shiftNodes(successor, successor.right)
		successor.right = target.right
		successor.right.parent = successor
	}
	tree.shiftNodes(target, successor)
	successor.left = target.left
	successor.left.parent = successor
}

func (tree *BinarySearchTree[T]) successor(node *BstNode[T]) *BstNode[T] {
	if node.right != nil {
		return tree.minimum(node.right)
	}
	parent := node.parent
	for parent != nil && node == parent.right {
		node = parent
		parent = parent.parent
	}
	return parent
}

func (tree *BinarySearchTree[T]) minimum(root *BstNode[T]) *BstNode[T] {
	node := root
	for node.left != nil {
		node = node.left
	}
	return node
}

func (tree *BinarySearchTree[T]) shiftNodes(u, v *BstNode[T]) {
	if u.parent == nil {
		tree.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else /* u == u.parent.right  */ {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

// Returns a BFS iterator which can be used to search all nodes.
func (tree *BinarySearchTree[T]) Iter() *BinarySearchTreeIterator[T] {
	q := newBstQueue[T]()
	q.push(tree.root)
	return &BinarySearchTreeIterator[T]{node: tree.root, q: q}
}

type BinarySearchTreeIterator[T any] struct {
	node *BstNode[T]
	q    *bstQueue[T]
}

func (it *BinarySearchTreeIterator[T]) next() bool {
	if it.q.len() == 0 {
		return false
	}
	it.node = it.q.pop()
	if it.node.left != nil {
		it.q.push(it.node.left)
	}
	if it.node.right != nil {
		it.q.push(it.node.right)
	}
	return true
}

func (it *BinarySearchTreeIterator[T]) value() *BstNode[T] {
	return it.node
}

type bstQueue[T any] struct {
	elems []*BstNode[T]
}

func newBstQueue[T any]() *bstQueue[T] {
	return &bstQueue[T]{elems: []*BstNode[T]{}}
}

func (q *bstQueue[T]) len() int {
	return len(q.elems)
}

func (q *bstQueue[T]) push(value *BstNode[T]) {
	q.elems = append(q.elems, value)
}

func (q *bstQueue[T]) pop() *BstNode[T] {
	// panics if called on empty queue
	value := q.elems[0]
	q.elems = q.elems[1:]
	return value
}
