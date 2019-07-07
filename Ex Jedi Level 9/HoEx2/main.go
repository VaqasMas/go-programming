package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() *person {
	*p = person{"Masood"}
	return p
}

type human interface {
	speak() *person
}

func main() {

	ps := person{"vaqas"}
	fmt.Println(ps)
	saySomething(&ps)
	// Without pointer below calling doesn't work.
	//saySomething(ps)
	//Below one will work.
	//fmt.Println(ps.speak())

}

func saySomething(h human) {

	fmt.Println("I'm", h.speak())
}
