package main

import (
	"fmt"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/lists/arraylist"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/lists/doublylinkedlist"
	genericlist "github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/lists/list"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/lists/singlylinkedlist"
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

func ListUsage() {
	//arrayListUsage()
	//genericListUsage()
	//singlyLinkedListUsage()
	doublyLinkedListUsage()
}
