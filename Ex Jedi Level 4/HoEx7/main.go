package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	xy := [][]string{x, y}
	for i, v := range xy {
		fmt.Println("Record: ", i)

		for j, z := range v {
			fmt.Printf("\t")
			fmt.Println("Index: ", j, "& It's Value: ", z)
		}
	}

}
