package main

import "fmt"

func main() {
	// accountAge := 2.6
	// accountAgeInt := int(accountAge)

	// fmt.Println("Your account has existed for", accountAgeInt, "years")
	// fmt.Printf("I am %v years old", 10)

	// playWithFormatting()
	// ifConditionals()
	fmt.Println(getCoo())
}

func playWithFormatting() {
	
	const name = "Andrew ng"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)

	fmt.Printf(msg)
}

func ifConditionals() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	}else{
		fmt.Println("Message not sent")
	}

	if length := -1; length < 1 {
		fmt.Println("Message code is invalid")
	}
}

func getCoo() (x,y int)  {
	x, y = 2,3
	return
}