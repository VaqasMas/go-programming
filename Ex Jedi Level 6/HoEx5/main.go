package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	sa := s.length * s.length
	return sa
}

func (c circle) area() float64 {
	ca := math.Pi * c.radius * c.radius
	return ca
}

type shape interface {
	area() float64
}

func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {

	s1 := square{
		length: 21.2,
	}

	c1 := circle{
		radius: 13.3,
	}

	info(s1)
	info(c1)

}
