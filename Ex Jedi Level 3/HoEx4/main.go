package main

import "fmt"

func main() {

	byear := 1984
	cyear := 2019

	for {
		if byear <= cyear {
			fmt.Println(byear)
			byear++
		} else {
			break
		}
	}
}
