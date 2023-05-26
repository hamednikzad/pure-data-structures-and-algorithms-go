package main

import (
	"fmt"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/trees/avl"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/trees/bst"
)

func bstUsage() {
	b := bst.NewBst[int]()
	min, _ := b.Min()
	max, _ := b.Max()
	fmt.Printf("Min:%d, Max:%d\n", min, max)

	b.Add(10)
	b.Add(5)
	b.Add(15)
	b.Add(13)
	b.Add(12)
	b.Add(14)
	b.Add(13)
	fmt.Println("Depth: ", b.GetDepth())

	b.TraversePreOrder()
	b.TraverseInOrder()
	b.TraversePostOrder()

	b.Remove(13)
	min, _ = b.Min()
	max, _ = b.Max()
	fmt.Printf("Min:%d, Max:%d\n", min, max)

	b.TraversePreOrder()
}

func avlUsage() {
	b := avl.NewAvl[int]()
	/*min, _ := b.Min()
	max, _ := b.Max()
	fmt.Printf("Min:%d, Max:%d\n", min, max)*/

	b.Add(5)
	b.Add(3)
	b.Add(7)
	b.Add(2)
	b.Add(4)
	b.Add(13)
	b.Add(17)
	fmt.Println("Depth: ", b.GetDepth())

	b.TraversePreOrder()
	b.TraverseInOrder()
	b.TraversePostOrder()

	//b.Remove(13)
	//min, _ = b.Min()
	//max, _ = b.Max()
	//fmt.Printf("Min:%d, Max:%d\n", min, max)
	//
	//b.TraversePreOrder()
}

func TreeUsage() {
	//bstUsage()
	avlUsage()
}
