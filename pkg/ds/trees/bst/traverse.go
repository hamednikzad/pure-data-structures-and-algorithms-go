package bst

import (
	"fmt"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/queues"
)

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

// TraverseLevelOrder Level
func (b *Bst[T]) TraverseLevelOrder() {
	fmt.Printf("TraverseLevelOrder(LRD) with Depth {%d}: ", b.GetDepth())
	traverseLevelOrder(b.Root)
	fmt.Println()
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

func traverseLevelOrder[T ds.Numeric](parent *node[T]) {
	if parent == nil {
		return
	}
	q := queues.New[*node[T]]()
	q.Enqueue(parent)

	for !q.IsEmpty() {
		r, _ := q.Dequeue()
		fmt.Print(r.Data, " ")

		if r.Left != nil {
			q.Enqueue(r.Left)
		}

		if r.Right != nil {
			q.Enqueue(r.Right)
		}
	}
}

func (b *Bst[T]) GetDepth() int {
	return getDepth(b.Root)
}
