package sync_once

import (
	"fmt"
	"sync"
)

var o sync.Once

func Run() {
	o.Do(func() {
		fmt.Println("sync_once func Run()")
	})
}
