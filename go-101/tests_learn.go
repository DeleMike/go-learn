package main

import (
	"example.com/prose"
	"fmt"
)

func main() {
	fmt.Println(prose.JoinWithCommas([]string{"rat", "cat", "dog"}))
}
