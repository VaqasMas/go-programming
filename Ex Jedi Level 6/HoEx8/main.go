package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x())
}

func foo() func() string {
	return func() string {
		return "This is func under foo"
	}
}

// Didn't understoo what I had to do, so couldn't do this ex. on my own.
