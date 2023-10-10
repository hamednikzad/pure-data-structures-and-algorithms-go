package mutexes

import (
	"fmt"
	"sync"
)

func Run() {
	var arr []int
	var wg sync.WaitGroup
	var m sync.Mutex

	const iterations = 100
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			arr = append(arr, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(len(arr))
}
