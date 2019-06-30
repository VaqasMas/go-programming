package main

import "fmt"

func main() {
	xi1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	xi2 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	xs1 := foo(xi1...)
	fmt.Println("Result of foo func: ", xs1)
	xs2 := bar(xi2)
	fmt.Println("Result of bar func: ", xs2)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(y []int) int {
	sum := 0
	for _, v := range y {
		sum += v
	}
	return sum
}
