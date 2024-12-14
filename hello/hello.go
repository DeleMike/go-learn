package main

import (
	"fmt"

	"example.com/playground"
)

func main() {
	// Get a greeting message and print it.
	message := playground.Hello("Akin")
	playground.Run()
	fmt.Println(message)

	var address playground.Address
	address.Street = "123 Oak Street"
	address.State = "NE"
	address.City = "Tokyo"
	address.PostalCode = "68111"
	fmt.Println(address)
	subscriber := playground.Subscriber{
		Name: "Akin",
		Rate: 450.0,
	}

	subscriber.Address = address

	fmt.Println(subscriber)
}
