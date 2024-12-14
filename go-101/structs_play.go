package main

import "fmt"

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	myStruct.number = 3.14
	myStruct.word = "hello"
	//myStruct.toggle = true
	fmt.Println(myStruct.number, "||", myStruct.word, "||", myStruct.toggle)
}
