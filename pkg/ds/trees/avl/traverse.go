package avl

import (
	"fmt"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds"
)

// TraversePreOrder DLR
func (b *Avl[T]) TraversePreOrder() {
	fmt.Printf("TraversePreOrder(DLR) with Depth {%d}: ", b.GetDepth())
	traversePreOrder(b.Root)
	fmt.Println()
}

// TraverseInOrder LDR
func (b *Avl[T]) TraverseInOrder() {
	fmt.Printf("TraverseInOrder(LDR) with Depth {%d}: ", b.GetDepth())
	traverseInOrder(b.Root)
	fmt.Println()
}

// TraversePostOrder LRD
func (b *Avl[T]) TraversePostOrder() {
	fmt.Printf("TraversePostOrder(LRD) with Depth {%d}: ", b.GetDepth())
	traversePostOrder(b.Root)
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

func (b *Avl[T]) GetDepth() int {
	return getDepth(b.Root)
}
