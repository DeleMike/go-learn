package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	//myStruct.number = 3.14
	//myStruct.word = "hello"
	////myStruct.toggle = true
	//fmt.Println(myStruct.number, "||", myStruct.word, "||", myStruct.toggle)
	//var bolts part
	//bolts.description = "This is a 6 by 6 bolt"
	//bolts.count = 18
	//showInfo(bolts)

	//var s subscriber
	//applyDiscount(&s)
	//fmt.Println(s.rate)

	var value myStruct
	value.myField = 3
	var pointer *myStruct = &value
	fmt.Println((*pointer).myField)
}

type myStruct struct {
	myField int
}
type part struct {
	description string
	count       int
}

func showInfo(p part) {
	fmt.Println(p.description, "||", p.count)
}
