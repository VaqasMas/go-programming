package main

import "fmt"

const (
	one   = 2017 + iota
	two   = 2017 + iota
	three = 2017 + iota
	four  = 2017 + iota
)

func main() {
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	fmt.Println(four)
}
