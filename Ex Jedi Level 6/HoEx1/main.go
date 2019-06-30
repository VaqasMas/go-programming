package main

import "fmt"

func main() {

	x := foo()
	y, z := bar()

	fmt.Println(x)
	fmt.Printf("%v\n%v", y, z)

}

func foo() int {
	return 42
}

func bar() (int, string) {

	return 45, "This is bar"
}
