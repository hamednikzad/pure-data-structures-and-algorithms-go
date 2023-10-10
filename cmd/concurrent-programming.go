package main

import (
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/concurrent-programming/mutexes"
	sync_once "github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/concurrent-programming/sync-once"
)

type concurrentProgramming struct {
}

func (p concurrentProgramming) mutexes() {
	mutexes.Run()
}

func (p concurrentProgramming) syncOnce() {
	sync_once.Run()
	sync_once.Run()
	sync_once.Run()
	sync_once.Run()
}
