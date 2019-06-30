package main

import "fmt"

func main() {
	x := 45
	fmt.Println("Value:", x)
	fmt.Println("Address:", &x)

	y := &x
	fmt.Println("Value:", *y)
	fmt.Println("Address:", &y)

}
