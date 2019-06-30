package main

import "fmt"

func main() {
	p1 := struct {
		first  string
		comp   map[string]string
		falvor []string
	}{
		first: "James",
		comp: map[string]string{
			"name":     "ABC",
			"location": "123",
		},
		falvor: []string{
			"chocolate",
			"strawberry",
		},
	}

	fmt.Println(p1)

	fmt.Println(p1.comp)
}

/*
	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "Miss",
		last:  "Moneypenny",
		age:   28,
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
*/
