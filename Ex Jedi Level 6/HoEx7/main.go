package main

import "fmt"

func main() {
	xi := []int{2, 4, 5, 6, 4, 7}

	xf := func(x []int) {
		for i, v := range x {
			fmt.Printf("Index:%v\t Value:%v\n", i, v)
		}
	}
	xf(xi)
}
