package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int

	const gs = 100

	var wg sync.WaitGroup

	fmt.Println("Start count:", counter)
	wg.Add(gs)
	for i := 1; i <= gs; i++ {
		go func() {
			v1 := counter
			runtime.Gosched()
			v1++
			counter = v1
			fmt.Println("Foo Counter:", counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Exit count:", counter)
}
