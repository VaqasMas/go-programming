package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	icecream  []string
}

func main() {
	p1 := person{
		firstname: "James",
		lastname:  "Bond",
		icecream:  []string{"chocolate", "hazelnuts"},
	}

	p2 := person{
		firstname: "Miss",
		lastname:  "Moneypenny",
		icecream:  []string{"strawberry", "vanilla"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for i, v := range p1.icecream {
		fmt.Println(i, v)
	}
	for i, v := range p2.icecream {
		fmt.Println(i, v)
	}
}
