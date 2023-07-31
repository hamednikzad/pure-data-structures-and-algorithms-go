package main

import (
	"fmt"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/statcks/linkedliststack"
)

func stackUsage() {
	var stack = linkedliststack.New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Print()
	fmt.Println("********************")
	first, err := stack.Peek()
	if err != nil {
		fmt.Println("Error in Calling Peek: ", err.Error())
	} else {
		fmt.Println("First:\t", first)
	}
	fmt.Println("********************")

	fmt.Println("Iterating over Stack")
	for !stack.IsEmpty() {
		item, err := stack.Pop()
		if err != nil {
			fmt.Println("Error in Calling Pop: ", err.Error())
		} else {
			fmt.Print(item, ",\t")
		}
	}
	fmt.Println("\nall elements poped")
	stack.Print()
}
