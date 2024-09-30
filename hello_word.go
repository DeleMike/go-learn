package main

import "fmt"

func main() {
	// accountAge := 2.6
	// accountAgeInt := int(accountAge)

	// fmt.Println("Your account has existed for", accountAgeInt, "years")
	// fmt.Printf("I am %v years old", 10)

	// playWithFormatting()
	// ifConditionals()
	// fmt.Println(getCoo())
	// testStruct()
	rectangle := Rect{width: 4, height: 5}
	area := fmt.Sprintf("The area of the rectangele is %.1f", (rectangle.area()))
	fmt.Println(area)
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
	} else {
		fmt.Println("Message not sent")
	}

	if length := -1; length < 1 {
		fmt.Println("Message code is invalid")
	}
}

func getCoo() (x, y int) {
	x, y = 2, 3
	return
}

func testStruct() {
	type Car struct {
		Make  string
		Model string
		age   float64
		price float64
		Wheel struct {
			Radius   float64
			Material string
		}
	}

	car := Car{Make: "Mercedes-Benz", Model: "gx-500", age: 0.2, price: 350000}
	car.Wheel.Material = "Black Puma"
	car.Wheel.Radius = 30.5
	fmt.Println("Car is ", car)

	// anonymous structs
	user := struct {
		name string
		age  int
	}{
		name: "Ada Lovelace",
		age:  70,
	}

	fmt.Println("User is ", user)

}

type Rect struct {
	width float64
	height float64

}

func (r Rect) area() float64{
	return r.width * r.height
}
