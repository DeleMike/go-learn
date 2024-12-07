package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sayHi("Akindele"))
	err := errors.New("this is an error")
	fmt.Println(err.Error())
}

func sayHi(name string) string {
	return fmt.Sprintf("Hello, %s\n", name)
}
