package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("OS:", runtime.GOARCH)
}
