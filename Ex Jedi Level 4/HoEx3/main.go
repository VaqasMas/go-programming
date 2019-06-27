package main

import "fmt"

func main() {
	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("Main Slice: ", sl)
	fmt.Println("Slice One: ", sl[:5])
	fmt.Println("Slice Two: ", sl[5:])
	fmt.Println("Slice Three: ", sl[2:7])
	fmt.Println("Slice Four: ", sl[1:6])

}
