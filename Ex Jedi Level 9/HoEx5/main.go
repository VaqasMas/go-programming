package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter int64

const gs = 100

var wg sync.WaitGroup

func main() {
	fmt.Println("Start count:", counter)
	wg.Add(gs)
	for i := 1; i <= gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Exit count:", counter)
}

// Had to watch the lecture about atomic again.
