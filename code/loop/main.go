package main

import "fmt"

func main(){
	var list = []string{"Apple", "orange","pineapple"}

	for index,value:=range list {
		fmt.Printf("Index %d and value %s\n",index,value)

	}
}