package main

import (
	"fmt"
)

func main() {
	amount := 6

	double(&amount)
	fmt.Println(amount)
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
