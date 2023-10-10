package main

import (
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/concurrent-programming/mutexes"
	synconce "github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/concurrent-programming/sync-once"
)

type concurrentProgramming struct {
}

func (p concurrentProgramming) mutexes() {
	mutexes.Run()
}

func (p concurrentProgramming) syncOnce() {
	synconce.Run()
	synconce.Run()
	synconce.Run()
	synconce.Run()
}
