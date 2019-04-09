package main

import "fmt"

var x = 7
var y =8
func main() {
	x,y = y,x // tuple assignment

	fmt.Printf("vale of x is: %d",x)
	fmt.Println()
	fmt.Printf("vale of y is: %d",y)

	fmt.Println()
	fmt.Println("greatest common divisor is :" ,gcd(3,9))
	//fmt.Println(3%9)
	fmt.Println("The 3 th fibonacci number is ",fib(3))
}

func gcd(x, y int) int {
	for y!=0{
		x, y = y, x%y
		fmt.Println(x,y)
	}
	return x
}

func fib (n int) int{
	x, y  = 0,1
	for i:=0;i<n;i++{
		x , y = y , x+y
	}
	return x
}
