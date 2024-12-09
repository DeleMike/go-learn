package main

import (
	"fmt"
)

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func createPointer() *float64 {
	myFloat := 98.5
	return &myFloat
}

func double(number *int) {
	*number *= *number
}

func sayHi(name string) string {
	return fmt.Sprintf("Hello, %s\n", name)
}
