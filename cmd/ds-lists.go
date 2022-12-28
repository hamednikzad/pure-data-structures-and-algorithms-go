package main

import (
	"fmt"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/lists/arraylist"
	genericlist "github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/lists/list"
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

func ListUsage() {
	//arrayListUsage()
	genericListUsage()
}
