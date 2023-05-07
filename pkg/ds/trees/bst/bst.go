package bst

import (
	"fmt"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds"
)

type Bst[T ds.Numeric] struct {
	Root *node[T]
}

func NewBst[T ds.Numeric]() *Bst[T] {
	return &Bst[T]{Root: nil}
}

func maxInt[T ds.Numeric](a, b T) T {
	if a > b {
		return a
	}
	return b
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

func getDepth[T ds.Numeric](parent *node[T]) int {
	if parent == nil {
		return 0
	}

	return maxInt(getDepth(parent.Left), getDepth(parent.Right)) + 1
}

func traversePreOrder[T ds.Numeric](parent *node[T]) {
	if parent == nil {
		return
	}

	fmt.Print(parent.Data, " ")
	traversePreOrder(parent.Left)
	traversePreOrder(parent.Right)
}

func traverseInOrder[T ds.Numeric](parent *node[T]) {
	for true {
		if parent == nil {
			return
		}
		traverseInOrder(parent.Left)
		fmt.Print(parent.Data, " ")
		parent = parent.Right
	}
}

func traversePostOrder[T ds.Numeric](parent *node[T]) {
	if parent == nil {
		return
	}

	traversePostOrder(parent.Left)
	traversePostOrder(parent.Right)
	fmt.Print(parent.Data, " ")
}

func (b *Bst[T]) GetDepth() int {
	return getDepth(b.Root)
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

// TraversePreOrder DLR
func (b *Bst[T]) TraversePreOrder() {
	fmt.Printf("TraversePreOrder(DLR) with Depth {%d}: ", b.GetDepth())
	traversePreOrder(b.Root)
	fmt.Println()
}

// TraverseInOrder LDR
func (b *Bst[T]) TraverseInOrder() {
	fmt.Printf("TraverseInOrder(LDR) with Depth {%d}: ", b.GetDepth())
	traverseInOrder(b.Root)
	fmt.Println()
}

// TraversePostOrder LRD
func (b *Bst[T]) TraversePostOrder() {
	fmt.Printf("TraversePostOrder(LRD) with Depth {%d}: ", b.GetDepth())
	traversePostOrder(b.Root)
	fmt.Println()
}
