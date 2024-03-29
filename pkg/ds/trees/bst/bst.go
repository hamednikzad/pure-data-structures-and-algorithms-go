package bst

import (
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds"
)

type Bst[T ds.Numeric] struct {
	Root *node[T]
}

func NewBst[T ds.Numeric]() *Bst[T] {
	return &Bst[T]{Root: nil}
}

func min[T ds.Numeric](node *node[T]) T {
	min := node.Data

	for node.Left != nil {
		min = node.Left.Data
		node = node.Left
	}

	return min
}

func max[T ds.Numeric](node *node[T]) T {
	max := node.Data

	for node.Right != nil {
		max = node.Right.Data
		node = node.Right
	}

	return max
}

func remove[T ds.Numeric](parent *node[T], data T) *node[T] {
	if parent == nil {
		return nil
	}

	if data < parent.Data {
		parent.Left = remove(parent.Left, data)
	} else if data > parent.Data {
		parent.Right = remove(parent.Right, data)
	} else { // node should be removed
		// node with only one child or no child
		if parent.Left == nil {
			return parent.Right
		}
		if parent.Right == nil {
			return parent.Left
		}
		// node with two children: Get the inorder successor (smallest in the right subtree)
		parent.Data = min(parent.Right)

		// Delete the inorder successor
		parent.Right = remove(parent.Right, parent.Data)
	}

	return parent
}

func (b *Bst[T]) Min() (T, bool) {
	if b.Root == nil {
		return 0, false
	}
	return min(b.Root), true
}

func (b *Bst[T]) Max() (T, bool) {
	if b.Root == nil {
		return 0, false
	}
	return max(b.Root), true
}

func (b *Bst[T]) Remove(data T) {
	b.Root = remove(b.Root, data)
}

func (b *Bst[T]) Add(data T) bool {
	n := newNode(data)
	if b.Root == nil {
		b.Root = n
		return true
	}

	next := b.Root
	prev := next

	for next != nil {
		prev = next
		if data < next.Data {
			next = next.Left
		} else if data > next.Data {
			next = next.Right
		} else {
			return false //Exists
		}
	}

	if data < prev.Data {
		prev.Left = n
	} else {
		prev.Right = n
	}

	return true
}
