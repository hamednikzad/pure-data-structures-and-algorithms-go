package main

import (
	"fmt"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/lists/arraylist"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/lists/doublylinkedlist"
	genericlist "github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/lists/list"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/lists/singlylinkedlist"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/lists/sortedlinkedlist"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/queues"
)

func arrayListUsage() {
	list := arraylist.New(2)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Print()

	fmt.Println("Is contains 3: ", list.Contains(3))

	err := list.Insert(2, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	list.Print()
}

func genericListUsage() {
	ls := genericlist.New[int](2)
	ls.Add(1)
	ls.Add(2)
	ls.Add(3)
	ls.Add(4)
	ls.Add(5)
	ls.Print()

	fmt.Println("Is contains 3: ", ls.Contains(3))

	err := ls.Insert(2, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	ls.Print()
}

func singlyLinkedListUsage() {
	var ll = singlylinkedlist.New[int]()
	ll.AddLast(5)
	ll.AddLast(2)
	ll.AddLast(3)

	ll.Print()
	fmt.Println("********************")
	ll.Traverse(fmt.Print)
	fmt.Println()
	fmt.Println("********************")
	var last = ll.GetValueAtIndex(2)
	fmt.Println("Last:\t", last)

	fmt.Println("********************")
	ll.Remove(1)
	ll.Print()
}

func doublyLinkedListUsage() {
	var ll = doublylinkedlist.New[int]()
	ll.AddLast(5)
	ll.AddLast(2)
	ll.AddLast(3)

	ll.Print()
	fmt.Println("********************")
	ll.Traverse(fmt.Print)
	fmt.Println()
	fmt.Println("********************")
	var last = ll.GetValueAtIndex(2)
	fmt.Println("Last:\t", last)

	fmt.Println("********************")
	ll.Remove(1)
	ll.Print()
}

func sortedLinkedListUsage() {
	var ll = sortedlinkedlist.New[int]()
	ll.Add(5)
	ll.Add(2)
	ll.Add(3)

	ll.Print()
	fmt.Println("********************")
	ll.Traverse(fmt.Print)
	fmt.Println()
	fmt.Println("********************")
	var last = ll.GetValueAtIndex(2)
	fmt.Println("Last:\t", last)

	fmt.Println("********************")
	ll.RemoveValue(2)
	ll.Print()
}

func queueListUsage() {
	var queue = queues.New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	queue.Print()
	fmt.Println("********************")
	err, first := queue.Peek()
	if err != nil {
		fmt.Println("Error in Calling Peek: ", err.Error())
	} else {
		fmt.Println("First:\t", first)
	}
	fmt.Println("********************")

	fmt.Println("Iterating over Queue")
	for !queue.IsEmpty() {
		err, item := queue.Dequeue()
		if err != nil {
			fmt.Println("Error in Calling Dequeue: ", err.Error())
		} else {
			fmt.Print(item, ",\t")
		}
	}
	fmt.Println("\nall elements dequeued")
	queue.Print()
}

func ListUsage() {
	//arrayListUsage()
	//genericListUsage()
	//singlyLinkedListUsage()
	//doublyLinkedListUsage()
	//queueListUsage()
	sortedLinkedListUsage()
}
