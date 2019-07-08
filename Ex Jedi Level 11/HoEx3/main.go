package main

import "fmt"

type customErr struct {
	custErr string
}

func (ce customErr) Error() string {
	fmt.Print("Error Func - ", ce.custErr)
	return ce.custErr
}

func main() {
	c := customErr{custErr: "Custom Message"}
	foo(c)
}

func foo(e error) {
	fmt.Println(" - Foo Func -", e)
}
