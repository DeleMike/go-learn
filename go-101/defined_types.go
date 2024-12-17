package main

import "fmt"

type Litres float64
type Gallons float64

func main() {
	var carFuel Gallons
	var bushFuel Litres

	carFuel = Gallons(1.2)
	bushFuel = Litres(4.5)

	fmt.Println(carFuel, bushFuel+2)
	carFuel += toGallons(Litres(40.0))
	bushFuel += toLitres(Gallons(30.0))
	fmt.Printf("carFuel: %0.1f || bushFuel: %0.1f\n", carFuel, bushFuel)

	carFuel = Gallons(Litres(2.0))
	fmt.Println(Litres(3) == Litres(Gallons(3)))

	user := AUser{
		id:     10,
		name:   "Akindele",
		age:    24,
		active: true,
	}

	fmt.Printf("I am %2d years old\n", user.getAge())
	fmt.Println(user.greetPerson())

	number := Number(4)
	fmt.Println(number)
	number.Double()
	fmt.Println(number)
	number.Display()

}

func toGallons(l Litres) Gallons {
	return Gallons(l * 0.264)
}

func toLitres(g Gallons) Litres {
	return Litres(g * 3.785)
}

type AUser struct {
	id     int
	name   string
	age    int
	active bool
}

func (u AUser) greetPerson() string {
	return "Hello " + u.name
}

func (u AUser) getAge() int {
	return u.age
}

type Number int

func (n *Number) Double() {
	*n *= 2
}

func (n *Number) Display() {
	fmt.Println(*n)
}
