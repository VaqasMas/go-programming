package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int

var wg sync.WaitGroup

func main() {

	fmt.Println("Before CPUs", runtime.NumCPU())
	fmt.Println("Before Goroutines", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	fmt.Println("--------")
	go bar()
	fmt.Println("Mid CPUs", runtime.NumCPU())
	fmt.Println("Mid Goroutines", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("After CPUs", runtime.NumCPU())
	fmt.Println("After Goroutines", runtime.NumGoroutine())
}

func foo() {
	for i := 1; i <= 10; i++ {
		//fmt.Println("This is foo Func")
		fmt.Println("foo Goroutines", runtime.NumGoroutine())
	}
	wg.Done()
}

func bar() {
	for i := 1; i <= 10; i++ {
		//fmt.Println("This is bar Func")
		fmt.Println("bar Goroutines", runtime.NumGoroutine())
	}
	wg.Done()
}
