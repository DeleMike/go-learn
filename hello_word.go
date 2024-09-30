package main

import "fmt"

func main() {
	// accountAge := 2.6
	// accountAgeInt := int(accountAge)

	// fmt.Println("Your account has existed for", accountAgeInt, "years")
	// fmt.Printf("I am %v years old", 10)

	playWithFormatting()
}

func playWithFormatting() {
	const name = "Andrew ng"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)

	fmt.Printf(msg)
}