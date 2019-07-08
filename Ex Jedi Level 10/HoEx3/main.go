package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

//Didn't know  how to resolve deadlock, only after watching solution
// I realized how to resolve it.
func gen() <-chan int {
	c := make(chan int)
	go func() { // did this after watching solution
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

//Did this myself, but was facing deadlock.
func receive(r <-chan int) {
	for v := range r {
		fmt.Println(v)
	}
}
