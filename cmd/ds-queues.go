package main

import (
	"fmt"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/queues"
)

func queueListUsage() {
	var queue = queues.New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	queue.Print()
	fmt.Println("********************")
	first, err := queue.Peek()
	if err != nil {
		fmt.Println("Error in Calling Peek: ", err.Error())
	} else {
		fmt.Println("First:\t", first)
	}
	fmt.Println("********************")

	fmt.Println("Iterating over Queue")
	for !queue.IsEmpty() {
		item, err := queue.Dequeue()
		if err != nil {
			fmt.Println("Error in Calling Dequeue: ", err.Error())
		} else {
			fmt.Print(item, ",\t")
		}
	}
	fmt.Println("\nall elements dequeued")
	queue.Print()
}
