package main

import "fmt"

type vehicle struct {
	door  string
	color string
}
type sedan struct {
	vehicle
	luxury bool
}
type truck struct {
	vehicle
	fourWheel bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			door:  "two",
			color: "blue",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			door:  "four",
			color: "black",
		},
		luxury: false,
	}

	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Println("Truck: ", t1.door, t1.color, t1.fourWheel)
	fmt.Println("Sedan: ", s1.door, s1.color, s1.luxury)
	//fmt.Println(t1.vehicle, t1.fourWheel)
}
