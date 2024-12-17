package main

import (
	"example.com/calendar"
	"fmt"
	"log"
)

func main() {
	event := calendar.Event{}

	err := event.SetTitle("Parent's Anniversary")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2010)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(10)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Date)
	fmt.Println(event.Title())

}
