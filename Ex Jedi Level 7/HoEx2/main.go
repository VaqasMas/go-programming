package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(pf *person) {
	fmt.Println("------------")

	//can also do, pf.frist...
	(*pf).first = "Miss"
	(*pf).last = "Moneypenny"
	(*pf).age = 27

}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println("Before calling Change Me:", p1)

	changeMe(&p1)

	fmt.Println("After calling Change Me:", p1)
}
