package main

import (
	"example.com/calendar"
	"fmt"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2010)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(10)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}
