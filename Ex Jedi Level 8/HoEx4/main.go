package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println("Before:", xi)
	sort.Ints(xi)
	fmt.Println("After:", xi)
	fmt.Println("---------------")
	fmt.Println("Before:", xs)
	sort.Strings(xs)
	fmt.Println("After:", xs)
}
