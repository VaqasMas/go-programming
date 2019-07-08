package main

import "fmt"

func main() {
	c := make(chan int)
	go send(c)

	receive(c)

	fmt.Println("Exit")

}

func send(s chan int) {
	for i := 1; i <= 100; i++ {
		s <- i
	}
	close(s)
}

func receive(r chan int) {
	for v := range r {
		fmt.Println("Pulled value:", v)
	}
}
