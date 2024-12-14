package playground

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi Player, %v. Welcome!", name)
	return message
}

func Run() {
	fmt.Println("Hello, playground")
}

type Address struct {
	Street, City, State, PostalCode string
}

type Subscriber struct {
	Name   string
	Rate   float64
	active bool
	Address
}
