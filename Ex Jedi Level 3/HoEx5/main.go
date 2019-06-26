package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		m := i % 4
		fmt.Printf("Number: %d & its Modulus: %d\n", i, m)
	}
}
