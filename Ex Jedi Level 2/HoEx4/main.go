package main

import "fmt"

func main() {
	num := 42

	fmt.Printf("%d\t	%b\t	%#x\n", num, num, num)

	bum := num << 1

	fmt.Printf("%d\t	%b\t	%#x\n", bum, bum, bum)
}
