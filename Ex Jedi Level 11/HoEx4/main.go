package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		se := fmt.Errorf("square root of negative number: %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", se}
	}
	return 42, nil
}

// Didn't do this exercise on my own. Had to copy the example from previous lecture to complete this ex.
