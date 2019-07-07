package main

import (
	"fmt"
	"sync"
)

var counter int

const gs = 100

var wg sync.WaitGroup

var mu sync.Mutex

func main() {
	fmt.Println("Start count:", counter)
	wg.Add(gs)
	for i := 1; i <= gs; i++ {
		go func() {
			mu.Lock()
			v1 := counter
			v1++
			counter = v1
			fmt.Println("Counter:", counter)
			mu.Unlock()
			//fmt.Println("Foo Counter:", counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Exit count:", counter)
}
