package main

import "fmt"

func main() {
	var x int
	var p = &x // &x means memory address of x
	*p = 67 // *p means value at address which is stored in p variable.
	fmt.Print(*p)

}
