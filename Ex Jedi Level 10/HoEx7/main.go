package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	const gr = 10

	wg.Add(gr)

	c := make(chan int)
	//fanout := make(chan int)

	for i := 0; i < gr; i++ {
		fmt.Println("Routine", i)
		go send(c)
		go receive(c)
	}

	fmt.Println("Exit")
	wg.Wait()

}

func send(s chan int) {

	for j := 0; j < 10; j++ {
		s <- j
	}
	//close(s)
	wg.Done()
}

func receive(r chan int) {

	for v := range r {
		fmt.Println(v)
	}
	//wg.Done()
}

//Didn't do this exercise correctly. Solutions are total different.
