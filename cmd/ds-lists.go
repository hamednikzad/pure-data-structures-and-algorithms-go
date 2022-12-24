package main

import (
	"fmt"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/lists/arraylist"
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

func ListUsage() {
	arrayListUsage()
}
