package main

import "github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/concurrent-programming/mutexes"

type concurrentProgramming struct {
}

func (p concurrentProgramming) mutexes() {
	mutexes.Run()
}
