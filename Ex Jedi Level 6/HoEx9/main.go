package main

import "fmt"

func main() {

	a := average(sum, []int{2, 4, 5, 3, 2, 6})
	fmt.Printf("Average result:%f\n", a)

}

func sum(x []int) float64 {
	total := 0.
	for _, v := range x {
		total += float64(v)
	}
	return total
}

func average(f func(sx []int) float64, y []int) float64 {
	var sx []int
	for _, v := range y {
		sx = append(sx, v)
	}
	lt := len(y)
	avg := f(sx) / float64(lt)
	return avg
}
