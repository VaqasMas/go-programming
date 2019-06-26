package main

import (
	"fmt"
)

func main() {
	//var s string
	var i int

	for i = 22; i <= 122; i++ {
		//s = string(i)
		fmt.Printf("ASCII: %#U\n", i)
	}
}
