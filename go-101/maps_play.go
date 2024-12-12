package main

import "fmt"

func main() {
	ranks := make(map[string]int)
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 4

	fmt.Println(ranks)

	myMap := map[string]int{
		"male":   10,
		"female": 20,
	}

	myMap["male"] = myMap["male"] * 5
	myMap["female"]++
	value, ok := myMap["r"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("not found")
	}
	fmt.Println(myMap)

}
