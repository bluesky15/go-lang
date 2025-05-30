package main

import (
	"fmt"
	"reflect"
)

func main() {
	//int
	var x = 5
	// string
	var y = "there"
	//bool
	var z = true
	//float64
	var num = 5.0
	fmt.Println(checkType(num))
	fmt.Println(checkType(x))
	fmt.Println(checkType(y))
	fmt.Println(checkType(z))
}

func checkType(i any) string {
	//fmt.Printf("the type is :%T\n",i)
	return reflect.TypeOf(i).Name()
}