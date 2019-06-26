package main

import "fmt"

func main() {

	favSport := "Basketball"
	switch favSport {
	case "Baseball":
		fmt.Println("My Favourite game is Baseball")

	case "Basketball":
		fmt.Println("My Favourite game is Basketball")

	case "Golf":
		fmt.Println("My Favourite game is Golf")

	default:
		fmt.Println("I don't like sports")
	}
}
