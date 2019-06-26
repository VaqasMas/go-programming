package main

import "fmt"

func main() {

	switch {
	case false:
		fmt.Println("False")

	case (2 == 2):
		fmt.Println("2 is equal to 2")

	case true:
		fmt.Println("True")
	}
}
