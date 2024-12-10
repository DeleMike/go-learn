package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	myFun("a", "v", "h")
	//notes := make ([]int, 2)
	//fmt.Println(notes)
	var letters = []string{"a", "b", "c"}
	//fmt.Println(letters)
	part1 := letters[:2]
	myFun("a", part1...)

	//fmt.Println(part1)
	//part2 := letters[1:]
	//fmt.Println(part2)
	//letters[1] = "X"
	//fmt.Println("=====================================")
	//
	//fmt.Println(letters)
	//fmt.Println(part1)
	//fmt.Println(part2)

}

func myFun(param1 string, param2 ...string) {
	fmt.Println(param1, param2)
}
