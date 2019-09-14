package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	fmt.Printf("Hello \n")
	dat, err := ioutil.ReadFile("test.txt")
	check(err)
	fmt.Print(string(dat))
}
