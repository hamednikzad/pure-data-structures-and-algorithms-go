package contexes

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func goFunc(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	for range time.Tick(500 * time.Millisecond) {
		if ctx.Err() != nil {
			fmt.Println(ctx.Err())
			return
		}
		fmt.Println("Tick!")
	}
}

func RunWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)

	go goFunc(&wg, ctx)

	time.Sleep(2 * time.Second)
	cancel()

	wg.Wait()
}

func RunWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go goFunc(&wg, ctx)

	wg.Wait()
}
