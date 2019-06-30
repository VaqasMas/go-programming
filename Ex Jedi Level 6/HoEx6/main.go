package main

import "fmt"

func main() {

	func() {
		fmt.Println("This is Anonymous function")
	}()

	func(n int) {
		fmt.Printf("%d\n", n)
	}(42)
}
